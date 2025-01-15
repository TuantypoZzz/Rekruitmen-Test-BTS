package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type checklistServiceImpl struct {
	repository.ChecklistRepository
}

func NewChecklistServiceImpl(checklistRepository *repository.ChecklistRepository) ChecklistService {
	return &checklistServiceImpl{ChecklistRepository: *checklistRepository}
}

func (service *checklistServiceImpl) Create(ctx context.Context, checklistModel model.CreateChecklist, username string) model.ChecklistModel {
	validation.Validate(checklistModel)

	checklist := entity.Checklist{
		Id:       uuid.New(),
		Title:    checklistModel.Title,
		Username: username,
	}
	checklist = service.ChecklistRepository.Insert(ctx, checklist)

	return model.ChecklistModel{
		Id:       checklist.Id.String(),
		Title:    checklist.Title,
		Username: checklist.Username,
	}
}

func (service *checklistServiceImpl) Delete(ctx context.Context, id int, username string) {
	checklist := service.ChecklistRepository.FindById(ctx, id)
	if checklist.Username != username {
		panic(exception.NotFoundError{
			Message: "Checklist Not Found",
		})
	}
	service.ChecklistRepository.Delete(ctx, checklist)
}

func (service *checklistServiceImpl) FindById(ctx context.Context, id int) model.ChecklistModel {
	checklist := service.ChecklistRepository.FindById(ctx, id)
	// exception.PanicLogging(err)
	return model.ChecklistModel{
		Id:       checklist.Id.String(),
		Title:    checklist.Title,
		Username: checklist.Username,
	}
}

func (service *checklistServiceImpl) FindAll(ctx context.Context, username string) (responses []model.ChecklistModel) {
	checklists := service.ChecklistRepository.FindAll(ctx, username)

	if len(checklists) == 0 {
		return []model.ChecklistModel{}
	}

	for _, checklist := range checklists {
		responses = append(responses, model.ChecklistModel{
			Id:       checklist.Id.String(),
			Title:    checklist.Title,
			Username: checklist.Username,
		})
	}
	return responses
}
