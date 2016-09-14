package core

import (
    "github.com/kataras/iris"
    "github.com/dgrijalva/jwt-go"
    conf "github.com/roporter/go-libs/go-config"
)

type MyCustomClaims struct {
	Name string 		`json:"name"`
	Username string 	`json:"username"`
	Email string 		`json:"email"`
	Role string 		`json:"role"`
	Capabilities string `json:"capabilities"`
	jwt.StandardClaims
}

type myIris struct {
	core 			iris.Framework
	config 			conf.Config
	//stats 			*stats.Stats
	//ipfilter		*ipfilter.IPFilter
	errorLogger 	iris.HandlerFunc
	//scheduler		*gocron.Scheduler
	//consul 			*consulapi.Client
	//jwt				*jwtmiddleware.Middleware
	system			SystemData
	//logLogger		*logrus.Logger
	//log				*logrus.Entry
	paths       	[]string
	mainRoutes  	interface{}
	pluginFolders 	[]string
	plugins     	interface{}
	version     	string
}

type SystemData struct {
	ip				string
	hostname		string
	city			string
	region 			string
	country 		string
	location 		string
	organisation	string
	getRoutes		int
	postRoutes		int
	headRoutes		int
	putRoutes		int
	deleteRoutes    int
}

type myPlugin struct{}
