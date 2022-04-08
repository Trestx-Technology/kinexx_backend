package db

import (
	"kinexx_backend/pkg/entity"
	"time"
)

type AccountService interface {
	SignUp(cred Credentials) (string, error)
	Login(cred Credentials) (entity.ProfileDB, string, error)
	SocialMedialogin(cred Credentials) (string, entity.ProfileDB, int64, error)
	VerifyEmailOTP(cred Credentials) (string, error)
	SendEmailOTP(email string) (string, error)
	ChangePassword(cred Credentials) (string, error)
}

// Credentials represents the model for an credentials
type Credentials struct {
	Email             string           `bson:"email" json:"email"`
	Password          string           `bson:"password" json:"password"`
	CreatedDate       time.Time        `bson:"created_date" json:"createdDate"`
	Status            string           `bson:"status" json:"status"`
	VerificationCode  string           `bson:"verification_code" json:"verificationCode"`
	FirstName         string           `bson:"first_name" json:"firstName"`
	LastName          string           `bson:"last_name" json:"lastName"`
	PhoneNo           string           `bson:"phone_no" json:"phoneNumber"`
	CountryCode       string           `bson:"country_code" json:"countryCode"`
	Designation       string           `bson:"designation" json:"designation"`
	Name              string           `bson:"name" json:"name"`
	DOB               string           `bson:"dob" json:"dob"`
	EmailLoginOTP     string           `bson:"email_login_otp" json:"emailLoginOtp"`
	EmailSentTime     time.Time        `bson:"email_sent_time" json:"emailSentTime"`
	VerifiedTime      time.Time        `bson:"verified_time" json:"verifiedTime"`
	UrlToProfileImage string           `bson:"url_to_profile_image" json:"urlToProfileImage"`
	About             string           `bson:"about" json:"about"`
	Address           entity.AddressDB `bson:"address" json:"address"`
	AuthToken         string           `bson:"auth_token" json:"auth_token"`
	Type              string           `bson:"type" json:"type"`
	// TermsChecked        bool             `bson:"terms_and_condition" json:"termsAndCondition"`
	PasswordResetCode   string      `bson:"password_reset_code" json:"passwordResetCode"`
	PasswordResetTime   time.Time   `bson:"password_reset_time" json:"passwordResetTime"`
	LastLoginDeviceInfo interface{} `bson:"last_login_device_info" json:"lastLoginDeviceInfo"`
	LastLoginLocation   string      `bson:"last_login_location" json:"lastLoginLocation"`
	Online              bool        `bson:"online" json:"online"`
}

type OTP struct {
	Email string `bson:"email" json:"email"`
	OTP   string `bson:"otp_code" json:"otp_code"`
}

type EmailVerification struct {
	Email string `bson:"email" json:"email"`
}
