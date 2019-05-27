package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var mysqlSchema = `
CREATE TABLE text_data(
	id int NOT NULL AUTO_INCREMENT,
	value text,
	PRIMARY KEY (id)
);`

var postgresSchema = `
CREATE TABLE text_data(
	id SERIAL PRIMARY KEY,
	value text
);`

var db *sqlx.DB

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

	var objmap map[string]interface{}

	var err error

	if err = json.Unmarshal([]byte(os.Getenv("DB_ADMIN")), &objmap); err != nil {
		panic(fmt.Sprintf("Unable to parse DB_ADMIN data - reason: %v", err))
	}

	db, err = sqlx.Connect(
		os.Getenv("DB_ENGINE"),
		fmt.Sprintf(
			"user=%s dbname=%s sslmode=require password=%s host=%s",
			objmap["username"].(string),
			os.Getenv("DB_NAME"),
			objmap["password"].(string),
			os.Getenv("DB_ENDPOINT"),
		),
	)

	if err != nil {
		panic(fmt.Sprintf("Unable to open db connection - reason: %v", err))
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Info("Backend called")

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello World - new!\n")

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			io.WriteString(w, pair[0]+"="+pair[1]+"\n")
		}

	})
	http.ListenAndServe(":7080", nil)
}
