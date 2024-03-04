package mongo

import (
	"github.com/labstack/echo/v4"
	"go-project/internal/models"
	"go-project/pkg/interfaces/database/mongodb"
)

type AlertRepository struct {
	collection mongodb.Collection
}

func NewAlertRepository(client mongodb.Client, dbName, collectionName string) *AlertRepository {
	alertCollection := client.Database(dbName).Collection(collectionName)
	return &AlertRepository{
		collection: alertCollection,
	}
}

func (ar *AlertRepository) GetAlertByOptions(ctx echo.Context, filter models.AlertReqFilter) ([]models.Alert, error) {
	return nil, nil
}
