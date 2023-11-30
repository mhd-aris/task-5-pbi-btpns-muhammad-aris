package main

import (
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/config"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/database"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/routers"
)

func init(){
	config.Load()
}

func main(){
	dbConfig := config.GetDatabaseConfig()
	db := database.NewDB(&dbConfig)
	
    api := app.New(db)
	
	routers.InitRoutes(api)

    api.Router.Run(":8080")

}