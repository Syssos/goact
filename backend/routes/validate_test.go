package routes

import (
    "net/http"
    "strings"
    "net/http/httptest"
    "testing"
    "encoding/json"
)

func TestCredentuals(t *testing.T) {
    t.Run("Testing with no credentuals", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "", Password: ""})))
        response := httptest.NewRecorder()
        ValidateUser(response, request)
        got := response.Code
        want := 401

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) != 0 {
            t.Errorf("Cookies set and shouldn't be %v", cookieCheck)
        }
    })
    t.Run("Testing with incorrect credentuals", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "ThistWrong", Password: "invalid"})))
        response := httptest.NewRecorder()
        ValidateUser(response, request)
        got := response.Code
        want := 401

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) != 0 {
            t.Errorf("Cookies set and shouldn't be %v", cookieCheck)
        }
    })
    t.Run("Testing with special characters in credentuals", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "&^(*%*&%@$*)*&@#$(", Password: "!@#$%^&*()_+{}|:\"<>?"})))
        response := httptest.NewRecorder()
        ValidateUser(response, request)
        got := response.Code
        want := 401

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) != 0 {
            t.Errorf("Cookies set and shouldn't be %v", cookieCheck)
        }
    })
    t.Run("Testing with correct credentuals", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "TestUser1", Password: "SomeTestpwd"})))
        response := httptest.NewRecorder()
        ValidateUser(response, request)
        got := response.Code
        want := 200

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) > 0 {
            if cookieCheck[0].Name != "token" {
                t.Errorf("Incorrect cookie name: %v", cookieCheck[0].Name)
            }
        } else {
            t.Errorf("Cookies not right %v", cookieCheck)
        }
    })
}

func BenchmarkNoCredsTime(b *testing.B) {
    
    b.Run("Testing Incorrect Creds", func(b *testing.B) {
        request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "sdfsdf", Password: "hgfjfgh"})))
        w := httptest.NewRecorder()
        handler := http.HandlerFunc(ValidateUser)
  
        b.ReportAllocs()
        b.ResetTimer()
  
        for i := 0; i < b.N; i++ {
            handler.ServeHTTP(w, request)
        }
    })
    b.Run("Testing with Correct Creds", func(b *testing.B) {
        request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "TestUser1", Password: "SomeTestpwd"})))
        w := httptest.NewRecorder()
        handler := http.HandlerFunc(ValidateUser)
  
        b.ReportAllocs()
        b.ResetTimer()
  
        for i := 0; i < b.N; i++ {
            handler.ServeHTTP(w, request)
        }
    })
}

func createJson(creds Credentuals) string {
    u, err := json.Marshal(creds)
    if err != nil {
        panic(err)
    }

    return string(u)
}