package repository

import (
	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewRepo(collectionName string) Repo {
	return &repo{
		CollectionName: collectionName,
	}
}
func (r *repo) Count(filter bson.M) (int64, error) {
	return trestCommon.Count(filter, bson.M{}, r.CollectionName)
}
