package http

import (
	"net/http"
)

type HttpConnection struct {
	Client http.Client
}

func NewHttpConnection() HttpConnection {
	client := http.Client{}

	return HttpConnection{Client: client}
}
