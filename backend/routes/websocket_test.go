package routes

import (
	"fmt"
	"log"
    "net/url"
    "net/http"
    "net/http/httptest"
	"net/http/cookiejar"
    "testing"

	"github.com/gorilla/websocket"
)

func TestWebsocketValidation(t *testing.T) {
	t.Run("Testing with no token", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
		response := httptest.NewRecorder()
		WebSock(response, request)
		got := response.Code
		want := 401
		
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Testing cookie with incorrect token", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        response := httptest.NewRecorder()
		cookie := http.Cookie{Name: "token", Value: "IncorrectValueForSure"}
		request.AddCookie(&cookie)
        
		WebSock(response, request)
        got := response.Code
        want := http.StatusBadRequest

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) > 0 {
			t.Errorf("Cookie present, potentual security issue, or forcing clients to re-validate")
		}
	})
	t.Run("Testing cookie with expired token", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        response := httptest.NewRecorder()
		cookie := http.Cookie{Name: "token", Value: GenerateExpired()}
		request.AddCookie(&cookie)
        
		WebSock(response, request)
        got := response.Code
        want := http.StatusBadRequest

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) > 0 {
			t.Errorf("Cookie present, potentual security issue, or forcing clients to re-validate")
		}
	})
}

func TestWebsocketConnection(t *testing.T) {
	// The intent for this test will for using it to ensure that the websocket endpoint can take a normal route, upgrade it, and allow for communication
	// Currently this test is not working and further research needs to be done into how this route can be test, the endpoint itself may need to be broken up
	// for easier testing purposes.

	// Generate dialer to handle ws after authentication
	d := websocket.Dialer{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
	}

	// set cookies to use
	cookies := []*http.Cookie{&http.Cookie{Name: "token", Value: GenerateTokenString()}}
	u, err := url.Parse("http://localhost:8080/ws")
	if err != nil {
		log.Fatal(err)
	}
	jar.SetCookies(u, cookies)
	d.Jar = jar

	// Dial websocket endpoint
	fmt.Println(d.Jar)
	c, resp, err := d.Dial("ws://" + "localhost:8080" + "/ws", nil)

	if err != nil {
		t.Fatal(err) // Current point of failure, ws fails to connect, I beleive it is due to the route requiring authentication.
	}
	
	if got, want := resp.StatusCode, http.StatusSwitchingProtocols; got != want {
		t.Errorf("resp.StatusCode = %q, want %q", got, want)
	}
	
	err = c.WriteJSON("test")
	if err != nil {
		t.Fatal(err)
	}
}