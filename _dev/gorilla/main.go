package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"bitbucket.org/roporter-mydev/color"
	"bitbucket.org/roporter-mydev/logging"
	"bitbucket.org/roporter-mydev/stats"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	//	var conf *viper.Viper

	fmt.Println(color.Red("TEST"))

	// Tell app how many resources are available
	runtime.GOMAXPROCS(runtime.NumCPU())

	//	conf = system.LoadConfig("config", []string{"./conf"})
	//	fmt.Println(conf)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in

	fmt.Println(r)

	n := negroni.New()
	stat := stats.New()

	n.UseFunc(logging.ServeHTTP)
	n.UseFunc(stat.ServeHTTP)
	n.Use(negroni.NewRecovery())

	n.UseHandler(r)
	//n.Run(":3000")
	graceful.Run(":3000", 10*time.Second, n)
}
