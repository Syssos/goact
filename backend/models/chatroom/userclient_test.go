package chatroom

import (
	"log"
    "net/http"
    "net/http/httptest"
	"testing"
  
    "github.com/google/uuid"
	"github.com/gorilla/websocket"

    // "github.com/Syssos/goact/routes"
)

func TestClient(t *testing.T) {
	room := NewRoom()
	go room.Start()


	r, _ := http.NewRequest(http.MethodGet, "/ws", nil)
	w := httptest.NewRecorder()

	conn, upgradeErr := Upgrade(w, r)
	if upgradeErr != nil {
		t.Error("Error Upgrading socket")
	}


	client := &Client{
        ID:       uuid.New(),
        UserName: "TestUser",
        Conn:     conn,
        Room:     room,
    }

    room.Register <- client
    if len(room.Clients) > 1 {
    	t.Error("Client not Registered")
    }

    room.Unregister <- client
    if len(room.Clients) < 1 {
    	t.Error("Client not Registered")
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    return conn, nil
}