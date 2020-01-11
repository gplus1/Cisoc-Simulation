package main

import (
	"net/http"

	"github.com/betsegawlemma/restaurant-rest/FeedBack/repository"
	"github.com/betsegawlemma/restaurant-rest/FeedBack/service"
	"github.com/betsegawlemma/restaurant-rest/delivery/http/handler"
	urepim "github.com/betsegawlemma/restaurant-rest/user/repository"
	usrvim "github.com/betsegawlemma/restaurant-rest/user/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	roleRepo := urepim.NewRoleGormRepo(dbconn)
	roleSrv := usrvim.NewRoleService(roleRepo)
	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)

	FeedBackRepo := repository.FeedBackGormRepo(dbconn)
	FeedBackrv := service.NewFeedBackervice(FeedBackRepo)
	adminFeedBackHandler := handler.NewAdminFeedBackHandler(FeedBackrv)

	router := httprouter.New()

	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)

	router.GET("/v1/admin/FeedBack/:id", adminFeedBackHandler.GetSingleFeedBack)
	router.GET("/v1/admin/FeedBack", adminFeedBackHandler.GetFeedBack)
	router.PUT("/v1/admin/FeedBack/:id", adminFeedBackHandler.PutFeedBack)
	router.POST("/v1/admin/FeedBack", adminFeedBackHandler.PostFeedBack)
	router.DELETE("/v1/admin/FeedBack/:id", adminFeedBackHandler.DeleteFeedBack)

	http.ListenAndServe(":8181", router)
}
