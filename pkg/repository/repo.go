package repository

import "go.mongodb.org/mongo-driver/bson"

type Repo interface {
	Count(filter bson.M) (int64, error)
}
