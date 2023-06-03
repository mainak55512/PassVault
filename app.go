package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StorePass(service_name, user_id, password string) {
	PassStore(service_name, user_id, password)
}

func (a *App) FetchCredDetails() []string {
	return Fetch_cred_details()
}


func (a *App) FetchPass(service_name string) []string {
	return Fetch_creds(service_name)
}

func (a *App) DelPass(service_name string) {
	Del_cred(service_name)
}

func (a *App) GenPass(length int, upper_case, numbers, spl_char bool) string {
	return Generate_password(length,upper_case,numbers,spl_char)
}

func (a *App) Login_cred( s string) bool {
	return ValidLogin(s)
}

func (a *App) SetPin( s,q1,q2 string){
	Setpin(s,q1,q2)
}

func (a *App) ChangePin( s string){
	Changepin(s)
}

func (a *App) IsNew() bool {
	return CheckConfig()
}

func (a *App)CheckSecurityQuestion(s1,s2 string) bool {
	return checkScQues(s1,s2)
}

func (a *App)GenerateRSAKeys(){
	Generate_rsa_Key_pair()
}