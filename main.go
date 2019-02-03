package main

import (
	"log"
	"net/http"

	common "github.com/alx-t/go-playground/common"
	"github.com/alx-t/go-playground/routes"
	"github.com/codegangsta/negroni"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}

	defer common.Db.Close()

	// recs, err := project.Read(common.Db, "")
	// if err != nil {
	// 	fatal(err)
	// }
	// log.Println(recs)

	// common.ExecuteCmd("ls -al")

	log.Println("Listening...")
	server.ListenAndServe()
}
