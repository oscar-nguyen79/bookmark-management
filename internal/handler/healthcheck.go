package handler

import (
	"net/http"

	"bookmark-management/internal/service"

	"github.com/gin-gonic/gin"
)

type HealthCheck interface {
	Check(c *gin.Context)
}

type healthCheckHandler struct {
	healthCheckService service.HealthCheck
}

func NewHealthCheck(healthCheckSvc service.HealthCheck) HealthCheck {
	return &healthCheckHandler{
		healthCheckService: healthCheckSvc,
	}
}

func (h *healthCheckHandler) Check(c *gin.Context) {
	response, err := h.healthCheckService.Check()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, response)
	}

	c.JSON(http.StatusOK, response)
}
