package handler

import (
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/service"
	"net/http"
)

type TransformHandlers struct {
	service service.TransformInterface
}

func NewTransformHandlers(service service.TransformInterface) *TransformHandlers {
	return &TransformHandlers{service: service}
}

func (h *TransformHandlers) Transform(r *ginext.Request) (*ginext.Response, error) {
	return ginext.NewResponseData(http.StatusOK, nil), nil
}
