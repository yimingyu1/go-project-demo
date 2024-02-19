package list

import (
	"awesomeProject1/internal/list/service"
	"awesomeProject1/pkg/app/api"
	"github.com/gin-gonic/gin"
	"log"
)

func Create(g *gin.Context) {
	var createTodoRequest CreateTodoRequest
	if err := g.ShouldBindJSON(&createTodoRequest); err != nil {
		log.Printf("createTodo bind params err: %v", err)
		api.SendResponse(g, err, nil)
		return
	}
	err := service.Svc.ListSvc.Create(g, createTodoRequest.Title)
	if err != nil {
		panic(err)
		return
	}
	api.SendResponse(g, nil, "创建成功")
}
