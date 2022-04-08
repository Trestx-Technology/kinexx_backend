package movies_service

import (
	"encoding/json"
	"kinexx_backend/pkg/services/movies_service/db"
	"net/http"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func FindMovies(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("getting connection", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var query = mux.Vars(r)["query"]
	var page = mux.Vars(r)["page"]
	data, err := db.FindMovies(query, page)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get connection data"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "User Not Available"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("connection sent", logrus.Fields{
		"duration": duration,
	})
}
