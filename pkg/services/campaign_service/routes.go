package campaign_service

import (
	"kinexx_backend/pkg/entity"
	campaignHandler "kinexx_backend/pkg/services/campaign_service/handler"
)

var CampaignServices = entity.Routes{
	entity.Route{
		Name:        "Campaign",
		Method:      "POST",
		Pattern:     "/campaign",
		HandlerFunc: campaignHandler.Add,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "GET",
		Pattern:     "/campaign",
		HandlerFunc: campaignHandler.GetAll,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "GET",
		Pattern:     "/my/campaign",
		HandlerFunc: campaignHandler.GetMy,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "GET",
		Pattern:     "/find/campaign/{name}",
		HandlerFunc: campaignHandler.Find,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "GET",
		Pattern:     "/campaign/{id}",
		HandlerFunc: campaignHandler.GetDetail,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "GET",
		Pattern:     "/count/campaign",
		HandlerFunc: campaignHandler.GetCount,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "PUT",
		Pattern:     "/campaign/{id}",
		HandlerFunc: campaignHandler.Update,
	},
	entity.Route{
		Name:        "Campaign",
		Method:      "DELETE",
		Pattern:     "/campaign/{id}",
		HandlerFunc: campaignHandler.Delete,
	},
}
