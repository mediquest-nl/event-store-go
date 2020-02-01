// Package p contains an HTTP Cloud Function.
//package main  // for compiler test of code only
package p

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	//"html"
	"net/http"
	"strings"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func StoreEvent(w http.ResponseWriter, r *http.Request) {

	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		"mediquest-sandbox:europe-west1:cvk-event-store-2",
		"EventStore",
		"postgres",
		"Pgsql4cloudFunctions!")
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		fmt.Fprint(w, "Error while opening database")
	}
	defer db.Close()

	var reqStr = fmt.Sprint(r)
	reqStr = strings.Replace(reqStr, "'", "\\'", -1) // escape single quotes in string (we should also escape backslashes
	var qry string = fmt.Sprint("INSERT INTO Events (EventMsg) VALUES ('", reqStr, "')")
	_, err = db.Exec(qry)

	w.Header().Add("Content-Type", "text/html")
	var body string = fmt.Sprint("<html><head>TEST</head><body>Hello World<p>", err, "</p></body></html>")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(body))
	return
}

/*
func main() {

}
*/
