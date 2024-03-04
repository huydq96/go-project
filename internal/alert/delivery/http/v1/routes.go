package v1

func (h *AlertHandler) MapRoutes() {
	h.group.GET("", h.GetAlerts)
}
