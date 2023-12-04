package main

import (
	"net/http"

	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/config"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/router"
)

func main() {
	config.Load()
	r := router.Init()
	r.StaticFS("/uploads", http.Dir("uploads"))
	r.Run(":8080")
}
