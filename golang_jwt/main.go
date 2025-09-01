package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

//without set JWT in cookie
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		_, err := VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}


//JWT Set in Cookies
// func AuthMiddleware(next http.Handler) http.Handler{
//  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//       cookie, err := r.Cookie("jwt")
// 		if err != nil {
// 			http.Error(w, "Unauthorized: missing cookie", http.StatusUnauthorized)
// 			return
// 		}

// 		_ , err = VerifyJWT(cookie.Value)
// 		if err != nil {
// 			http.Error(w, "Unauthorized: invalid or expired token", http.StatusUnauthorized)
// 			return
// 		}

// 	next.ServeHTTP(w , r)
//  })
// }

// If using Cookies

// Logout is easy → just clear the cookie from the client by setting it with an expired date.

// func HandleLogout(w http.ResponseWriter, r *http.Request) {
//     // overwrite the jwt cookie with expired one
//     http.SetCookie(w, &http.Cookie{
//         Name:     "jwt",
//         Value:    "",
//         Path:     "/",
//         HttpOnly: true,
//         Secure:   false,
//         Expires:  time.Unix(0, 0), // expired in the past
//         MaxAge:   -1,
//     })

//     w.WriteHeader(http.StatusOK)
//     w.Write([]byte("Logged out successfully"))
// }

func HandleLogin(w http.ResponseWriter , r *http.Request){
   user := &Users{}

  decoder := json.NewDecoder(r.Body)
  decoder.Decode(&user)

  if user.Password!= userdemo.Password && user.Name !=userdemo.Name{
	http.Error(w , "try to login With wrong credentials",http.StatusUnauthorized)
	return
  }

  token, err := user.GenerateToken()
  if err!=nil{
	http.Error(w , "Failed To generate Token",http.StatusInternalServerError)
	return
  }


//   http.SetCookie(w ,&http.Cookie{
//     Name: "jwt",
// 	Value: token,
// 	Path: "/",
// 	HttpOnly: true,
//     Secure: false,
// 	SameSite: http.SameSiteStrictMode,
//   })
 
 json.NewEncoder(w).Encode(map[string]string{
	"token": token,
	"massage":"\nLogin successful! JWT set in cookie",
})

}


func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(userdemo)


	w.WriteHeader(200)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	w.Write([]byte("Hello, your user ID is " + userID))
}



func main() {
	router := http.NewServeMux()


   router.Handle("POST /login",http.HandlerFunc(HandleLogin))

   router.Handle("GET /users", http.HandlerFunc(GetAllUsers))
   router.Handle("GET /users/{id}", AuthMiddleware(http.HandlerFunc(ProtectedHandler)))

	fmt.Println("Server Starting on port 8000")
	err := http.ListenAndServe(":8000",router)
	if err!= nil{
		fmt.Println("Failed to starting Server")
	}

}

