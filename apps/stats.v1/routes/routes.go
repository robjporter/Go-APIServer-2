package routes

import (
	"../controllers"
	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/stats"
}

func Routes() *bone.Mux {
	tmp := bone.New()
	tmp.GetFunc("/", controllers.StatsHome)
	tmp.GetFunc("/dashboard", controllers.StatsDashboard)
	tmp.GetFunc("/raw", controllers.StatsRaw)
	tmp.GetFunc("/averageresponsetime", controllers.StatsAverageResponseTime)
	tmp.GetFunc("/methodtypes", controllers.StatsRequestTypeCounts)
	tmp.GetFunc("/uptime", controllers.StatsUpTime)
	tmp.GetFunc("/time", controllers.StatsTime)
	tmp.GetFunc("/count", controllers.StatsCount)
	tmp.GetFunc("/calls", controllers.StatsCalls)
	tmp.GetFunc("/requests", controllers.StatsRequests)
	tmp.GetFunc("/codes", controllers.StatsCodes)
	return tmp
}
