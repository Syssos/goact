package chatroom

import (
	"reflect"
	"testing"
)

func TestChatroom(t *testing.T) {
	t.Run("Testing Chatroom returns empty Room w/ Start method", func(t *testing.T) {
		// This test is used to ensure that the start method doesn't get removed by mistake as it is 
		// responsible for the chatrooms users communication and is a vitial component in the app
		newroom := NewRoom()

		st := reflect.TypeOf(newroom)
		_, ok := st.MethodByName("Start")
		if !ok {
			t.Error("No start Method exists")
		}
	})
}