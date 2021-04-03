package main

import (
	"github.com/kieranlavelle/vita-intellectus/pkg/api"
)

func main() {
	// habitRouter := api.CreateRoutes()
	taskRouter, pool := api.TaskRouter()
	defer pool.Close()

	// go log.Fatal(http.ListenAndServe("0.0.0.0:8004", habitRouter))
	taskRouter.Run(":8004")
}
