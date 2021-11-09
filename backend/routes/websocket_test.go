package routes

import (
    // "net/http"
    // "net/http/httptest"
    "testing"

	// "github.com/gorilla/websocket"
)

func TestWebsocketValidation(t *testing.T) {
	// t.Run("Testing with no token", func(t *testing.T) {
	// 	request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
	// 	response := httptest.NewRecorder()
	// 	WebSock(response, request)
	// 	got := response.Code
	// 	want := 401
		
	// 	if got != want {
	// 		t.Errorf("got %v, want %v", got, want)
	// 	}
	// })
	// t.Run("Testing cookie with incorrect token", func(t *testing.T) {
	// 	request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
 //        response := httptest.NewRecorder()
	// 	cookie := http.Cookie{Name: "token", Value: "IncorrectValueForSure"}
	// 	request.AddCookie(&cookie)
        
	// 	WebSock(response, request)
 //        got := response.Code
 //        want := http.StatusBadRequest

 //        cookieCheck := response.Result().Cookies()
 //        if got != want {
 //            t.Errorf("got %v, want %v", got, want)
 //        }
 //        if len(cookieCheck) > 0 {
	// 		t.Errorf("Cookie present, potentual security issue, or forcing clients to re-validate")
	// 	}
	// })
	// t.Run("Testing cookie with expired token", func(t *testing.T) {
	// 	request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
 //        response := httptest.NewRecorder()
	// 	cookie := http.Cookie{Name: "token", Value: GenerateExpired()}
	// 	request.AddCookie(&cookie)
        
	// 	WebSock(response, request)
 //        got := response.Code
 //        want := http.StatusBadRequest

 //        cookieCheck := response.Result().Cookies()
 //        if got != want {
 //            t.Errorf("got %v, want %v", got, want)
 //        }
 //        if len(cookieCheck) > 0 {
	// 		t.Errorf("Cookie present, potentual security issue, or forcing clients to re-validate")
	// 	}
	// })
	return
}

