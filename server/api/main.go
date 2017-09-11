package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	flag "github.com/ogier/pflag"
	"stash.open.ch/hack/backend/src/ssomiddleware"
)

// Repo Singleton
var Repo *ProjectRepo

var port uint16
var dbpath string

func init() {
	flag.Uint16VarP(&port, "port", "p", 8080, "port to listen on")
	flag.StringVarP(&dbpath, "db", "d", "./hackathon.db", "Database path")

}

func main() {
	flag.Parse()
	router := NewRouter()
	Repo, _ = NewRepo(dbpath)
	defer Repo.Close()

	ssoMiddleware := ssomiddleware.New(ssomiddleware.Options{
		UsernameHeader:  "X-Auth-Username",
		EmailHeader:     "X-Auth-Email",
		GivenNameHeader: "X-Auth-GivenName",
		SurnameHeader:   "X-Auth-Surname",
		Debug:           false,
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.RecoveryHandler()(handlers.CompressHandler(handlers.CORS()(handlers.CombinedLoggingHandler(os.Stdout, ssoMiddleware.Handler(handlers.ProxyHeaders(router))))))))

}
