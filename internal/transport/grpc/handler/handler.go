package handler

import "github.com/sirupsen/logrus"

type Handler struct {
	LocationHandler *LocationHandler
}

func  New(logger *logrus.Logger) *Handler {
	var locationHandler LocationHandler

	return &Handler{
		LocationHandler: locationHandler.New(logger),
	}
}