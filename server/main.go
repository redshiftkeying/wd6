package main

import "github.com/redshiftkeying/zen-go/routers"

func main() {
	r := routers.Router

	routers.SetRouters("development")
	r.Run()
}
