package main

import (
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/config"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/router"
)

func init(){
	config.Load()
}

func main(){
	myApp := app.NewApp()

    r := router.SetupRoutes(myApp)

    if err := r.Run(":8080"); err != nil {
        panic(err)
    }

}