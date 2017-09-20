// Package  c2netapi provides a rest api for c2net iot hub functions
package c2netapi

import (
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

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
 
