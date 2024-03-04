package usecase

import (
	"github.com/labstack/echo/v4"
	"go-project/config"
	"go-project/internal/models"
)

type AlertUseCase struct {
	cfg  *config.Config
	repo models.AlertRepository
}

func NewAlertUseCase(cfg *config.Config, repo models.AlertRepository) *AlertUseCase {
	return &AlertUseCase{
		cfg:  cfg,
		repo: repo,
	}
}

func (a *AlertUseCase) GetAlertByOptions(ctx echo.Context, req models.AlertReqFilter) ([]models.Alert, error) {
	return a.repo.GetAlertByOptions(ctx, req)
}
