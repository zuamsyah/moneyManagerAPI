package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"moneyManagerAPI/models"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var JWT_SIGNATURE_KEY = "LKJS@3KJDKS"

func AuthOtpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.User
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &u)

	if u.PhoneNumber == "08923712733" && u.Otp == 123456 {
		tokenString, err := CreateToken(u.PhoneNumber)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No phone number found")
		}

		response := map[string]interface{}{"token": tokenString}

		result, _ := json.Marshal(&response)

		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func CreateToken(phone_number string) (string, error) {
	mapClaim := jwt.MapClaims{}
	mapClaim["phone_number"] = phone_number
	mapClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaim)
	token, _ := sign.SignedString([]byte(JWT_SIGNATURE_KEY))
	return token, nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/api/auth_otp" {
			w.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Missing authorization header")
				return
			}

			_, err := verifyToken(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid token")
				return
			}

			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		}
		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SIGNATURE_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token.Claims, nil
}
