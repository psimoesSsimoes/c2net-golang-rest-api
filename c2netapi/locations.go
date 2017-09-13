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
 * Locations where sensors are going to be placed
 */
type SensorArea struct {
	Idc2net int    `json:"idc2net"`
	Idnode  string `json:"idnode"`
	Name    string `json:"name"`
}

/**
returns all sensor areas
*/
func AllSensorAreas(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		var areas []SensorArea

		rows, err := db.Query("SELECT * FROM locations")

		for rows.Next() {

			var area SensorArea
			err = rows.Scan(&area.Idc2net, &area.Idnode, &area.Name)
			if err != nil {
				log.Error(err.Error())
				json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select an area from database"})
			}
			areas = append(areas, area)
		}
		log.Info(areas)
		json.NewEncoder(w).Encode(areas)
	}
}

func InsertSensorArea(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		decoder := json.NewDecoder(r.Body)
		log.Info(decoder)
		var area SensorArea
		err = decoder.Decode(&area)

		if err != nil {
			log.Println(err.Error())
		}
		if ValidSensorArea(area) {
			stmt, _ := db.Prepare("INSERT INTO locations (cod_res_c2net,cod_res_node,name_location) values (?,?,?)")
			_, err = stmt.Exec(area.Idc2net, area.Idnode, area.Name)
			if err != nil {
				log.Info(err.Error()) // proper error handling instead of panic in your app
				json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
			} else {

				json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted SensorArea Into the Database", Body: fmt.Sprintf("%+v\n", area)})
			}
		}
	}
}

func EditArea(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		decoder := json.NewDecoder(r.Body)
		var area SensorArea
		err = decoder.Decode(&area)

		if err != nil {
			log.Error(err.Error())
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Received JSON invalid"})
			return
		}
		stmt, _ := db.Prepare("UPDATE locations SET cod_res_node = ? , name_location = ? where cod_res_c2net = ?")
		_, err = stmt.Exec(area.Idnode, area.Name, area.Idc2net)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't execute update"})
			return
		}

		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: fmt.Sprintf("Succefully updated SensorArea : %s", area.Idc2net)})

	}

}

func DeleteArea(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	vars := mux.Vars(r)
	idToDelete := vars["id"]

	stmt, _ := db.Prepare("DELETE FROM locations WHERE cod_res_c2net = ?")
	if err != nil {

		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
		return
	}

	_, err = stmt.Exec(idToDelete)

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
	} else {
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully deleted SensorArea from the Database"})
	}

}

func DeleteAllAreas(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/psimoes/Github/c2net-golang-rest-api/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		stmt, err := db.Prepare("DELETE FROM locations")
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't prepare delete on c2net sqlite db"})
			return
		}
		_, err = stmt.Exec()

		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't execute delete on c2net sqlite db"})
			return
		}

		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Succefully deleted all Sensor Areas"})

	}

}
