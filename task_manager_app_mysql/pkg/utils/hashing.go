package utils

import(

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string , error){
  
	hashValue , err := bcrypt.GenerateFromPassword([]byte (password),13)

	return string(hashValue), err	
}

func CheckPassword( hashPassword string ,password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))

	return err==nil
}