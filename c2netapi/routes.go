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

	Route{"AllSensors", "GET", "/allsensors", AllSensors},
	Route{"NewSensor", "POST", "/sensor", InsertSensor},
	Route{"DeleteSensor", "DELETE", "/sensor/{id}", DeleteSensor},
	Route{"DeleteAllSensors", "DELETE", "/sensors", DeleteAllSensors},
	Route{"InsertHubId", "POST", "/hub", InsertHubId},
	Route{"AllListeners", "GET", "/comm", AllCommListeners},
	Route{"UpdateListeners", "POST", "/comm", UpdateCommListeners},
	Route{"AllSensingElements", "GET", "/allsensing", AllSensingElements},
	Route{"InsertSensingElements", "POST", "/insertsensings", InsertSensingElements},
	Route{"RestartHub", "GET", "/restart", RestartHub},
	Route{"InsertAlive", "POST", "/alive", InsertAlive},
	Route{"AllAlive", "GET", "/allalive", AllAlive},
}