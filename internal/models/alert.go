package models

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Alert struct {
		ID primitive.ObjectID `json:"id" bson:"_id"`
	}

	AlertReqFilter struct {
		Keyword string `json:"keyword" query:"keyword"`
	}
)

type AlertUseCase interface {
	GetAlertByOptions(ctx echo.Context, req AlertReqFilter) ([]Alert, error)
}

type AlertRepository interface {
	GetAlertByOptions(ctx echo.Context, req AlertReqFilter) ([]Alert, error)
}
