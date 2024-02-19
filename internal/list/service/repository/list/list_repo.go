package list

import (
	"awesomeProject1/internal/list/service/model"
	"context"
	"github.com/jinzhu/gorm"
	"log"
)

type ListRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) *ListRepo {
	return &ListRepo{
		db: db,
	}
}

func (l *ListRepo) GetList(c context.Context) []*model.ListModel {
	var listModels []*model.ListModel
	l.db.Debug().Find(&listModels)
	return listModels
}

func (l *ListRepo) UpdateDone(c context.Context, id uint64) {
	listMode := &model.ListModel{
		ID: id,
	}
	l.db.Debug().Model(listMode).Update("todo_status", 1)
}

func (l *ListRepo) DeletedTodo(c context.Context, listModel *model.ListModel) {
	l.db.Debug().Delete(listModel)
}

func (l *ListRepo) CreateTodo(c context.Context, listModel *model.ListModel) error {
	err := l.db.Debug().Create(listModel).Error
	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
