package controllers

import (
	"github.com/alx-t/go-playground/models"
)

type (
	// For Get - /projects
	ProjectsResource struct {
		Data []models.Project `json:"projects"`
	}

	// For Post/Put - /projects
	// For Get - /projects/id
	ProjectResource struct {
		Data models.Project `json:"project"`
	}
)
