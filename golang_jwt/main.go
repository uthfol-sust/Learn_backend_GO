package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Users struct{
	Id int     `json:"id"`
	Name string `json:"name"`
	Password string `json:"-"`
}

var userdemo = Users{Id:1 ,Name:"Uthpol",Password: "123456" }

type MyCustomClaims struct{
	Id int     
	Name string
	Role string
	jwt.RegisteredClaims
}

func (u *Users) GenerateToken() (string, error) {
	claims := MyCustomClaims{
		Id:   u.Id,
		Name: u.Name,
        Role: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MyTaskManager",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("my_secret_key")) 
}


func VerifyJWT(tokenString string) (*MyCustomClaims , error){
   token , err := jwt.ParseWithClaims(tokenString ,&MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
	return []byte("my_secret_key"), nil
   })

   if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}
	return claims, nil
}

func HandleLogin(w http.ResponseWriter , r *http.Request){
  user := &Users{}
  er:=json.NewDecoder(r.Body).Decode(user)
  
  fmt.Println(user)
  if user.Password!=userdemo.Password || user.Name!=userdemo.Name || er!=nil{
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	return
   }


  token , err := user.GenerateToken()
  if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

http.SetCookie(w,&http.Cookie{
	Name: "jwt",
	Value: token,
	Path: "/",
	HttpOnly: true,
	Secure: false,
	SameSite: http.SameSiteStrictMode,
})
 
 json.NewEncoder(w).Encode(map[string]string{
	"token": token,
	"massage":"Login successful! JWT set in cookie",
})

}



func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login",HandleLogin).Methods("POST")
	


	log.Fatal(http.ListenAndServe(":8000",router))
}

