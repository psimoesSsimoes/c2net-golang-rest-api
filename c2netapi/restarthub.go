// Package  c2netapi provides a rest api for c2net iot hub functions
package c2netapi

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type Hub struct {
	State int `json:"state"`
}

func RestartHub(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/bin/bash", "-c", "sudo killall java").Output()
	if err != nil {
		log.Info(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "cd /home/pi/C2NET/c2net-iot-hub && sudo ./setEnv.sh").Output()
	if err != nil {
		log.Info(err)
	}

}

func StopHub(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {
		stmt, _ := db.Prepare("UPDATE ON_OFF SET state = ? WHERE state = 1")
		_, err = stmt.Exec(0)
		if err != nil {
			log.Info(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert hub state in database"})
		} else {
			_, err := exec.Command("/bin/bash", "-c", "sudo killall java").Output()
			if err != nil {
				log.Info(err)
			}

			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted hubState Into the Database"})
		}
	}

}
func StartHub(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {
		stmt, _ := db.Prepare("UPDATE ON_OFF SET state = ? WHERE state = 0")
		_, err = stmt.Exec(1)
		if err != nil {
			log.Info(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert hub state in database"})
		} else {
			_, err = exec.Command("/bin/bash", "-c", "cd /home/pi/C2NET/c2net-iot-hub && sudo ./setEnv.sh").Output()
			if err != nil {
				log.Info(err)
			}

			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted hubState Into the Database"})
		}
	}

}


func StatusHub(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "/home/pi/C2NET/c2net-iot-hub/tables/c2net.db")
	defer db.Close()

	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Couldn't open c2net sqlite db"})
	} else {
	
		rows, err := db.Query("SELECT * FROM ON_OFF")

			var s Hub
		for rows.Next() {


			err = rows.Scan(&s.State)
			if err != nil {
				log.Error(err.Error())
				json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select an sensor from database"})
				return
			}
		}
		log.Info(s)
		json.NewEncoder(w).Encode(s)

	}

}

