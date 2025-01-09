package main

import (
	"github.com/anonydev/e-commerce-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	// listen and serve on 0.0.0.0:8002 (for windows "localhost:8002")
	r.Run(":8002")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run()
}
