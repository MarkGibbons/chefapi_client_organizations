package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/MarkGibbons/chefapi_client"
	"github.com/gorilla/mux"
	"log"
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
	fmt.Printf("FLAGS %+v\n", flags)
	log.Fatal(http.ListenAndServe(":"+flags.Port, r))
}

// orgs will send an array of the organizations found on the chef server
func orgs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := chefapi_client.Client()
	orgList, err := client.Organizations.List()
	fmt.Printf("ORGLIST %+v ERR %+v", orgList, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message": "%+v"}`, err)
		w.Write([]byte(msg))
		return
	}
	orgNames := make([]string, 0, len(orgList))
	for k := range orgList {
		orgNames = append(orgNames, k)
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
