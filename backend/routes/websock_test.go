package routes

import (
    "net/http"
    "testing"
)

func TestValidateFMTCheck(t *testing.T) {
    t.Run("Testing Incorrect Cookie name", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        tkStr := "IncorrectValueForSure"

        cookie := http.Cookie{Name: "noket", Value: tkStr}
        request.AddCookie(&cookie)

        res, err  := ValidateCookieFMT(request)
        if err == nil {
            t.Error("error not detected for bad token name")
        }

        if res != "" {
            t.Error("Accepted incorrectly named cookie")
        }
    })
    t.Run("Testing Correct Cookie name", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        tkStr := GenerateExpired()
        
        cookie := http.Cookie{Name: "token", Value: tkStr}
        request.AddCookie(&cookie)

        res, err  := ValidateCookieFMT(request)
        if err != nil {
            t.Error("Error Validating Cookie")
        }
        if res != tkStr {
            t.Error("Couldn't validate correct string")
        }
    })
}

func TestValidateCookieJWT(t *testing.T) {
    t.Run("Testing Cookie validation with bad token", func(t *testing.T) {
        res, err := ValidateCookieJWT("TheHamburgler")
        if err == nil {
            t.Error("Error not detected while validating JWT String")
        }
        if res != "" {
            t.Error("Validation string not blank with incorrect JWT format")
        }
    })

    t.Run("Testing Cookie validation with expired token", func(t *testing.T) {
        res, err  := ValidateCookieJWT(GenerateExpired())
        if err == nil {
            t.Error("Error not detected while validating JWT String")
        }
        if res != "" {
            t.Error("Could not detect bad token string")
        }
    })

    t.Run("Testing Cookie validation with valid token", func(t *testing.T) {
        res, err  := ValidateCookieJWT(GenerateTokenString())
        if err != nil {
            t.Error("Error Validating JWT String")
        }

        if res != "TestUser1" {
            t.Error("Could not validate good token string, or obtain username")
        }
    })
}