package repository

import (
	"context"

	"github.com/91diego/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var implementation Repository

// Set dependencies
func setRepository(repository Repository) {
	implementation = repository
}

// GetStudent
func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

// SetStudent
func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}
