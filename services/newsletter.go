package services

import "github.com/Linus-Boehm/go-serverless-suite/itf"

type newletterService struct {
}

func NewNewsletterService() itf.NewsWriter {
	return &newletterService{}
}
