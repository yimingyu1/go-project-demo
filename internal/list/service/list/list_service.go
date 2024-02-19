package list

import (
	"awesomeProject1/internal/list/service/model"
	"awesomeProject1/internal/list/service/repository/list"
	"context"
)

type ListService struct {
	listRepo *list.ListRepo
}

func NewListService() *ListService {
	db := model.GetDB()
	return &ListService{
		list.NewListRepo(db),
	}
}

func (srv *ListService) Create(ctx context.Context, todoContext string) error {

	listModel := &model.ListModel{
		TodoContext: todoContext,
	}
	return srv.listRepo.CreateTodo(ctx, listModel)
}

func (srv *ListService) GetList(ctx context.Context) []*model.ListModel {
	return srv.listRepo.GetList(ctx)
}

func (srv *ListService) UpdateDone(ctx context.Context, id uint64) {
	srv.listRepo.UpdateDone(ctx, id)
}

func (srv *ListService) DeletedTodo(ctx context.Context, id uint64) {
	listModel := &model.ListModel{
		ID: id,
	}
	srv.listRepo.DeletedTodo(ctx, listModel)
}
