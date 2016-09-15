package main

import (
	//"fmt"
	"flag"
	//"time"
	//"strings"

	"./core"

	//"github.com/kataras/iris"
)

func main() {
	flag.IntVar(&core.Instance, "instance",1,"Instance ID Number for this application execution cycle.")
	flag.Parse()
	core.Test()
	//api.run("./config/config.json")
}
