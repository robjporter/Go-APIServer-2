package routes

import (
	"../controllers"
	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/auth"
}

func Routes() *bone.Mux {
	tmp := bone.New()
	tmp.PostFunc("/login", controllers.Login)
	tmp.PostFunc("/impersonate", controllers.Impersonate)
	tmp.GetFunc("/validate", controllers.Validate)
	tmp.GetFunc("/signup", controllers.Signup)
	return tmp
}
