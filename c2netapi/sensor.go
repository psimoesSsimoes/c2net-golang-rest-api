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
	Nodeid   string `json:"nodeid"`
	Typeid   string `json:"typeid"`
	Typename string `json:"type_name"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Freq     string `json:"freq"`
}

var gsensors map[string]Sensor

/**
returns all sensor areas
*/
func AllSensors(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()
	

	var sensors []Sensor
	if gsensors==nil{
	rows, err := db.Query("SELECT * FROM sensors")
	gsensors=make(map[string]Sensor,0)
	for rows.Next() {

		var s Sensor
		err = rows.Scan(&s.Nodeid, &s.Typeid, &s.Typename, &s.Id, &s.Name, &s.Freq)
		if err != nil {
			log.Error(err.Error())
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select an sensor from database"})
			return
		}
		sensors = append(sensors, s)
		gsensors[s.Nodeid+s.Id]=s
	}
}else{
	for _,v := range gsensors{

		sensors = append(sensors, v)
	}
}
	log.Info(sensors)
	json.NewEncoder(w).Encode(sensors)
}

func InsertSensor(w http.ResponseWriter, r *http.Request) {

	if gsensors == nil {
		gsensors = make(map[string]Sensor, 0)
	}

	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {

		decoder := json.NewDecoder(r.Body)
		var sensors []Sensor
		err = decoder.Decode(&sensors)
		
		if err != nil {
			log.Error(err.Error())
		}
		defer db.Close()
		for _, v := range sensors {
			_, ok := gsensors[v.Nodeid+v.Id]
			
			if !ok {

				stmt, _ := db.Prepare("INSERT INTO sensors(nodeid,typeid,typename,id,name,freq) values(?,?,?,?,?,?)")
				_, err = stmt.Exec(v.Nodeid, v.Typeid, v.Typename, v.Id, v.Name, v.Freq)
				if err != nil {
					json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
					return
				} else {
					gsensors[v.Nodeid+v.Id] = v
				}

			} else {
				stmt, _ := db.Prepare("UPDATE sensors SET typeid = (?),typename = (?),name=(?),freq=(?) where nodeid = (?) and id =(?)")

				_, err = stmt.Exec(v.Typeid, v.Typename, v.Name, v.Freq, v.Nodeid, v.Id)
				if err == nil {
					gsensors[v.Nodeid+v.Id] = v
				}

			}
		}
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted SensorArea Into the Database", Body: fmt.Sprintf("%+v\n", sensors)})

	}
}
func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {
		decoder := json.NewDecoder(r.Body)
		var sensor Sensor
		err = decoder.Decode(&sensor)
		if err != nil {
			log.Error(err.Error())
		}
		stmt, _ := db.Prepare("DELETE FROM sensors WHERE nodeid = (?) and id = (?)")
		_, err = stmt.Exec(sensor.Nodeid, sensor.Id)
		if err != nil {
			log.Error(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert sensor area in database"})
		} else {
			delete(gsensors, sensor.Nodeid+sensor.Id)
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Deleted Sensor from the Database", Body: fmt.Sprintf("%+v\n", sensor)})
		}

	}
}
