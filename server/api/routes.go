package main

import "net/http"

// Route defines a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of Route objects
type Routes []Route

// Define our routes
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},
	Route{
		"ProjectIndex",
		"GET",
		"/api/v1/projects",
		ProjectIndex,
	},
	Route{
		"ProjectCreate",
		"POST",
		"/api/v1/projects",
		ProjectCreate,
	},
	Route{
		"ProjectEdit",
		"PUT",
		"/api/v1/projects/{projectId}",
		ProjectEdit,
	},
	Route{
		"ProjectShow",
		"GET",
		"/api/v1/projects/{projectId}",
		ProjectShow,
	},
	Route{
		"ProjectDelete",
		"DELETE",
		"/api/v1/projects/{projectId}",
		ProjectDelete,
	},
	Route{
		"ProjectJoin",
		"POST",
		"/api/v1/projects/{projectId}/join",
		ProjectJoin,
	},
	Route{
		"ProjectLike",
		"POST",
		"/api/v1/projects/{projectId}/like",
		ProjectLike,
	},
	Route{
		"UserShow",
		"GET",
		"/api/v1/user",
		UserShow,
	},
}
