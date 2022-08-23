package handlers

import (
	"encoding/json"
	"net/http"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func Handler(w http.ResponseWriter, err error, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get credentials"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		err := json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(bson.M{"status": true, "data": data})
	if err != nil {
		return
	}

}
