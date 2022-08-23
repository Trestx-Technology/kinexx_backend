package view_content_service

import (
	"kinexx_backend/pkg/entity"
	viewContentHandler "kinexx_backend/pkg/services/view_content_service/handler"
)

var ViewContentServices = entity.Routes{
	entity.Route{
		Name:        "View Content",
		Method:      "POST",
		Pattern:     "/content",
		HandlerFunc: viewContentHandler.Add,
	},
	entity.Route{
		Name:        "View Content",
		Method:      "GET",
		Pattern:     "/content/{campaign_id}",
		HandlerFunc: viewContentHandler.GetForCampaign,
	},
	entity.Route{
		Name:        "View Content",
		Method:      "PUT",
		Pattern:     "/content/{id}",
		HandlerFunc: viewContentHandler.Update,
	},
	entity.Route{
		Name:        "View Content",
		Method:      "DELETE",
		Pattern:     "/campaign/{id}",
		HandlerFunc: viewContentHandler.Delete,
	},
}
