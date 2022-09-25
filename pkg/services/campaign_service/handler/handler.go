package campaignHandler

import (
	"encoding/json"
	"io/ioutil"
	handlers "kinexx_backend/pkg/handler"
	campaignDB "kinexx_backend/pkg/services/campaign_service/db"
	campaignEntity "kinexx_backend/pkg/services/campaign_service/entity"
	"kinexx_backend/pkg/services/campaign_service/repository"
	"kinexx_backend/pkg/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	campaignService = campaignDB.NewCampaignService(repository.NewCampaignRepository("campaigns"))
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
	data, err := campaignService.AddCampaign(brand)
	data.ContentURL = utils.CreatePreSignedDownloadUrl(data.ContentURL)
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
	data, err := campaignService.UpdateCampaign(brand, mux.Vars(r)["id"])
	data.ContentURL = utils.CreatePreSignedDownloadUrl(data.ContentURL)
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
	data, err := campaignService.DeleteCampaign(mux.Vars(r)["id"])
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
	data, err := campaignService.GetAllCampaign()
	for i := range data {
		data[i].ContentURL = utils.CreatePreSignedDownloadUrl(data[i].ContentURL)
		data[i].VideoURL = utils.CreatePreSignedDownloadUrl(data[i].VideoURL)
	}
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}
func GetCount(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := campaignService.Count()

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
	data, err := campaignService.GetMyCampaign(dataToken["userid"].(string))
	for i := range data {
		data[i].ContentURL = utils.CreatePreSignedDownloadUrl(data[i].ContentURL)
		data[i].VideoURL = utils.CreatePreSignedDownloadUrl(data[i].VideoURL)
	}
	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}
func Find(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting brand", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := campaignService.GetMyCampaign(mux.Vars(r)["name"])
	for i := range data {
		data[i].ContentURL = utils.CreatePreSignedDownloadUrl(data[i].ContentURL)
		data[i].VideoURL = utils.CreatePreSignedDownloadUrl(data[i].VideoURL)
	}
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
	_, done2 := utils.CheckToken(w, r)
	if done2 {
		return
	}
	data, err := campaignService.GetCampaign(mux.Vars(r)["id"])
	data.ContentURL = utils.CreatePreSignedDownloadUrl(data.ContentURL)
	data.VideoURL = utils.CreatePreSignedDownloadUrl(data.VideoURL)

	handlers.Handler(w, err, data)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("brand updated", logrus.Fields{
		"duration": duration,
	})
}

func parseJSON(w http.ResponseWriter, r *http.Request) (*campaignEntity.Campaign, bool) {
	var brand *campaignEntity.Campaign
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
