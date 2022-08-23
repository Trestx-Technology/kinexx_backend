package entity

import (
	entity2 "kinexx_backend/pkg/services/hobby_service/entity"
	"time"

	"github.com/ryanbradynd05/go-tmdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Profile represents the model for an profile
type ProfileDB struct {
	ID                      primitive.ObjectID  `bson:"_id" json:"id"`
	Email                   string              `bson:"email" json:"email"`
	Status                  string              `bson:"status" json:"status"`
	FirstName               string              `bson:"first_name" json:"firstName"`
	LastName                string              `bson:"last_name" json:"lastName"`
	Name                    string              `bson:"name" json:"name"`
	DOB                     string              `bson:"dob" json:"dob"`
	Designation             string              `bson:"designation" json:"designation"`
	Gender                  string              `bson:"gender" json:"gender"`
	Featured                string              `bson:"featured" json:"featured"`
	PhoneNo                 string              `bson:"phone_no" json:"phoneNumber"`
	Address                 AddressDB           `bson:"address" json:"address"`
	About                   string              `bson:"about" json:"about"`
	UrlToProfileImage       string              `bson:"url_to_profile_image" json:"url_to_profile_image"`
	TermsChecked            bool                `bson:"terms_and_condition" json:"termsAndConditions"`
	Password                string              `bson:"password" json:"password"`
	CreatedTime             time.Time           `bson:"created_time" json:"createdTime"`
	EmailLoginOTP           string              `bson:"email_login_otp" json:"emailLoginOtp"`
	OTP                     string              `bson:"otp_code" json:"otp_code"`
	UpdateTime              time.Time           `bson:"update_time" json:"updateTime"`
	EmailSentTime           time.Time           `bson:"email_sent_time" json:"emailSentTime"`
	VerificationCode        string              `bson:"verification_code" json:"verificationCode"`
	PasswordResetCode       string              `bson:"password_reset_code" json:"passwordResetCode"`
	CountryCode             string              `bson:"country_code" json:"countryCode"`
	PasswordResetTime       time.Time           `bson:"password_reset_time" json:"passwordResetTime"`
	LastLoginDeviceID       string              `bson:"last_login_device_id" json:"lastLoginDeviceID"`
	LastLoginDeviceName     string              `bson:"last_login_device_name" json:"lastLoginDeviceName"`
	LastLoginLocation       string              `bson:"last_login_location" json:"lastLoginLocation"`
	Movies                  []int               `bson:"movies" json:"movies"`
	MoviesList              []*tmdb.Movie       `json:"movies_list"`
	Hobbies                 []string            `bson:"hobbies" json:"hobbies"`
	HobbiesList             []entity2.HobbiesDB `json:"hobbies_list"`
	PortfolioVideos         []string            `bson:"portfolio_videos" json:"portfolio_videos"`
	PortfolioVideosCaptions []string            `bson:"portfolio_videos_captions" json:"portfolio_videos_captions"`
	SelectedPortFolioVideo  string              `bson:"selected_portfolio_video" json:"selected_portfolio_video"`
	ProfileVideo            []string            `bson:"profile_video" json:"profile_video"`
	Online                  bool                `bson:"online" json:"online"`
	Experiences             []Experience        `bson:"experiences" json:"experiences"`
	Blocked                 []string            `bson:"blocked" json:"blocked"`
}

type AddressDB struct {
	Address     string      `bson:"address" json:"address"`
	Country     string      `bson:"country" json:"country"`
	Pin         string      `bson:"pin" json:"pin"`
	City        string      `bson:"city" json:"city"`
	State       string      `bson:"state" json:"state"`
	GeoLocation interface{} `bson:"geo_location" json:"geo_location,omitempty"`
}

type UnRegisteredUsersDB struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	DeviceID   string             `bson:"device_id" json:"DeviceID"`
	DeviceName string             `bson:"device_name" json:"DeviceName"`
	Location   string             `bson:"location" json:"Location"`
	Time       time.Time          `bson:"time" json:"Time"`
}
