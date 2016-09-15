package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"./routes"
	"./system"
	"bitbucket.org/roporter-mydev/logging"
	"bitbucket.org/roporter-mydev/secure"
	"bitbucket.org/roporter-mydev/xrequestid2"
	"bitbucket.org/roporter-mydev/xrevision"

	"github.com/blang/semver"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"github.com/gorilla/context"
	"github.com/tylerb/graceful"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var plugins []system.Module
	var application = &system.Application{}
	defer application.Close()

	application.Versions.BuildRevision = system.GetGitBuild(true)
	application.Versions.HeaderRevision = system.GetGitBuild(false)
	application.Versions.AppVersion, _ = semver.New(system.GetAppBuild(application.Versions.BuildRevision))
    
	configLocations := []string{"conf"}
	secureMiddleware := secure.New(secure.Options{FrameDeny: true, IsDevelopment: true})

	application.Startup()
	application.LoadConfig(configLocations)
	application.Init()
	application.DisplayIntro()
	application.SetupSecurity()
	application.ConnectToDatabase()
	application.LoadTemplates()

	r := bone.New()
	r = routes.Include(r)

	plugins = application.GetPlugins()

	for i := 0; i < len(plugins); i++ {
		mod := plugins[i]
		r.Handle("/assets/"+mod.Name+"/", http.StripPrefix("/assets/"+mod.Name, http.FileServer(http.Dir(mod.Path+"/public"))))
	}
	// Add this last after all other routes - StaticRoute function of Bone, checks a slice for shortest match - this is the shortest match, so needs to be final match
	r.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir(application.Configuration.GetString("PublicPath")))))

	r.HandleFunc("/", yourHandler)
	r.HandleFunc("/test", testHandler)
	r.GetFunc("/reg/#var^[a-z]$/#var2^[0-9]$", showVar)

	//r.DisplayAllStaticRoutes()

	n := negroni.New()
	n.UseFunc(logging.ServeHTTP)
	n.UseFunc(application.Settings.Stats.ServeHTTP)
	n.UseFunc(xrequestid2.ServeHTTP)
	n.UseFunc(application.ServeHTTP)
	n.Use(xrevision.New(application.Versions.HeaderRevision))
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.Use(negroni.NewRecovery())
    
	n.UseHandler(r)

	graceful.Run(":"+application.Configuration.GetString("host.port"), time.Duration(application.Configuration.GetInt("host.timeout"))*time.Second, n)
}

func showVar(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(bone.GetAllValues(req)["var"]))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "application")
	if val != nil {
		application := val.(*system.Application)

		fmt.Println(application.Configuration.GetString("PublicPath"))
	}

	w.Write([]byte("Done!\n"))
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func debugPrintVersion() {
	//fmt.Printf("Major: %d\n", application.Versions.AppVersion.Major)
	//fmt.Printf("Minor: %d\n", application.Versions.AppVersion.Minor)
	//fmt.Printf("Patch: %d\n", application.Versions.AppVersion.Patch)
	//fmt.Printf("Pre: %s\n", application.Versions.AppVersion.Pre)
	//fmt.Printf("Build: %s\n", application.Versions.AppVersion.Build)
}
