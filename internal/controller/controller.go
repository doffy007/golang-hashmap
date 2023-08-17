package controller

import (
	"sync"
)

type ShortenURLController struct {
	mtx    sync.Mutex
	urlMap map[string]string
}

func NewShortenURLController() *ShortenURLController {
	return &ShortenURLController{
		urlMap: map[string]string{},
	}
}
