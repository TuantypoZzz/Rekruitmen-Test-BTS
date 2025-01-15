package service

import (
	"context"
	"golang-todo-app/model"
)

type ChecklistService interface {
	Create(ctx context.Context, checklistModel model.CreateChecklist, username string) model.ChecklistModel
	// UpdateStatus(ctx context.Context, ChecklistModel model.UpdateChecklistStatus, id string, username string) model.ChecklistModel
	// Update(ctx context.Context, productModel model.ChecklistCreateOrUpdateModel, id int) model.ChecklistCreateOrUpdateModel
	Delete(ctx context.Context, id int, username string)
	FindById(ctx context.Context, id int) model.ChecklistModel
	FindAll(ctx context.Context, username string) []model.ChecklistModel
}
