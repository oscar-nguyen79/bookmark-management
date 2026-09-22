package handler

import (
	"bookmark-management/internal/app/service"
	"net/http"

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

// @Summary API health check
// @Description API health check
// @Tags Health-check
// @Produce application/json
// @Success      200  {object}  map[string]string ""
// @Failure      500  {object}  string "Internal Server Error"
// @Router       /health-check [get]
func (h *healthCheckHandler) Check(c *gin.Context) {
	response, err := h.healthCheckService.Check()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"message": "Service unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
