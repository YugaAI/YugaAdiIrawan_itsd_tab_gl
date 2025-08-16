package handlers

import (
	"interview/BackendAPIHumanCapital/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	GetEmployeeByNIK(nik string) (*models.Employee, error)
	UpdateEmployeeByName(nik string, newName string) (string, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

type ResponseMessage struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
type ResponseSuccessUpdate struct {
	Message string `json:"message"`
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/api/v1")
	route.GET("/info/:nik", h.GetEmployeByNIK)
	route.POST("/info/:nik", h.Update)
}
