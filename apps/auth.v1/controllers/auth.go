package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/context"

	"../../../render"
	"../../../system"
	"bitbucket.org/roporter-mydev/auth"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")

	val := context.Get(r, "application")
	if val != nil {
		application := val.(*system.Application)

		if succ, _ := auth.Authenticate(token, *application.DBOptions); succ == true {
			fmt.Println("DONE")
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("Email")
	pass := r.Header.Get("Password")
	accept := ""
	result := ""
	token := ""
	err := errors.New("")

	if len(r.Header["Accept"]) == 1 {
		accept = r.Header["Accept"][0]
	} else {
		accept = "application/json"
	}

	val := context.Get(r, "application")
	if val != nil {
		application := val.(*system.Application)
		DBType := application.Configuration.GetString("database.type")

		if DBType == "bolt" {
			token, err = auth.LoginBolt(email, pass, application.Bolt.UserDBStore, *application.DBOptions)
		} else if DBType == "ledis" {
			token, err = auth.LoginLedis(email, pass, application.Ledis.UserDBStore, *application.DBOptions)
		} else if DBType == "mongo" {
			token, err = auth.LoginMongo(email, pass, application.Mongo.UserDBStore, *application.DBOptions)
		}

		if err == nil {
			timeout := application.Configuration.GetString("certificate.Timeout")
			result = `{"Login":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Duration":"` + timeout + `","Token":"` + token + `"}}`
		} else {
			result = `{ "CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Error":"` + err.Error() + `"}`
		}

		if accept == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if accept == "application/xml" {
			render.RenderXML(w, http.StatusOK, `<xml><token>`+token+`</token></xml>`)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func Impersonate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	email := r.Header.Get("Email")
	result := ""
	newToken := ""
	userId := ""
	accept := ""

	if len(r.Header["Accept"]) == 1 {
		accept = r.Header["Accept"][0]
	} else {
		accept = "application/json"
	}

	fmt.Println(accept)
	fmt.Println(userId)

	val := context.Get(r, "application")
	if val != nil {
		application := val.(*system.Application)
		DBType := application.Configuration.GetString("database.type")
		err := errors.New("")

		if DBType == "bolt" {
			userId, err = auth.AuthenticateUserIdBolt(token, application.Bolt.UserDBStore, *application.DBOptions)
		} else if DBType == "ledis" {
			userId, err = auth.AuthenticateUserIdLedis(token, application.Ledis.UserDBStore, *application.DBOptions)
		} else if DBType == "mongo" {
			userId, err = auth.AuthenticateUserIdMongo(token, application.Mongo.UserDBStore, *application.DBOptions)
		}

		cont := false

		if err == nil {
			valid, err2 := auth.ValidateScope(token, "IMPERSO", *application.DBOptions)
			if err2 == nil {
				cont = valid
			} else {
				//result = system.GenerateResponse()
				result = ""
			}

			if cont {
			}

			if err == nil {
				newToken = ""
				if DBType == "bolt" {
					newToken, err = auth.ImpersonateBolt(email, application.Bolt.UserDBStore, *application.DBOptions)
				} else if DBType == "ledis" {
					newToken, err = auth.ImpersonateLedis(token, application.Ledis.UserDBStore, *application.DBOptions)
				} else if DBType == "mongo" {
					newToken, err = auth.ImpersonateMongo(token, application.Mongo.UserDBStore, *application.DBOptions)
				}

				if err == nil {
					timeout := application.Configuration.GetString("certificate.Timeout")
					result = `{"Impersonate":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","User":"` + email + `","Duration":"` + timeout + `","Token":"` + newToken + `"}}`
					render.RenderJSON(w, http.StatusOK, result)
				} else {
					result = `{"Impersonate":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Error":"` + err.Error() + `"}}`
					render.RenderJSON(w, http.StatusUnauthorized, result)
				}
			} else {
				result = `{"Impersonate":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Error":"` + err.Error() + `"}}`
				render.RenderJSON(w, http.StatusUnauthorized, result)
			}
		}
	}
}

func Signup(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Signup")
}

func RequestToJsonObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(jsonDoc)
	if err != nil {
		return err
	}
	return nil
}
