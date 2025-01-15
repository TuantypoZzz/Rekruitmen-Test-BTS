package repository

import (
	"context"
	"golang-todo-app/entity"
)

type ChecklistRepository interface {
	Insert(ctx context.Context, checklist entity.Checklist) entity.Checklist
	Update(ctx context.Context, checklist entity.Checklist) entity.Checklist
	Delete(ctx context.Context, checklist entity.Checklist)
	FindById(ctx context.Context, id int) entity.Checklist
	FindAll(ctx context.Context, username string) []entity.Checklist
}
