package utils

// import(
//   "github.com/golang-jwt/jwt/v5"
//   "time"
// )

// func CreateToken(username ,role string , id int ){

//     tokenString, err := claims.SignedString(secretKey)
//     if err != nil {
//         return "", err
//     }
//   claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": username,                    
// 		"iss": "taskmanager",                 
// 		"aud": role,           
// 		"exp": time.Now().Add(time.Hour).Unix(), 
// 		"iat": time.Now().Unix(),               
// 	})

//   // Print information about the created token
// 	fmt.Printf("Token claims added: %+v\n", claims)
// 	return tokenString, nil
// }