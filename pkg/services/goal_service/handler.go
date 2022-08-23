package goalservice

import (
	"encoding/json"
	"io/ioutil"
	"kinexx_backend/pkg/entity"
	"kinexx_backend/pkg/repository/goals"
	goalDB "kinexx_backend/pkg/services/goal_service/db"
	"net/http"
	"strings"
	"time"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	goalservice = goalDB.NewGoalService(goals.NewGoalRepository("goals"))
)

func MakeGoal(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("making goal", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var goal *entity.GoalDB
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &goal)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := goalservice.MakeGoal(goal)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to make goal"))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to make goal"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("comment updated", logrus.Fields{
		"duration": duration,
	})
}
func GetAllGoals(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("making group", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}

	data, err := goalservice.GetAllGoal()
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to delete group"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to delete group"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("comment updated", logrus.Fields{
		"duration": duration,
	})
}
