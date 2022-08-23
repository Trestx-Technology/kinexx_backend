package view_service

import (
	"kinexx_backend/pkg/entity"
	viewHandler "kinexx_backend/pkg/services/view_service/handler"
)

var (
	ViewService = entity.Routes{
		entity.Route{
			Name:        "View",
			Method:      "POST",
			Pattern:     "/view",
			HandlerFunc: viewHandler.Add,
		},
		entity.Route{
			Name:        "View",
			Method:      "GET",
			Pattern:     "/view",
			HandlerFunc: viewHandler.GetAll,
		},
		entity.Route{
			Name:        "View",
			Method:      "GET",
			Pattern:     "/my/view",
			HandlerFunc: viewHandler.GetMy,
		},
		entity.Route{
			Name:        "View",
			Method:      "GET",
			Pattern:     "/view/{id}",
			HandlerFunc: viewHandler.GetDetail,
		},
		entity.Route{
			Name:        "View",
			Method:      "PUT",
			Pattern:     "/view/{id}",
			HandlerFunc: viewHandler.Update,
		},
		entity.Route{
			Name:        "View",
			Method:      "DELETE",
			Pattern:     "/view/{id}",
			HandlerFunc: viewHandler.Delete,
		},
	}
)
