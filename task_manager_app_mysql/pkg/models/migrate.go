package models

func MigrateAll(){
	UserAutoMigrate()
    TaskAutoMigrate()
    EmailVerificationAutoMigrate()
    UseResetPasswordAutomigrate()
}