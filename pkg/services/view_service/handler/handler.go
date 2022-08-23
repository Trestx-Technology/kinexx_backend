package viewHandler

import (
	"encoding/json"
	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	handlers "kinexx_backend/pkg/handler"
	viewDB "kinexx_backend/pkg/services/view_service/db"
	viewEntity "kinexx_backend/pkg/services/view_service/entity"
	viewRepository "kinexx_backend/pkg/services/view_service/repository"
	"kinexx_backend/pkg/utils"
	"net/http"
	"time"
)

var (
	service = viewDB.NewService(viewRepository.NewRepository("view"))
)

func parseJSON(w http.ResponseWriter, r *http.Request) (*viewEntity.ViewDB, bool) {
	var brand *viewEntity.ViewDB
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &brand)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		err := json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		if err != nil {
			return nil, true
		}
		return nil, true
	}
	return brand, false
}

func Add(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	brand, done := parseJSON(w, r)
	if done {
		return
	}
	data, err := service.Add(brand)
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}

func Update(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	brand, done := parseJSON(w, r)
	if done {
		return
	}
	data, err := service.Update(brand, mux.Vars(r)["id"])
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}
func Delete(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := service.Delete(mux.Vars(r)["id"])
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})

}
func GetAll(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := service.GetAll()
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}
func GetMy(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dataToken, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := service.GetMy(dataToken["userid"].(string))
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dataToken, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := service.Get(mux.Vars(r)["id"], dataToken["userid"].(string))
	handlers.Handler(w, err, data[0])
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}
