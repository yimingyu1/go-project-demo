package service

import "awesomeProject1/internal/list/service/list"

var Svc *Service

type Service struct {
	ListSvc *list.ListService
}

func Init() {
	Svc = initSvc()
}

func initSvc() *Service {
	return &Service{
		ListSvc: list.NewListService(),
	}
}
