package controller

import "net/http"

type CardController interface {
	GetCard(resp http.ResponseWriter, req *http.Request)
}
