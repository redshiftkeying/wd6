package main

import "github.com/redshiftkeying/wd6/server/routers"

func main() {
	r := routers.Router
	SetServerOptions("")
	routers.SetRouters("development")
	r.Run()
}
