package c2netapi

/*
 * Resp... the typical struct used to send back json responses
 */
type HttpResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func ValidSensorArea(area SensorArea) bool {
	if area.Idc2net < 1 {
		return false
	}
	if area.Idnode == "" || area.Name == "" {
		return false
	}
	return true
}
