package routes

import (
    "strings"
    "testing"
    "net/http"
    "encoding/json"
    "net/http/httptest"
)

func TestCredentuals(t *testing.T) {
    t.Run("Testing with no credentuals", func(t *testing.T) {
        creds := createJson(Credentuals{Username: "", Password: ""})

        response := httptest.NewRecorder()
        request, err := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(creds))
        if err != nil {
            t.Error("Error creating test request")
        }

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
        creds := createJson(Credentuals{Username: "ThistWrong", Password: "invalid"})
        
        response := httptest.NewRecorder()
        request, err := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(creds))
        if err != nil {
            t.Error("Error creating test request")
        }

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
        creds := createJson(Credentuals{Username: "&^(*%*&%@$*)*&@#$(", Password: "!@#$%^&*()_+{}|:\"<>?"})
        
        response := httptest.NewRecorder()
        request, err := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(creds))
        if err != nil {
            t.Error("Error creating test request")
        }

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
        creds := createJson(Credentuals{Username: "TestUser1", Password: "SomeTestpwd"})
        
        response := httptest.NewRecorder()
        request, err := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(creds))
        if err != nil {
            t.Error("Error creating test request")
        }

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

func createJson(creds Credentuals) string {
    u, err := json.Marshal(creds)
    if err != nil {
        panic(err)
    }

    return string(u)
}