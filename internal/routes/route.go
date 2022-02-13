package routes

import (
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/handlers/commands"
	"github.com/yemekSepeti/internal/handlers/queries"
	"net/http"
	"regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Routes(api database.API) []route {
	return []route{
		newRoute("GET", "/api/keys/([^/]+)", queries.GetObjectByKey(api)),
		newRoute("GET", "/api/keys", queries.GetObjectList(api)),
		newRoute("POST", "/api/keys", commands.CreateObject(api)),
		newRoute("DELETE", "/api/keys", commands.FlushObjectList(api)),
	}
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}
