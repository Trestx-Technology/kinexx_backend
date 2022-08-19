package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	Years        int                `bson:"years" json:"years"`
	ExperienceID primitive.ObjectID `bson:"experience_id" json:"experience_id"`
	Title        string             `bson:"title" json:"title"`
	CompanyName  string             `bson:"company_name" json:"companyName"`
	Location     string             `bson:"location" json:"location"`
	StartDate    string             `bson:"start_date" json:"startDate"`
	EndDate      string             `bson:"end_date" json:"endDate"`
	Industry     string             `bson:"industry" json:"industry"`
	Description  string             `bson:"description" json:"description"`
	Video        string             `bson:"video" json:"video"`
}
