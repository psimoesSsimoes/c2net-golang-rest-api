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
type Sensor struct {
	Idsensor string `json:"id"`
	Name     string `json:"name"`
	Tipo     string `json:"tipo"`
}

/**
returns all sensor areas
*/
func AllSensors(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		var sensors []Sensor
		log.Info("aRRIVED HERE")
		rows, err := db.Query("SELECT * FROM sensor")

		for rows.Next() {

			var s Sensor
			err = rows.Scan(&s.Idsensor, &s.Name, &s.Tipo)
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

func InsertSensor(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		decoder := json.NewDecoder(r.Body)
		log.Info(decoder)
		var sensors []Sensor
		err = decoder.Decode(&sensors)
		log.Info(len(sensors))
		if err != nil {
			log.Error(err.Error())
		}
		log.Info("passed decode")
		for _, v := range sensors {

			log.Info(v)
			stmt, _ := db.Prepare("INSERT INTO sensor(id,name,tipo) values(?,?,?)")
			_, err = stmt.Exec(v.Idsensor, v.Name, v.Tipo)
			if err != nil {
				log.Info("entered error")
				log.Error(err.Error()) // proper error handling instead of panic in your app
				json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
			}
		}
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted SensorArea Into the Database", Body: fmt.Sprintf("%+v\n", sensors)})

	}
}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	vars := mux.Vars(r)
	idToDelete := vars["id"]

	stmt, _ := db.Prepare("DELETE FROM sensor WHERE id = ?")
	if err != nil {

		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor in database"})
		return
	}

	_, err = stmt.Exec(idToDelete)

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor in database"})
	} else {
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully deleted Sensor from the Database"})
	}

}

func DeleteAllSensors(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		stmt, err := db.Prepare("DELETE FROM sensor")
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

		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Succefully deleted all Sensors"})

	}

}
