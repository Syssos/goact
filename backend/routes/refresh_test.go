package routes

import (
	"time"
    "strings"
    "testing"
    "net/http"
    "net/http/httptest"

	"github.com/dgrijalva/jwt-go"
)

func TestTokenRefresh(t *testing.T) {
    t.Run("Testing with no token", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/refresh", nil)
        response := httptest.NewRecorder()
        
        Refresh(response, request)
        
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
	t.Run("Testing token with long expiry timer", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/refresh", nil)
        response := httptest.NewRecorder()
		
		cookie := http.Cookie{Name: "token", Value: GenerateTokenString()}
		request.AddCookie(&cookie)
        
		Refresh(response, request)
        
        got := response.Code
        want := 400

        cookieCheck := response.Result().Cookies()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
        if len(cookieCheck) > 0 {
			t.Errorf("Cookie present, potentual security issue, or forcing clients to re-validate")
		}
    })
	t.Run("Testing token with short expiry timer", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/refresh", nil)
        response := httptest.NewRecorder()
		
		cookie := http.Cookie{Name: "token", Value: GenerateAlmostExpired()}
		request.AddCookie(&cookie)
        
		Refresh(response, request)
        
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
			// Getting cookie data from "token" cookie
			// Then validating if the cookie, is a valid jwt token
			tokenStr := cookieCheck[0].Value
			claims := &Claims{}
			tkn, err := jwt.ParseWithClaims(tokenStr, claims,
				func(t *jwt.Token) (interface{}, error){
					return jwtKey, nil
				})
			
			// check if theres an error validating jwt token.
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					t.Errorf("Incorrect cookie signature")
				}
				t.Errorf("Error not nil validating jwt returned")
			}

			if !tkn.Valid {
				t.Errorf("Returned token is not valid")
			}

			if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 60 * time.Second {
				t.Errorf("Token was not refreshed")
			}
		} else {
			t.Errorf("No Cookie Present, cookie was not refreshed")
		}
    })
}

func GenerateTokenString() string {
	request, _ := http.NewRequest(http.MethodPost, "/validate", strings.NewReader(createJson(Credentuals{Username: "TestUser1", Password: "SomeTestpwd"})))
	response := httptest.NewRecorder()
	ValidateUser(response, request)
	
	return response.Result().Cookies()[0].Value
}

func GenerateAlmostExpired() string {
	expirationTime := time.Now().Add(time.Second * 45)
	claims := &Claims{
		Username: "TestUser1",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Create a signed token string securing the claims authenticity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}

	return tokenString
}

func GenerateExpired() string {
	count := 10
	expirationTime := time.Now().Add(time.Duration(-count) * time.Minute)
	claims := &Claims{
		Username: "TestUser1",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Create a signed token string securing the claims authenticity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}

	return tokenString
}