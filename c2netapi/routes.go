package c2netapi

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	// All Tags Routes
	Route{"AllAreas", "GET", "/allareas", AllSensorAreas},
	Route{"NewArea", "POST", "/area", InsertSensorArea},
	Route{"EditArea", "PUT", "/area", EditArea},
	Route{"DeleteArea", "DELETE", "/area/{id}", DeleteArea},
	Route{"DeleteAllAreas", "DELETE", "/areas", DeleteAllAreas},
	Route{"Logger", "GET", "/log", Logger},

	Route{"AllSensors", "GET", "/allsensors", AllSensors},
	Route{"NewSensor", "POST", "/sensor", InsertSensor},
	Route{"DelSensor", "DELETE", "/sensor", DeleteSensor},
	Route{"UpdateTYPEsensor", "POST", "/update", UpdateSelected},
	Route{"AllSelected", "GET", "/allselected", AllSelected},
	Route{"AllCategories", "GET", "/allcategories", AllCategories},
	Route{"AllTypesSensors", "GET", "/alltypes", AllTypes},
	Route{"InsertHubId", "POST", "/hub", InsertHubId},
	Route{"AllListeners", "GET", "/comm", AllCommListeners},
	Route{"UpdateListeners", "POST", "/comm", UpdateCommListeners},
	Route{"RestartHub", "GET", "/restart", RestartHub},
	Route{"InsertAlive", "POST", "/alive", InsertAlive},
	Route{"AllAlive", "GET", "/allalive", AllAlive},
	Route{"HubStop", "GET", "/hubstop", StopHub},
	Route{"HubStart", "GET", "/hubstart", StartHub},
	Route{"HubStatus", "GET", "/hubstatus", StatusHub},
}
