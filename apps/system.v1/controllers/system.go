package controllers

import (
	"fmt"
	"net/http"
	"time"

	"../../../render"
	//"../../../system"

	"strconv"

	"github.com/shirou/gopsutil/cpu"
)

func CPUHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("CPUHome!\n"))
}

func SysHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("SysHome!\n"))
}

func SysCoreRaw(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusOK, "AllData: Data")
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusOK, "AllData: Data")
		} else {
			fmt.Println("NOT SURE WHERE WE ARE!")
		}
	}
}

func SysCoreCount(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUCounts(true)
	if err != nil {
		tmp2 = 0
	}
	result := `{ "CPU" : { "CORES" : { "COUNT" : "` + strconv.Itoa(tmp2) + `"}}}`
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

func SysCPUPercent1Sec(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUPercent(time.Second, true)
	if err != nil {
		panic("No return while checking CPU")
		return
	}

	result := `{ "CPU" : { "USAGE" : {`
	for i := 0; i < len(tmp2); i++ {
		result += `"CORE" : "` + strconv.Itoa(int(tmp2[i])) + `"`
		if i+1 != len(tmp2) {
			result += `,`
		}
	}
	result += "}}}"
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

func SysCPUPercent1SecTot(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUPercent(time.Second, false)
	if err != nil {
		panic("No return while checking CPU")
		return
	}

	result := `{ "CPU" : { "USAGE" : {`
	for i := 0; i < len(tmp2); i++ {
		result += `"CORE" : "` + strconv.Itoa(int(tmp2[i])) + `"`
		if i+1 != len(tmp2) {
			result += `,`
		}
	}

	result += "}}}"
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

func SysCPUInfo(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	//result := `{ "CPU" : { "INFO" : ` + tmp3 + `}}`
	result := tmp2[0]
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

func SysCPUInfoVendor(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "VENDOR_ID" : "` + tmp2[0].VendorID + `"}}}`
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

func SysCPUInfoFamily(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "FAMILY" : "` + tmp2[0].Family + `"}}}`
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

func SysCPUInfoModel(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "MODEL" : "` + tmp2[0].Model + `"}}}`
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

func SysCPUInfoCores(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "CORES" : "` + strconv.FormatInt(int64(tmp2[0].Cores), 10) + `"}}}`
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

func SysCPUInfoName(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "NAME" : "` + tmp2[0].ModelName + `"}}}`
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

func SysCPUInfoSpeed(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "SPEED_Mhz" : "` + strconv.FormatInt(int64(tmp2[0].Mhz), 10) + `"}}}`
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

func SysCPUInfoCache(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}
	result := `{ "CPU" : { "INFO" : { "CACHE" : "` + strconv.FormatInt(int64(tmp2[0].CacheSize), 10) + `"}}}`
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

func SysCPUInfoFlags(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	tmp2, err := cpu.CPUInfo()
	if err != nil {
		panic("No CPU Info received.")
		return
	}

	result := `{ "CPU" : { "INFO" : { "FLAGS" : {`
	flags := tmp2[0].Flags
	for i := 0; i < len(flags); i++ {
		result += `"FLAG` + strconv.Itoa(i+1) + `" : "` + flags[i] + `"`
		if i+1 != len(flags) {
			result += `,`
		}
	}
	result += "}}}}"

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

func SysRamCount(w http.ResponseWriter, req *http.Request) {
}
