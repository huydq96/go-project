package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-project/config"
	"go-project/internal/models"
	"go-project/pkg/defs"
)

type AlertHandler struct {
	cfg          *config.Config
	group        *echo.Group
	alertUseCase models.AlertUseCase
}

func NewAlertHandler(cfg *config.Config, group *echo.Group, alertUseCase models.AlertUseCase) *AlertHandler {
	return &AlertHandler{
		cfg:          cfg,
		group:        group,
		alertUseCase: alertUseCase,
	}
}

func (h *AlertHandler) GetAlerts(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, defs.SuccessResponse{
		Data:    "sample data",
		Message: "ok",
	})
}
