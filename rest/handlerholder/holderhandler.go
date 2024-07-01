package handlers

import "yet-again-templates/rest/domain"

func NewHandlerHolder(repo domain.Repository) *HandlerHolder {
	return &HandlerHolder{
		repo: repo,
	}
}

type HandlerHolder struct {
	repo domain.Repository
}
