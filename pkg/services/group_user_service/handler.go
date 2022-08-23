package groupuserservice

import (
	"encoding/json"
	groupUserdb "kinexx_backend/pkg/services/group_user_service/db"
	"kinexx_backend/pkg/services/group_user_service/group_user"
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
	groupUserService = groupUserdb.NewGroupUserService(group_user.NewGroupUserRepository("group_users"))
)

func AddUserToGroup(w http.ResponseWriter, r *http.Request) {
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
	// var group_user *entity.GroupUserDB

	var groupID = mux.Vars(r)["groupID"]
	var userID = mux.Vars(r)["userID"]
	var status = mux.Vars(r)["status"]
	data, err := groupUserService.AddUserToGroup(groupID, userID, status)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set group"})
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

func RemoveUserFromGroup(w http.ResponseWriter, r *http.Request) {
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

	var groupID = mux.Vars(r)["groupID"]
	var userID = mux.Vars(r)["userID"]

	err = groupUserService.RemoveUserFromGroup(userID, groupID)
	// why not return (data, err) instead of err in RemoveUserFromGroup

	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to remove record"})
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

func GetGroupsForUser(w http.ResponseWriter, r *http.Request) {
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

	var userID = mux.Vars(r)["userID"]

	SliceOfGroups, err := groupUserService.GetGroupsForUser(userID)
	// why not return (data, err) instead of err in RemoveUserFromGroup

	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "No group available"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": SliceOfGroups})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("comment updated", logrus.Fields{
		"duration": duration,
	})
}

func GetUsersInGroup(w http.ResponseWriter, r *http.Request) {
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

	var groupID = mux.Vars(r)["groupID"]

	SliceOfUsers, err := groupUserService.GetUsersInGroup(groupID)

	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set comment"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "No group available"})
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
