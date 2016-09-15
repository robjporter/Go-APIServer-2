package routes

import (
	"../controllers"

	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/core"
}

func Routes() *bone.Mux {
	tmp := bone.New()
	tmp.GetFunc("/", controllers.CoreHome)
	tmp.GetFunc("/home", controllers.CoreHome)
	tmp.GetFunc("/about", controllers.CoreAbout)
	return tmp
}
