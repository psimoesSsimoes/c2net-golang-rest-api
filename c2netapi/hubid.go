// Package  c2netapi provides a rest api for c2net iot hub functions
package c2netapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type HubId struct {
	Id int `json:"id"`
}

func InsertHubId(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
		return
	}

	decoder := json.NewDecoder(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
		return
	}

	stmt, _ := db.Prepare("DELETE FROM hubid")
	if err != nil {
		log.Error(err)
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't prepare delete on c2net sqlite db"})
		return
	}
	_, err = stmt.Exec()

	if err != nil {
		log.Error(err)
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't execute delete hubid  sqlite db"})
		return
	}
	var hub HubId
	err = decoder.Decode(&hub)
	log.Info(hub)

	stmt, _ = db.Prepare("INSERT INTO hubid (id) values (?)")

	_, err = stmt.Exec(hub.Id)

	if err != nil {
		log.Info(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert id in database"})
	} else {

		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted SensorArea Into the Database", Body: fmt.Sprintf("%+v\n", hub)})
	}

}
