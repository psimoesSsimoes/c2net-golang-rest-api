package c2netapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

/*
 * sensor placed on an location
 */
type Alive struct {
	Idsensor string `json:"idsensor"`
	Idc2net  string `json:"idc2net"`
	Name     string `json:"name"`
}

/**
returns all sensor areas
*/
func AllAlive(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		var sensors []Alive
		log.Info("aRRIVED HERE")
		rows, err := db.Query("SELECT * FROM alive_sensors")

		for rows.Next() {

			var s Alive
			err = rows.Scan(&s.Idsensor, &s.Idc2net, &s.Name)
			if err != nil {
				log.Error(err.Error())
				json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select an sensor from database"})
				return
			}
			sensors = append(sensors, s)
		}
		log.Info(sensors)
		json.NewEncoder(w).Encode(sensors)
	}
}

func InsertAlive(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		decoder := json.NewDecoder(r.Body)
		log.Info(decoder)
		var s Alive
		err = decoder.Decode(&s)

		if err != nil {
			log.Println(err.Error())
		}
		stmt, _ := db.Prepare("INSERT INTO alive_sensors (cod_sensor,cod_res_node,name_sensor) values (?,?,?)")
		_, err = stmt.Exec(s.Idsensor, s.Idc2net, s.Name)
		if err != nil {
			log.Info(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
		} else {

			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted AliveArea Into the Database", Body: fmt.Sprintf("%+v\n", s)})
		}
	}
}

func DeleteAlive(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	vars := mux.Vars(r)
	idToDelete := vars["id"]

	stmt, _ := db.Prepare("DELETE FROM alive_sensors WHERE cod_res_node = ?")
	if err != nil {

		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor in database"})
		return
	}

	_, err = stmt.Exec(idToDelete)

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor in database"})
	} else {
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully deleted Alive from the Database"})
	}

}

func AllAlivesPerArea(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	vars := mux.Vars(r)
	idToSearch := vars["id"]
	var sensors []Alive
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM alive_sensors where cod_res_node = %s", idToSearch))

	for rows.Next() {

		var s Alive
		err = rows.Scan(&s.Idsensor, &s.Idc2net, &s.Name)
		if err != nil {
			log.Error(err.Error())
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select an sensor from database"})
			return
		}
		sensors = append(sensors, s)
	}
	log.Info(sensors)
	json.NewEncoder(w).Encode(sensors)

}
