package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/MarkGibbons/chefapi_lib"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type restInfo struct {
	Cert string
	Key  string
	Port string
}

var flags restInfo

func main() {
	flagInit()

	r := mux.NewRouter()
	r.HandleFunc("/orgs", orgs)
	l, err := net.Listen("tcp4", ":"+flags.Port)
	if err != nil {
		panic(err.Error)
	}
	log.Fatal(http.ServeTLS(l, r, flags.Cert, flags.Key))
}

// orgs will send an array of the organizations found on the chef server
func orgs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orgNames, err := chefapi_lib.AllOrgs()
	if err != nil {
		msg, code := chefapi_lib.ChefStatus(err)
                http.Error(w, msg, code)
		return
	}
	orgJson, err := json.Marshal(orgNames)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message": "%+v"}`, err)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(orgJson)
	return
}

func flagInit() {
	restcert := flag.String("restcert", "", "Rest Certificate File")
	restkey := flag.String("restkey", "", "Rest Key File")
	restport := flag.String("restport", "8111", "Rest interface https port")
	flag.Parse()
	flags.Cert = *restcert
	flags.Key = *restkey
	flags.Port = *restport
	return
}
