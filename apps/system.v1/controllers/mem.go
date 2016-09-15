package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../../../render"
	//"../../../system"
	"github.com/shirou/gopsutil/mem"
)

func MemInfoSummary(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := "TMP" //`{ "HOST" : { "NET" : "` + tmp3 + `"}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapSummary(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := tmp2 //`{ "HOST" : { "NET" : "` + tmp3 + `"}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapTotal(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := `{ "HOST" : { "MEMORY" : { "SWAP" : { "TOTAL" : "` + strconv.Itoa(int(tmp2.Total)) + `"}}}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapUsed(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := `{ "HOST" : { "MEMORY" : { "SWAP" : { "USED" : "` + strconv.Itoa(int(tmp2.Used)) + `"}}}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapFree(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := `{ "HOST" : { "MEMORY" : { "SWAP" : { "FREE" : "` + strconv.Itoa(int(tmp2.Free)) + `"}}}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapUsedPercent(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := `{ "HOST" : { "MEMORY" : { "SWAP" : { "USEDPERCENT" : "` + strconv.Itoa(int(tmp2.UsedPercent)) + `"}}}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapSin(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := `{ "HOST" : { "MEMORY" : { "SWAP" : { "SIN" : "` + strconv.Itoa(int(tmp2.Sin)) + `"}}}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemSwapSout(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.SwapMemory()

	result := `{ "HOST" : { "MEMORY" : { "SWAP" : { "SOUT" : "` + strconv.Itoa(int(tmp2.Sout)) + `"}}}}`
	fmt.Println(tmp2)
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualSummary(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := tmp2

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualTotal(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "TOTAL" : "` + strconv.Itoa(int(tmp2.Total)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualAvailable(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "AVAILABLE" : "` + strconv.Itoa(int(tmp2.Available)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualUsed(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "USED" : "` + strconv.Itoa(int(tmp2.Used)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
	//used_percent":62.35542297363281,"free":200646656,"active":6214094848,"inactive":5853253632,"buffers":0,"cached":6266642432,"wired":3115225088,"shared":0}
}

func MemInfoMemVirtualUsedPercent(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "USEDPERCENT" : "` + strconv.Itoa(int(tmp2.UsedPercent)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualFree(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "FREE" : "` + strconv.Itoa(int(tmp2.Free)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualActive(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "ACTIVE" : "` + strconv.Itoa(int(tmp2.Active)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualInactive(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "INACTIVE" : "` + strconv.Itoa(int(tmp2.Inactive)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualBuffers(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "BUFFERS" : "` + strconv.Itoa(int(tmp2.Buffers)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualCached(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "CACHED" : "` + strconv.Itoa(int(tmp2.Cached)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func MemInfoMemVirtualWired(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, _ := mem.VirtualMemory()

	result := `{ "HOST" : { "MEMORY" : { "VIRTUAL" : { "WIRED" : "` + strconv.Itoa(int(tmp2.Wired)) + `"}}}}`

	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, result)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, result)
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}
