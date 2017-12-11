package c2netapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CommunicationListeners struct {
	CommunicationListner string
}

type Comm struct {
	Name string `json:"name,omitempty"`
}

func ReadDataListeners() ([]string, bool) {
	dat, err := ioutil.ReadFile("/opt/c2net-iot-hub/Resources/configs.xml")
	if err != nil {
		return nil, false
	}
	split := strings.Split(fmt.Sprintf("%s", dat), "\n")

	dtl := make([]string, 0)

	for i := 2; i < len(split)-1; i++ {
		x := strings.Replace(strings.Replace(split[i], "<CommunicationListner>", "", -1), "</CommunicationListner>", "", -1)
		if x != "" {
			dtl = append(dtl, x)
		}
	}
	return dtl, true

}

func WriteDataListeners(newinput []string) bool {
	d1 := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<CommunicationListeners>\n"
	for i := 0; i < len(newinput); i++ {
		s := "	<CommunicationListner>" + newinput[i] + "</CommunicationListner>"
		d1 += s
		d1 += "\n"
	}
	d1 += "</CommunicationListeners>"
	err := ioutil.WriteFile("/opt/c2net-iot-hub/Resources/configs.xml", []byte(d1), 0644)
	if err != nil {
		return false
	}
	return true

}

func AllCommListeners(w http.ResponseWriter, r *http.Request) {
	dtl, flag := ReadDataListeners()
	if flag != true {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to read configs file"})
		return
	}
	json.NewEncoder(w).Encode(dtl)
}

func UpdateCommListeners(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var conn []Comm
	err := decoder.Decode(&conn)
	if err != nil {
		log.Fatal(err)
	}
	connNames := []string{}
	for _, v := range conn {
		connNames = append(connNames, v.Name)
	}

	flag := WriteDataListeners(connNames)
	if flag != true {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to write to configs file"})
		return
	}

	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Success!! data listeners updated"})
}
