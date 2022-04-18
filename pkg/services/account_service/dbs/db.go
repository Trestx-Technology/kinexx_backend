package db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository"
	"kinexx_backend/pkg/utils"
	"math/rand"
	"strconv"
	"strings"

	"errors"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	repo = repository.NewProfileRepository("users")
)

type accountService struct{}

func NewSignUpService(repository repository.ProfileRepository) AccountService {
	repo = repository
	return &accountService{}
}
func (*accountService) SignUp(cred Credentials) (string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("sign up failed no password", err)
		return "", err
	}
	cred.Email = strings.ToLower(cred.Email)
	_, err := checkUser(cred.Email, cred.PhoneNo)

	if err != nil {
		trestCommon.ECLog2("sign up user not found", err)
		if err.Error() == "mongo: no documents in result" {
			var serv *accountService
			tokenString, err := serv.hashAndInsertData(cred)
			if err != nil {
				trestCommon.ECLog3("sign up not successful", err, logrus.Fields{"email": cred.Email})
			}
			return tokenString, err
		} else {
			return "", err

		}
	}

	return "", errors.New("email already registed")
}

func (*accountService) SendVerificationEmail(email, pemail, uid string) (string, error) {
	emailSentTime := time.Now()
	if email == "" {
		return "email sent successfully", nil
	}
	email = strings.ToLower(email)
	verificationCode := trestCommon.GetRandomString(16)
	sendCode, err := trestCommon.Encrypt(email + ":" + verificationCode)
	if err != nil {
		trestCommon.ECLog2("send verification email encryption failed", err)
		return "", err
	}
	var userData entity.ProfileDB
	if pemail != "" {
		userData, _ = checkUser(pemail, "")
	} else {
		userData, _ = checkUser(email, "")
	}
	_, err = utils.SendVerificationCode(email, userData.FirstName+" "+userData.LastName, sendCode)
	if err != nil {
		trestCommon.ECLog2("send verification email failed", err)
		return "", err
	}
	if pemail != "" {
		_, err = repo.UpdateOne(bson.M{"email": pemail}, bson.M{"$set": bson.M{"email": email, "email_sent_time": emailSentTime, "verification_code": verificationCode}})
		if err != nil {
			trestCommon.ECLog2("send verification email update failed", err)
			return "", err
		}
		return trestCommon.CreateToken(uid, email, "", "")
	} else {
		_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"email_sent_time": emailSentTime, "verification_code": verificationCode}})
		if err != nil {
			trestCommon.ECLog2("send verification email update failed", err)
			return "", err
		}
		return "email sent successfully", nil
	}

}

func (*accountService) SendEmailOTP(email string) (string, error) {
	email = strings.ToLower(email)
	emailSentTime := time.Now()
	if email == "" {
		return "", errors.New("email id required")
	}
	_, err := checkUser(email, "")
	if err != nil {
		return "", errors.New("user account doesn't exist")
	}
	randomOTP := 1000 + rand.Intn(9999-1000)

	_, err = utils.EmailLoginOTP(email, email, strconv.Itoa(randomOTP), "change_password")
	if err != nil {
		trestCommon.ECLog2("send verification email failed", err)
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"email_sent_time": emailSentTime, "email_login_otp": strconv.Itoa(randomOTP)}})
	if err != nil {
		trestCommon.ECLog2("send verification email update failed", err)
		return "", err
	}
	return "email sent successfully", nil
}

func (serv *accountService) Login(cred Credentials) (entity.ProfileDB, string, error) {
	if cred.Password == "" {
		err := errors.New("password missing")
		trestCommon.ECLog2("login failed no password", err)
		return entity.ProfileDB{}, "", err
	}
	cred.Email = strings.ToLower(cred.Email)
	salt := viper.GetString("salt")
	userData, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog2("login failed user not found", err)
		return entity.ProfileDB{}, "", err
	}
	// if userData.Status == "created" {
	// 	err = errors.New("user not verified")
	// 	trestCommon.ECLog2("login failed user has not verified his/her email address", err)
	// 	return "", err
	// }
	if userData.Status == "deleted" || userData.Status == "archived" {
		err = errors.New("unauthorized user")
		trestCommon.ECLog2("login failed user has deleted or archived his profile", err)
		return entity.ProfileDB{}, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(cred.Password+salt))
	if err != nil {
		trestCommon.ECLog2("login failed password hash doesn't match", err)
		return entity.ProfileDB{}, "", err
	}
	tokenString, err := trestCommon.CreateToken(userData.ID.Hex(), cred.Email, userData.FirstName, userData.Status)
	if err != nil {
		trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{"email": cred.Email, "name": userData.FirstName, "status": userData.Status})
		return entity.ProfileDB{}, "", err
	}
	repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"login_time": time.Now()}})

	return userData, tokenString, nil
}
func (*accountService) VerifyEmailOTP(cred Credentials) (string, error) {
	cred.Email = strings.ToLower(cred.Email)
	userData, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", err
	}

	if cred.EmailLoginOTP != userData.EmailLoginOTP {
		err = errors.New("unauthorized user")
		trestCommon.ECLog3("verify user verification code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.EmailLoginOTP, "code provided by user": cred.EmailLoginOTP})
		return "", err
	}
	tokenString, err := trestCommon.CreateToken(userData.ID.Hex(), cred.Email, userData.FirstName, userData.Status)
	if err != nil {
		trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{"email": cred.Email, "name": userData.FirstName, "status": userData.Status})
		return "", err
	}
	return tokenString, nil
}
func (*accountService) ChangePassword(cred Credentials) (string, error) {
	userData, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	setFilter := bson.M{}
	salt := viper.GetString("salt")
	hash, err := bcrypt.GenerateFromPassword([]byte(cred.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	cred.Password = string(hash)
	setFilter["$set"] = bson.M{"verified_time": time.Now(), "password": cred.Password, "login_time": time.Now()}
	return repo.UpdateOne(bson.M{"_id": userData.ID}, setFilter)
}

func calculatePercentage(profile entity.ProfileDB) int64 {
	initial := 12
	if profile.FirstName != "" && profile.LastName != "" {
		initial = initial + 11
	}
	if profile.About != "" {
		initial = initial + 11
	}
	if profile.PhoneNo != "" {
		initial = initial + 11
	}
	if profile.UrlToProfileImage != "" {
		initial = initial + 11
	}
	if profile.UrlToProfileImage != "" {
		initial = initial + 11
	}
	if profile.Address.Address != "" {
		initial = initial + 3
	}
	if profile.Address.City != "" {
		initial = initial + 2
	}
	if profile.Address.Country != "" {
		initial = initial + 2
	}
	if profile.Address.Pin != "" {
		initial = initial + 2
	}
	if profile.Address.State != "" {
		initial = initial + 2
	}
	return int64(initial)
}

func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrlAWS(fileName, path)
			return downUrl
		}
	}
	return ""
}
func (*accountService) VerifyEmail(cred Credentials) (string, error) {

	userData, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", err
	}

	if cred.VerificationCode != userData.VerificationCode {
		err = errors.New("unauthorized user")
		trestCommon.ECLog3("verify user verification code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.VerificationCode, "code provided by user": cred.VerificationCode})
		return "", err
	}
	if userData.Status == "verified" {
		err = errors.New("user already verified")
		trestCommon.ECLog3("verify user verification user already verified", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"verified_time": time.Now(), "status": "verified"}})
	if err != nil {
		trestCommon.ECLog3("verify user unable to update status", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	return "verified", nil
}

func (*accountService) VerifyOTP(cred OTP) (string, error) {

	userData, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", err
	}

	if cred.OTP != userData.OTP {
		err = errors.New("unauthorized user")
		trestCommon.ECLog3("verify user verification code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.OTP, "code provided by user": cred.OTP})
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"verified_time": time.Now(), "phone_status": "verified"}})
	if err != nil {
		trestCommon.ECLog3("verify user unable to update status", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	return "verified", nil
}

func checkUser(email, mobile string) (entity.ProfileDB, error) {
	var userData entity.ProfileDB
	if email == "" {
		err := errors.New("email missing")
		trestCommon.ECLog2("check user failed no email", err)
		return userData, err
	}
	if !trestCommon.ValidateEmail(email) {
		err := errors.New("invalid email")
		trestCommon.ECLog2("check user failed invalid email", err)
		return userData, err
	}

	return repo.FindOne(bson.M{"$or": bson.A{bson.M{"email": email}, bson.M{"phone_no": mobile}}}, bson.M{})

}

func (*accountService) hashAndInsertData(cred Credentials) (string, error) {
	salt := viper.GetString("salt")

	hash, err := bcrypt.GenerateFromPassword([]byte(cred.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
		return "", err
	}
	cred.Password = string(hash)
	cred.CreatedDate = time.Now()
	cred.Status = "created"
	userid, err := repo.InsertOne(cred)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
		return "", nil
	}
	// _, err = serv.SendOTP(cred.Email, cred.PhoneNo)
	// if err != nil {
	// 	trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": cred.Email})
	// }
	return trestCommon.CreateToken(userid, cred.Email, "", cred.Status)
}
func (*accountService) SendResetLink(email string) (string, error) {
	var cred Credentials
	cred.Email = email
	_, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog2("user not found", err)
		return "", err

	}
	emailSentTime := time.Now()
	verificationCode := trestCommon.GetRandomString(16)
	resetCode, err := trestCommon.Encrypt(email + ":" + verificationCode)
	if err != nil {
		trestCommon.ECLog2("send reset link encryption failed", err)
		return "", err
	}
	_, err = trestCommon.SendResetPasswordLink(email, resetCode)
	if err != nil {
		trestCommon.ECLog2("send reset password link failed", err)
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"email_sent_time": emailSentTime, "password_reset_code": verificationCode}})
	if err != nil {
		trestCommon.ECLog2("send reset link update failed", err)
		return "", err
	}
	return "Reset link sent successfully", nil
}

func (*accountService) VerifyResetLink(cred Credentials) (string, string, error) {

	userData, err := checkUser(cred.Email, "")
	if err != nil {
		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
		return "", "", err
	}

	if cred.PasswordResetCode != userData.PasswordResetCode {
		err = errors.New("unauthorized user")
		trestCommon.ECLog3("verify user password reset code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.PasswordResetCode, "code provided by user": cred.PasswordResetCode})
		return "", "", err
	}
	_, err = repo.UpdateOne(bson.M{"_id": userData.ID}, bson.M{"$set": bson.M{"password_reset_time": time.Now()}})
	if err != nil {
		trestCommon.ECLog3("verify user unable to update status", err, logrus.Fields{"email": cred.Email})
		return "", "", err
	}

	return "verified", userData.Email, nil
}

// func (*accountService) UpdatePassword(cred Credentials) (string, error) {
// 	userData, err := checkUser(cred.Email, "")
// 	if err != nil {
// 		trestCommon.ECLog3("verify user not found", err, logrus.Fields{"email": cred.Email})
// 		return "", err
// 	}
// 	if cred.PasswordResetCode != userData.PasswordResetCode {
// 		err = errors.New("unauthorized user")
// 		trestCommon.ECLog3("verify user password reset code didn't match", err, logrus.Fields{"email": cred.Email, "db verify code": userData.PasswordResetCode, "code provided by user": cred.PasswordResetCode})
// 		return "", err
// 	}
// 	if cred.Password == "" {
// 		err := errors.New("password missing")
// 		trestCommon.ECLog2("update password failed no password", err)
// 		return "", err
// 	}
// 	if err != nil {
// 		trestCommon.ECLog2("update password failed user not found", err)
// 		return "", err
// 	}

// 	salt := viper.GetString("salt")

// 	hash, err := bcrypt.GenerateFromPassword([]byte(cred.Password+salt), 5)
// 	if err != nil {
// 		trestCommon.ECLog3("hash password", err, logrus.Fields{"email": cred.Email})
// 		return "", err
// 	}
// 	cred.Password = string(hash)
// 	_, err = repo.UpdateOne(bson.M{"email": cred.Email}, bson.M{"$set": bson.M{"password": cred.Password, "update_time": time.Now(), "password_reset_time": time.Now()}})
// 	if err != nil {
// 		trestCommon.ECLog2("password update failed", err)
// 		return "", err
// 	}

// 	return "password updated successfully", nil
// }

func (serv *accountService) SocialMedialogin(cred Credentials) (string, entity.ProfileDB, int64, error) {
	if cred.Email == "" {
		return "", entity.ProfileDB{}, 0, errors.New("email id required")
	}
	userData, err := checkUser(cred.Email, "")
	if err != nil {
		cred.CreatedDate = time.Now()
		cred.Status = "verified"

		userid, err := repo.InsertOne(cred)
		if err != nil {
			trestCommon.ECLog3("Insert failed", err, logrus.Fields{"email": cred.Email})
			return "", entity.ProfileDB{}, 0, err
		}
		userDetails, err := checkUser(cred.Email, "")
		tokenString, err := trestCommon.CreateToken(userid, cred.Email, "", "verified")
		if err != nil {
			trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{"email": cred.Email, "name": userDetails.FirstName, "status": userDetails.Status})
			return "", entity.ProfileDB{}, 0, err
		}
		userDetails.PasswordResetCode = ""
		userDetails.VerificationCode = ""
		userDetails.EmailLoginOTP = ""
		newUrl := createPreSignedDownloadUrl(userDetails.UrlToProfileImage)
		per := calculatePercentage(userDetails)

		userDetails.UrlToProfileImage = newUrl
		return tokenString, userDetails, per, nil
	}
	setFilter := bson.M{}
	setFilter["$set"] = bson.M{"login_time": time.Now(), "last_login_device_info": cred.LastLoginDeviceInfo, "last_login_location": cred.LastLoginLocation}
	_, err = repo.UpdateOne(bson.M{"_id": userData.ID}, setFilter)
	if err != nil {
		trestCommon.ECLog3("social media login unable to update status", err, logrus.Fields{"email": cred.Email})
		return "", entity.ProfileDB{}, 0, err
	}
	tokenString, err := trestCommon.CreateToken(userData.ID.Hex(), cred.Email, "", userData.Status)
	if err != nil {
		trestCommon.ECLog3("login failed unable to create token", err, logrus.Fields{"email": cred.Email, "name": userData.FirstName, "status": userData.Status})
		return "", entity.ProfileDB{}, 0, err
	}
	userData.PasswordResetCode = ""
	userData.VerificationCode = ""
	userData.EmailLoginOTP = ""

	newUrl := createPreSignedDownloadUrl(userData.UrlToProfileImage)
	per := calculatePercentage(userData)

	userData.UrlToProfileImage = newUrl
	return tokenString, userData, per, nil
}
