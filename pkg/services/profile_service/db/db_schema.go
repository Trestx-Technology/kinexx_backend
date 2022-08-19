package db

import (
	"kinexx_backend/pkg/entity"
)

type ProfileService interface {
	UpdateProfile(profile *Profile, userid string) (string, error)
	SelectPromoVideos(profile *Profile, userid string) (string, error)
	UpdatePromoVideos(profile *Profile, userid string) (string, error)
	UpdateHobbies(profile *Profile, userid string) (string, error)
	UpdateMovies(profile *Profile, userid string) (string, error)
	ChangeUserStatus(userid string) (string, error)
	GetProfile(userID string) (entity.ProfileDB, int64, error)
	GetExp(userID string) ([]entity.Experience, error)
	ChangePassword(profile *Profile, userid string) (string, error)
	GetAllProfile(userID string) ([]entity.ProfileDB, error)
	GetUserProfile(userID string) (entity.ProfileDB, int64, error)
	SearchUser(search string) ([]entity.ProfileDB, error)
	UpdateVideo(profile *Profile, userid string) (string, error)
	AddAndUpdateExperience(exp entity.Experience, userID, experienceID string) (string, error)
	RemoveExperience(experienceID, userID string) (string, error)
	BlockUser(userID string, blockUserID string) (string, error)
}

type Profile struct {
	Status                  string      `bson:"status" json:"status"`
	FirstName               string      `bson:"first_name" json:"firstName"`
	Name                    string      `bson:"name" json:"name"`
	DOB                     string      `bson:"dob" json:"dob"`
	LastLoginDeviceInfo     interface{} `bson:"last_login_device_info" json:"lastLoginDeviceInfo"`
	LastLoginLocation       string      `bson:"last_login_location" json:"lastLoginLocation"`
	Email                   string      `bson:"email" json:"email"`
	Password                string      `bson:"password" json:"password"`
	LastName                string      `bson:"last_name" json:"lastName"`
	PhoneNo                 string      `bson:"phone_no" json:"phoneNumber"`
	CountryCode             string      `bson:"country_code" json:"countryCode"`
	Designation             string      `bson:"designation" json:"designation"`
	Gender                  string      `bson:"gender" json:"gender"`
	Featured                string      `bson:"featured" json:"featured"`
	Speciality              []string    `bson:"speciality" json:"speciality"`
	Categories              []string    `bson:"categories" json:"categories"`
	Address                 string      `bson:"address" json:"address"`
	Country                 string      `bson:"country" json:"country"`
	Pin                     string      `bson:"pin" json:"pin"`
	City                    string      `bson:"city" json:"city"`
	State                   string      `bson:"state" json:"state"`
	About                   string      `bson:"about" json:"about"`
	UrlToProfileImage       string      `bson:"url_to_profile_image" json:"urlToProfileImage"`
	Movie                   string      `json:"movie"`
	Hobby                   string      `json:"hobby"`
	PortfolioVideo          string      `json:"portfolio_video"`
	SelectedPortFolioVideo  string      `json:"selected_portfolio_video"`
	ProfileVideo            string      `json:"profile_video"`
	Online                  bool        `bson:"online" json:"online"`
	PortfolioVideosCaptions string      `json:"portfolio_videos_captions"`
}
