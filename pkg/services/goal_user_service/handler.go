package goaluserservice

import (
	"encoding/json"
	goal_group "kinexx_backend/pkg/services/goal_user_service/goals_group"

	goalGroupdb "kinexx_backend/pkg/services/goal_user_service/db"

	"net/http"
	"strings"
	"time"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	goalUserService = goalGroupdb.NewGoalUserService(goal_group.NewGoalGroupRepository("goal_groups"))
)

func AddUserToGoal(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting comment", logrus.Fields{
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
	// var goal_user *entity.GoalUserDB

	var goalID = mux.Vars(r)["goalID"]
	var userID = mux.Vars(r)["userID"]

	data, err := goalUserService.AddUserToGoal(goalID, userID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set comment"})
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

func RemoveUserFromGoal(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting comment", logrus.Fields{
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

	var goalID = mux.Vars(r)["goalID"]
	var userID = mux.Vars(r)["userID"]

	err = goalUserService.RemoveUserFromGoal(userID, goalID)
	// why not return (data, err) instead of err in RemoveUserFromGoal

	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set comment"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": "User Removed"})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("comment updated", logrus.Fields{
		"duration": duration,
	})
}

func GetGoalsForUser(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting comment", logrus.Fields{
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

	var userID = mux.Vars(r)["groupID"]

	SliceOfGoals, err := goalUserService.GetGoalsForUser(userID)
	// why not return (data, err) instead of err in RemoveUserFromGoal

	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set comment"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": SliceOfGoals})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("comment updated", logrus.Fields{
		"duration": duration,
	})
}

func GetUsersInGoal(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting comment", logrus.Fields{
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

	var goalID = mux.Vars(r)["goalID"]

	SliceOfUsers, err := goalUserService.GetUsersInGoal(goalID)
	// why not return (data, err) instead of err in RemoveUserFromGoal

	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set comment"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": SliceOfUsers})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("comment updated", logrus.Fields{
		"duration": duration,
	})
}
