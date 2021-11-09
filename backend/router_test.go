package main

import (
	// "net/http"
	// "net/http/httptest"
	"testing"
)

func TestRouterCreation(t *testing.T) {
	// t.Run("Testing router generation", func(t *testing.T) {
	// 	tt := []struct {
	// 		routeToHandle string
	// 		wantedValue int
	// 	}{
	// 		{"/", 200}, //status ok, homepage has index file present
	// 		{"/refresh", 401}, //status unautherized, requires token
	// 		{"/validate", 405}, //Method not allowed, validate only accepts posts requests
	// 		{"/ws", 401}, //status unautherized, requires token
	// 	}

	// 	for _, tc := range tt {
	// 		rr := httptest.NewRecorder()
	// 		req, err := http.NewRequest("GET", tc.routeToHandle, nil)
	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}

	// 		router := CreateRouter()
	// 		router.ServeHTTP(rr, req)

	// 		if rr.Code != tc.wantedValue {
	// 			t.Errorf("handler for %v should return: %v, got: %v", tc.routeToHandle, http.StatusOK, rr.Code)
	// 		}
	// 	}
	// })
}