package c2netapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

/*
 * sensor placed on an location
 */
type Sensor struct {
	Cat      string `json:"cat"`
	Code     string `json:"code"`
	Sensor   string `json:"sensor"`
	Selected int    `json:"selected"`
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
		rows, err := db.Query("SELECT * FROM types_of_sensors")

		for rows.Next() {

			var s Sensor
			err = rows.Scan(&s.Cat, &s.Code, &s.Sensor, &s.Selected)
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

/**
returns all selected
*/
func AllSelected(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		var sensors []Sensor
		log.Info("aRRIVED HERE")
		rows, err := db.Query("SELECT * FROM types_of_sensors where selected=1")

		for rows.Next() {

			var s Sensor
			err = rows.Scan(&s.Cat, &s.Code, &s.Sensor, &s.Selected)
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

func UpdateSelected(w http.ResponseWriter, r *http.Request) {
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
			stmt, _ := db.Prepare("UPDATE types_of_sensors SET selected = 1  WHERE code = (?)")
			_, err = stmt.Exec(v.Code)
			if err != nil {
				log.Info("entered error")
				log.Error(err.Error()) // proper error handling instead of panic in your app
				json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to update selected in database"})
			}
		}
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully updated selected", Body: fmt.Sprintf("%+v\n", sensors)})

	}
}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
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

		stmt, _ := db.Prepare("UPDATE types_of_sensors SET selected = 0  WHERE code = (?)")

		_, err = stmt.Exec(v.Code)
	}
	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully deleted Sensor from the Database"})

}

func DeleteAllSensors(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		stmt, err := db.Prepare("DELETE FROM sensors")
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
