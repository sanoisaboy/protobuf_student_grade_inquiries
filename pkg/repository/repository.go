package repository

import (
	"context"
)

type Repository interface {
	ListStudent(ctx context.Context) ([]User, error)
	CreateStudent(ctx context.Context, name string, id int, point int) (int, error)
	GetStudent(ctx context.Context, id int) (User, error)
	UpdateStudent(ctx context.Context, id int, point int) (int, []User, error)
	DeleteStudent(ctx context.Context, id int) (int, error)
}
