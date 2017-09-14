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
 * Locations joined with Alive sensors
 */
type SensingElement struct {
	Idc2net      int    `json:"idc2net"`
	Idnode       string `json:"idnode"`
	LocationName string `json:"loc_name"`
	IdSensor     string `json:"id_sensor"`
	SensorName   string `json:"sensor_name"`
}

/**
returns all sensing_elements areas
*/
func AllSensingElements(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		var areas []SensingElement

		log.Info("bam")
		rows, err := db.Query("select locations.cod_res_c2net,locations.cod_res_node,locations.name_location, alive_sensors.cod_sensor,alive_sensors.name_sensor from locations inner join alive_sensors on locations.cod_res_node = alive_sensors.cod_res_node")

		for rows.Next() {

			var area SensingElement
			err = rows.Scan(&area.Idc2net, &area.Idnode, &area.LocationName, &area.IdSensor, &area.SensorName)
			if err != nil {
				log.Error(err.Error())
				json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select an sensing element from database"})
			}
			areas = append(areas, area)
		}
		log.Info("bam")
		log.Info(areas)
		json.NewEncoder(w).Encode(areas)
	}
}

func InsertSensingElements(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		decoder := json.NewDecoder(r.Body)
		log.Info(decoder)
		var areas []SensingElement
		err = decoder.Decode(&areas)

		if err != nil {
			log.Println(err.Error())
		}

		for _, area := range areas {
			//check if exists
			if area.Idnode == "" || area.LocationName == "" {

			} else {
				stmt, _ := db.Prepare("INSERT INTO locations (cod_res_c2net,cod_res_node,name_location) values (?,?,?)")
				_, err = stmt.Exec(area.Idc2net, area.Idnode, area.LocationName)
				//need to handle error of primary key repeated
				if err != nil {
					log.Info(err.Error()) // proper error handling instead of panic in your app
					json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to insert an sensing element from database"})
				}
			}
			if area.IdSensor == "" || area.Idnode == "" || area.SensorName == "" {
			} else {
				stmt, _ := db.Prepare("INSERT INTO alive_sensors (cod_sensor,cod_res_node,name_sensor) values (?,?,?)")
				_, err = stmt.Exec(area.IdSensor, area.Idnode, area.SensorName)
				if err != nil {
					log.Info(err.Error()) // proper error handling instead of panic in your app
					json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to insert an sensing element from database"})
				}
				json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted SensingElement Into the Database", Body: fmt.Sprintf("%+v\n", area)})
			}
		}
	}
}
