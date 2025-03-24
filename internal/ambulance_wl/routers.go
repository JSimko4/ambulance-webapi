/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xsimkoj2@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the AmbulanceConditionsAPI part of the API
	AmbulanceConditionsAPI AmbulanceConditionsAPI
	// Routes for the AmbulanceWaitingListAPI part of the API
	AmbulanceWaitingListAPI AmbulanceWaitingListAPI
	// Routes for the AmbulancesAPI part of the API
	AmbulancesAPI AmbulancesAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"GetConditions",
			http.MethodGet,
			"/api/waiting-list/:ambulanceId/condition",
			handleFunctions.AmbulanceConditionsAPI.GetConditions,
		},
		{
			"CreateWaitingListEntry",
			http.MethodPost,
			"/api/waiting-list/:ambulanceId/entries",
			handleFunctions.AmbulanceWaitingListAPI.CreateWaitingListEntry,
		},
		{
			"DeleteWaitingListEntry",
			http.MethodDelete,
			"/api/waiting-list/:ambulanceId/entries/:entryId",
			handleFunctions.AmbulanceWaitingListAPI.DeleteWaitingListEntry,
		},
		{
			"GetWaitingListEntries",
			http.MethodGet,
			"/api/waiting-list/:ambulanceId/entries",
			handleFunctions.AmbulanceWaitingListAPI.GetWaitingListEntries,
		},
		{
			"GetWaitingListEntry",
			http.MethodGet,
			"/api/waiting-list/:ambulanceId/entries/:entryId",
			handleFunctions.AmbulanceWaitingListAPI.GetWaitingListEntry,
		},
		{
			"UpdateWaitingListEntry",
			http.MethodPut,
			"/api/waiting-list/:ambulanceId/entries/:entryId",
			handleFunctions.AmbulanceWaitingListAPI.UpdateWaitingListEntry,
		},
		{
			"CreateAmbulance",
			http.MethodPost,
			"/api/ambulance",
			handleFunctions.AmbulancesAPI.CreateAmbulance,
		},
		{
			"DeleteAmbulance",
			http.MethodDelete,
			"/api/ambulance/:ambulanceId",
			handleFunctions.AmbulancesAPI.DeleteAmbulance,
		},
	}
}
