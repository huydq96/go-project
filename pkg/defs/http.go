package defs

type (
	SuccessResponse struct {
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message,omitempty"`
	}
)
