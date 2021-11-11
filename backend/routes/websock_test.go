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

        res  := ValidateCookieFMT(request)

        if res != "" {
            t.Error("Could not detect bad token name")
        }
    })
    t.Run("Testing Correct Cookie name", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/ws", nil)
        tkStr := GenerateExpired()
        
        cookie := http.Cookie{Name: "token", Value: tkStr}
        request.AddCookie(&cookie)

        res  := ValidateCookieFMT(request)

        if res != tkStr {
            t.Error("Could not detect bad token name")
        }
    })
}

func TestValidateCookieJWT(t *testing.T) {
    t.Run("Testing Cookie validation with bad token", func(t *testing.T) {
        res := ValidateCookieJWT("TheHamburgler")

        if res != "" {
            t.Error("Validation string not blank with incorrect JWT format")
        }
    })

    t.Run("Testing Cookie validation with expired token", func(t *testing.T) {
        res  := ValidateCookieJWT(GenerateExpired())

        if res != "" {
            t.Error("Could not detect bad token string")
        }
    })

    t.Run("Testing Cookie validation with valid token", func(t *testing.T) {
        res  := ValidateCookieJWT(GenerateTokenString())

        if res != "TestUser1" {
            t.Error("Could not good token string")
        }
    })
}