package viewContentHandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	handlers "kinexx_backend/pkg/handler"
	viewContentDB "kinexx_backend/pkg/services/view_content_service/db"
	viewContentEntity "kinexx_backend/pkg/services/view_content_service/entity"
	"kinexx_backend/pkg/services/view_content_service/repository"
	"kinexx_backend/pkg/utils"
	"net/http"
	"time"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	viewContentService = viewContentDB.NewService(viewContentRepository.NewRepository("view_contents"))
)

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
	data, err := viewContentService.Add(brand)
	data.Banner = utils.CreatePreSignedDownloadUrl(data.Banner)
	data.Cover = utils.CreatePreSignedDownloadUrl(data.Cover)
	data.VideoURL = utils.CreatePreSignedDownloadUrl(data.VideoURL)
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
	data, err := viewContentService.Update(brand, mux.Vars(r)["id"])
	data.Banner = utils.CreatePreSignedDownloadUrl(data.Banner)
	data.Cover = utils.CreatePreSignedDownloadUrl(data.Cover)
	data.VideoURL = utils.CreatePreSignedDownloadUrl(data.VideoURL)
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
	data, err := viewContentService.Delete(mux.Vars(r)["id"])
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})

}

func GetForCampaign(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := viewContentService.GetForCampaign(mux.Vars(r)["campaign_id"])
	for i := range data {
		data[i].Banner = utils.CreatePreSignedDownloadUrl(data[i].Banner)
		data[i].Cover = utils.CreatePreSignedDownloadUrl(data[i].Cover)
		data[i].VideoURL = utils.CreatePreSignedDownloadUrl(data[i].VideoURL)
	}
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}
func parseJSON(w http.ResponseWriter, r *http.Request) (*viewContentEntity.ViewContentDB, bool) {
	var brand *viewContentEntity.ViewContentDB
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
