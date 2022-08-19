package db

import (
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository"
	hobby_db "kinexx_backend/pkg/services/hobby_service/db"
	movies_service "kinexx_backend/pkg/services/movies_service/db"
	"strconv"
	"strings"

	"errors"
	"time"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = repository.NewProfileRepository("users")
)

type profileService struct{}

func NewProfileService(repository repository.ProfileRepository) ProfileService {
	repo = repository
	return &profileService{}
}

func (*profileService) UpdateProfile(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	if profile.FirstName != "" {
		setParameters["first_name"] = profile.FirstName
	}
	if profile.LastName != "" {
		setParameters["last_name"] = profile.LastName
	}
	if profile.PhoneNo != "" {
		setParameters["phone_no"] = profile.PhoneNo
	}
	if profile.Designation != "" {
		setParameters["designation"] = profile.Designation
	}
	if profile.Name != "" {
		setParameters["name"] = profile.Name
	}
	if profile.Gender != "" {
		setParameters["gender"] = profile.Gender
	}
	if profile.Featured != "" {
		setParameters["featured"] = profile.Featured
	}
	if profile.DOB != "" {
		setParameters["dob"] = profile.DOB
	}
	if len(profile.Speciality) > 0 {
		setParameters["speciality"] = profile.Speciality
	}
	if len(profile.Categories) > 0 {
		setParameters["categories"] = profile.Categories
	}
	if profile.Status != "" {
		setParameters["status"] = profile.Status
	}
	if profile.Address != "" {
		setParameters["address.address"] = profile.Address
	}
	if profile.State != "" {
		setParameters["address.state"] = profile.State
	}
	if profile.City != "" {
		setParameters["address.city"] = profile.City
	}
	if profile.Country != "" {
		setParameters["address.country"] = profile.Country
	}
	if profile.Pin != "" {
		setParameters["address.pin"] = profile.Pin
	}
	if profile.UrlToProfileImage != "" {
		setParameters["url_to_profile_image"] = profile.UrlToProfileImage
	}
	if profile.About != "" {
		setParameters["about"] = profile.About
	}
	if profile.SelectedPortFolioVideo != "" {
		setParameters["selected_portfolio_video"] = profile.SelectedPortFolioVideo
	}
	setParameters["online"] = true
	if profile.Movie != "" {
		movieID := []int{}
		for _, idS := range strings.Split(profile.Movie, ",") {
			mov, err := strconv.Atoi(idS)
			if err == nil {
				movieID = append(movieID, mov)
			}
		}
		setParameters["movies"] = movieID
	}
	if profile.Hobby != "" {
		setParameters["hobbies"] = strings.Split(profile.Hobby, ",")
	}
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}
	return result, nil
}
func (*profileService) UpdateVideo(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	setParameters["profile_video"] = strings.Split(profile.ProfileVideo, ",")
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}
func (*profileService) UpdateMovies(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	movieID := []int{}
	for _, idS := range strings.Split(profile.Movie, ",") {
		mov, err := strconv.Atoi(idS)
		if err == nil {
			movieID = append(movieID, mov)
		}
	}
	setParameters["movies"] = movieID

	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}
func (*profileService) UpdateHobbies(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	setParameters["hobbies"] = strings.Split(profile.Hobby, ",")
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}
func (*profileService) UpdatePromoVideos(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	setParameters["portfolio_videos"] = strings.Split(profile.PortfolioVideo, ",")
	setParameters["portfolio_videos_captions"] = strings.Split(profile.PortfolioVideosCaptions, ",")
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}
func (*profileService) ChangeUserStatus(userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	user, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	setParameters["online"] = !user.Online
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
			})
		return "", err
	}

	return result, nil
}
func (*profileService) SelectPromoVideos(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	setParameters["selected_portfolio_video"] = profile.SelectedPortFolioVideo
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}

	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}
func (*profileService) GetProfile(userID string) (entity.ProfileDB, int64, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return entity.ProfileDB{}, 0, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.ProfileDB{}, 0, err
	}
	profile, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, 0, err
	}
	// profile.Password = ""
	newUrl := createPreSignedDownloadUrl(profile.UrlToProfileImage)
	profile.UrlToProfileImage = newUrl
	per := calculatePercentage(profile)
	profile.MoviesList, _ = movies_service.GetMovie(profile.Movies, "0")
	profile.HobbiesList, _ = hobby_db.GetHobbies(profile.Hobbies, "0")
	profile.SelectedPortFolioVideo = createPreSignedDownloadUrl(profile.SelectedPortFolioVideo)
	for j := range profile.PortfolioVideos {
		profile.PortfolioVideos[j] = createPreSignedDownloadUrl(profile.PortfolioVideos[j])
	}
	for j := range profile.ProfileVideo {
		profile.ProfileVideo[j] = createPreSignedDownloadUrl(profile.ProfileVideo[j])
	}
	for j := range profile.Experiences {
		profile.Experiences[j].Video = createPreSignedDownloadUrl(profile.Experiences[j].Video)
	}
	return profile, per, nil
}

func (*profileService) GetExp(userID string) ([]entity.Experience, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return []entity.Experience{}, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return []entity.Experience{}, err
	}
	profile, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return []entity.Experience{}, err
	}

	for j := range profile.Experiences {
		profile.Experiences[j].Video = createPreSignedDownloadUrl(profile.Experiences[j].Video)
	}
	return profile.Experiences, nil
}

func (*profileService) GetUserProfile(userID string) (entity.ProfileDB, int64, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return entity.ProfileDB{}, 0, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.ProfileDB{}, 0, err
	}
	profile, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, 0, err
	}
	// profile.Password = ""
	newUrl := createPreSignedDownloadUrl(profile.UrlToProfileImage)
	profile.UrlToProfileImage = newUrl
	per := calculatePercentage(profile)
	profile.MoviesList, _ = movies_service.GetMovie(profile.Movies, "1")
	profile.HobbiesList, _ = hobby_db.GetHobbies(profile.Hobbies, "0")
	profile.SelectedPortFolioVideo = createPreSignedDownloadUrl(profile.SelectedPortFolioVideo)
	for j := range profile.PortfolioVideos {
		profile.PortfolioVideos[j] = createPreSignedDownloadUrl(profile.PortfolioVideos[j])
	}
	for j := range profile.ProfileVideo {
		profile.ProfileVideo[j] = createPreSignedDownloadUrl(profile.ProfileVideo[j])
	}
	if strings.Contains(profile.Featured, "s3") {
		newUrl = createPreSignedDownloadUrl(profile.Featured)
		if newUrl != "" {
			profile.Featured = newUrl
		}
	}
	return profile, per, nil
}
func (*profileService) AddAndUpdateExperience(exp entity.Experience, experienceID, userID string) (string, error) {
	uid, _ := primitive.ObjectIDFromHex(userID)
	profile, err := repo.FindOne(bson.M{"_id": uid}, bson.M{})
	if err != nil {
		return "", err
	}
	if experienceID != "" {
		id, _ := primitive.ObjectIDFromHex(experienceID)
		for i := range profile.Experiences {
			if profile.Experiences[i].ExperienceID == id {
				exp.ExperienceID = id
				profile.Experiences[i] = exp
				break
			}
		}
	} else {
		exp.ExperienceID = primitive.NewObjectID()
		profile.Experiences = append(profile.Experiences, exp)
	}
	set := bson.M{"experiences": profile.Experiences}

	return repo.UpdateOne(bson.M{"_id": uid}, bson.M{"$set": set})
}

func (*profileService) BlockUser(userID string, blockUserID string) (string, error) {
	uid, _ := primitive.ObjectIDFromHex(userID)
	profile, err := repo.FindOne(bson.M{"_id": uid}, bson.M{})
	if err != nil {
		return "", err
	}
	found := false
	for i := range profile.Blocked {
		if profile.Blocked[i] == blockUserID {
			profile.Blocked = append(profile.Blocked[:i], profile.Blocked[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		profile.Blocked = append(profile.Blocked, blockUserID)
	}
	set := bson.M{"blocked": profile.Blocked}

	return repo.UpdateOne(bson.M{"_id": uid}, bson.M{"$set": set})
}

func (*profileService) RemoveExperience(experienceID, userID string) (string, error) {
	uid, _ := primitive.ObjectIDFromHex(userID)
	profile, err := repo.FindOne(bson.M{"_id": uid}, bson.M{})
	if err != nil {
		return "", err
	}
	if experienceID != "" {
		id, _ := primitive.ObjectIDFromHex(experienceID)
		for i := range profile.Experiences {
			if profile.Experiences[i].ExperienceID == id {
				profile.Experiences = append(profile.Experiences[:i], profile.Experiences[i+1:]...)
				break
			}
		}
	}
	set := bson.M{"experiences": profile.Experiences}

	return repo.UpdateOne(bson.M{"_id": uid}, set)
}

func (*profileService) GetAllProfile(userID string) ([]entity.ProfileDB, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return []entity.ProfileDB{}, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return []entity.ProfileDB{}, err
	}
	profile, err := repo.Find(bson.M{"_id": bson.M{"$ne": id}}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, err
	}
	// profile.Password = ""
	for i := range profile {
		newUrl := createPreSignedDownloadUrl(profile[i].UrlToProfileImage)
		profile[i].UrlToProfileImage = newUrl
		profile[i].Password = ""
		profile[i].MoviesList, _ = movies_service.GetMovie(profile[i].Movies, "1")
		profile[i].HobbiesList, _ = hobby_db.GetHobbies(profile[i].Hobbies, "0")
		profile[i].SelectedPortFolioVideo = createPreSignedDownloadUrl(profile[i].SelectedPortFolioVideo)
		for j := range profile[i].PortfolioVideos {
			profile[i].PortfolioVideos[j] = createPreSignedDownloadUrl(profile[i].PortfolioVideos[j])
		}
		for j := range profile[i].Experiences {
			profile[i].Experiences[j].Video = createPreSignedDownloadUrl(profile[i].Experiences[j].Video)
		}
		for j := range profile[i].ProfileVideo {
			profile[i].ProfileVideo[j] = createPreSignedDownloadUrl(profile[i].ProfileVideo[j])
		}
		if strings.Contains(profile[i].Featured, "s3") {
			newUrl = createPreSignedDownloadUrl(profile[i].Featured)
			if newUrl != "" {
				profile[i].Featured = newUrl
			}
		}
	}
	return profile, nil
}

func (*profileService) SearchUser(search string) ([]entity.ProfileDB, error) {
	filter := bson.A{}
	if search != "" {
		filter = append(filter, bson.M{"name": bson.M{"$regex": search, "$options": "i"}})
		filter = append(filter, bson.M{"email": bson.M{"$regex": search, "$options": "i"}})
		filter = append(filter, bson.M{"hobbies": bson.M{"$regex": search, "$options": "i"}})
	}
	profile, err := repo.Find(bson.M{"$or": filter}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, err
	}
	// profile.Password = ""
	for i := range profile {
		newUrl := createPreSignedDownloadUrl(profile[i].UrlToProfileImage)
		profile[i].UrlToProfileImage = newUrl
		profile[i].Password = ""
		profile[i].MoviesList, _ = movies_service.GetMovie(profile[i].Movies, "1")
		profile[i].HobbiesList, _ = hobby_db.GetHobbies(profile[i].Hobbies, "0")
		profile[i].SelectedPortFolioVideo = createPreSignedDownloadUrl(profile[i].SelectedPortFolioVideo)
		for j := range profile[i].PortfolioVideos {
			profile[i].PortfolioVideos[j] = createPreSignedDownloadUrl(profile[i].PortfolioVideos[j])
		}
		for j := range profile[i].Experiences {
			profile[i].Experiences[j].Video = createPreSignedDownloadUrl(profile[i].Experiences[j].Video)
		}
		for j := range profile[i].ProfileVideo {
			profile[i].ProfileVideo[j] = createPreSignedDownloadUrl(profile[i].ProfileVideo[j])
		}
		if strings.Contains(profile[i].Featured, "s3") {
			newUrl = createPreSignedDownloadUrl(profile[i].Featured)
			if newUrl != "" {
				profile[i].Featured = newUrl
			}
		}
	}
	return profile, nil
}

func GetProfilesForIDs(userIDs []string) ([]entity.ProfileDB, error) {
	filter := bson.A{}
	if len(userIDs) == 0 {
		return []entity.ProfileDB{}, errors.New("no userids")
	}
	for _, userID := range userIDs {
		id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			trestCommon.ECLog3(
				"GetProfile section",
				err,
				logrus.Fields{
					"user_id": userID,
				},
			)
			return []entity.ProfileDB{}, err
		}
		filter = append(filter, bson.M{"_id": id})
	}
	profile, err := repo.Find(bson.M{"$or": filter}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, err
	}
	// profile.Password = ""
	var profiles []entity.ProfileDB
	for prof := range profile {
		newUrl := createPreSignedDownloadUrl(profile[prof].UrlToProfileImage)
		profile[prof].UrlToProfileImage = newUrl
		profile[prof].Password = ""
		//profile[prof].MoviesList, _ = movies_service.GetMovie(profile[prof].Movies, "1")
		profile[prof].HobbiesList, _ = hobby_db.GetHobbies(profile[prof].Hobbies, "0")
		profile[prof].SelectedPortFolioVideo = createPreSignedDownloadUrl(profile[prof].SelectedPortFolioVideo)
		for j := range profile[prof].PortfolioVideos {
			profile[prof].PortfolioVideos[j] = createPreSignedDownloadUrl(profile[prof].PortfolioVideos[j])
		}
		for j := range profile[prof].ProfileVideo {
			profile[prof].ProfileVideo[j] = createPreSignedDownloadUrl(profile[prof].ProfileVideo[j])
		}
		if strings.Contains(profile[prof].Featured, "s3") {
			newUrl = createPreSignedDownloadUrl(profile[prof].Featured)
			if newUrl != "" {
				profile[prof].Featured = newUrl
			}
		}
		for j := range profile[prof].Experiences {
			profile[prof].Experiences[j].Video = createPreSignedDownloadUrl(profile[prof].Experiences[j].Video)
		}
		profiles = append(profiles, profile[prof])

	}
	return profiles, nil
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
func checkUser(userID string) (entity.ProfileDB, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"CheckUser section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.ProfileDB{}, err
	}
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}

func (*profileService) ChangePassword(profile *Profile, userid string) (string, error) {
	// if profile.Password == "" {
	// 	trestCommon.ECLog3("new password", errors.New("there was an error while updating the password"), logrus.Fields{"email": profile.Email})
	// 	return "", errors.New("there was an error while updating the password")
	// }
	salt := viper.GetString("salt")
	hash, err := bcrypt.GenerateFromPassword([]byte(profile.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash new password", err, logrus.Fields{"email": profile.Email})
		return "", err
	}
	password := string(hash)
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	userData, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog3("hash new password", err, logrus.Fields{"email": profile.Email})
		return "", err
	}
	if userData.Email != profile.Email {
		trestCommon.ECLog3("new password", errors.New("there was an error while updating the password"), logrus.Fields{"email": profile.Email})
		return "", errors.New("there was an error while updating the password")
	}
	if password != "" {
		setParameters["password"] = password
	}
	setParameters["update_time"] = time.Now()
	setParameters["password_updated_on_time"] = time.Now()
	setParameters["last_login_device_info"] = profile.LastLoginDeviceInfo
	setParameters["last_login_location"] = profile.LastLoginLocation
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	_, err = repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile password",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}
	_, err = sendConfirmationEmail(userData.Email)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": profile.Email})
	}
	return "password changed successfully", nil
}

func sendConfirmationEmail(email string) (string, error) {
	emailSentTime := time.Now()
	verificationCode := trestCommon.GetRandomString(16)
	sendCode, err := trestCommon.Encrypt(email + ":" + verificationCode)
	if err != nil {
		trestCommon.ECLog2("send verification email encryption failed", err)
		return "", err
	}
	_, err = trestCommon.SendPasswordConfirmation(email, sendCode)
	if err != nil {
		trestCommon.ECLog2("send verification email failed", err)
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"password_reset_confirmation_email_sent_time": emailSentTime}})
	if err != nil {
		trestCommon.ECLog2("send verification email update failed", err)
		return "", err
	}
	return "email sent successfully", nil
}
