package router

import (
	handler "github.com/bootcamp-go/desafio-go-web/cmd/server/handlers"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	e *gin.Engine
	list []domain.Ticket

}

func NewRouter(r *gin.Engine, list []domain.Ticket) *Router {
	return &Router{
		e: r,
		list: list,
	}
}


func (r *Router) Setup() {
	r.e.Use(gin.Logger())
	r.e.Use(gin.Recovery())
}

func (r *Router) MapRoutes() {
	db := r.list
	repo := tickets.NewRepository(db)
	service := tickets.NewService(repo)
	handler := handler.NewService(service)

	gr := r.e.Group("/ticket")
	{
		gr.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
		gr.GET("/getAverage/:dest", handler.AverageDestination())
	}
	r.e.Run(":8080")
}
