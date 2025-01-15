package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type checklistServiceImpl struct {
	*gorm.DB
}

func NewChecklistRepositoryImpl(DB *gorm.DB) ChecklistRepository {
	return &checklistServiceImpl{DB: DB}
}

func (repository *checklistServiceImpl) Insert(ctx context.Context, checklist entity.Checklist) entity.Checklist {
	err := repository.DB.WithContext(ctx).Create(&checklist).Error
	exception.PanicLogging(err)
	return checklist
}

func (repository *checklistServiceImpl) Update(ctx context.Context, checklist entity.Checklist) entity.Checklist {
	err := repository.DB.WithContext(ctx).Where("id = ?", checklist.Id).Updates(&checklist).Error
	exception.PanicLogging(err)
	return checklist
}

func (repository *checklistServiceImpl) Delete(ctx context.Context, checklist entity.Checklist) {
	repository.DB.WithContext(ctx).Where("id = ?", checklist.Id).Delete(&checklist)
}

func (repository *checklistServiceImpl) FindById(ctx context.Context, id int) entity.Checklist {
	var checklist entity.Checklist
	result := repository.DB.WithContext(ctx).Where("id = ?", id).First(&checklist)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "checklist Not Found",
		})
	}
	return checklist
}

func (repository *checklistServiceImpl) FindAll(ctx context.Context, username string) []entity.Checklist {
	var checklists []entity.Checklist
	repository.DB.WithContext(ctx).Where("username = ?", username).Find(&checklists)
	return checklists
}
