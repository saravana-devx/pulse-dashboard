package jobs

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateJobs(c *gin.Context) {
}

func (h *Handler) GetJobById(c *gin.Context) {
}

func (h *Handler) GetAllJobs(c *gin.Context) {
}

func (h *Handler) UpdateJob(c *gin.Context) {
}

func (h *Handler) DeleteJob(c *gin.Context) {
}
