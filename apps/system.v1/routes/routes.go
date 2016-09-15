package routes

import (
	"../controllers"

	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/system"
}

func Routes() *bone.Mux {
	tmp := bone.New()
	// TESTING
	tmp.GetFunc("/tmp", controllers.TMP)
	tmp.GetFunc("/tmp2", controllers.TMP2)

	// SYSTEM
	tmp.GetFunc("/", controllers.SysHome)
	tmp.GetFunc("/raw", controllers.SysCoreRaw)
	tmp.GetFunc("/cores", controllers.SysCoreCount)
	tmp.GetFunc("/ram", controllers.SysRamCount)

	// CPU
	tmp.GetFunc("/cpu", controllers.CPUHome)
	tmp.GetFunc("/cpu/percent/sec", controllers.SysCPUPercent1Sec)
	tmp.GetFunc("/cpu/percent/sec/tot", controllers.SysCPUPercent1SecTot)
	tmp.GetFunc("/cpu/info/raw", controllers.SysCPUInfo)
	tmp.GetFunc("/cpu/info/vendor", controllers.SysCPUInfoVendor)
	tmp.GetFunc("/cpu/info/family", controllers.SysCPUInfoFamily)
	tmp.GetFunc("/cpu/info/model", controllers.SysCPUInfoModel)
	tmp.GetFunc("/cpu/info/cores", controllers.SysCPUInfoCores)
	tmp.GetFunc("/cpu/info/name", controllers.SysCPUInfoName)
	tmp.GetFunc("/cpu/info/speed", controllers.SysCPUInfoSpeed)
	tmp.GetFunc("/cpu/info/cache", controllers.SysCPUInfoCache)
	tmp.GetFunc("/cpu/info/flags", controllers.SysCPUInfoFlags)

	// HOST
	tmp.GetFunc("/host", controllers.HostHome)
	tmp.GetFunc("/host/boottime", controllers.HostBootTime)
	tmp.GetFunc("/host/os", controllers.HostOS)
	tmp.GetFunc("/host/version", controllers.HostVersion)
	tmp.GetFunc("/host/users", controllers.HostActiveUsers)
	tmp.GetFunc("/host/virtualisation/system", controllers.HostVirtualisationSystem)
	tmp.GetFunc("/host/virtualisation/role", controllers.HostVirtualisationRole)

	// NET
	tmp.GetFunc("/net", controllers.NetInfoAdapterSummary)
	tmp.GetFunc("/net/bytes/sent", controllers.NetInfoAdapterSummaryBytesSent)
	tmp.GetFunc("/net/bytes/recv", controllers.NetInfoAdapterSummaryBytesRecv)
	tmp.GetFunc("/net/packets/sent", controllers.NetInfoAdapterSummaryPacketsSent)
	tmp.GetFunc("/net/packets/recv", controllers.NetInfoAdapterSummaryPacketsRecv)
	tmp.GetFunc("/net/errors/in", controllers.NetInfoAdapterSummaryErrorIn)
	tmp.GetFunc("/net/errors/out", controllers.NetInfoAdapterSummaryErrorOut)
	tmp.GetFunc("/net/adapters/summary", controllers.NetInfoAdaptersSummary)
	tmp.GetFunc("/net/adapters/:interface/summary", controllers.NetInfoAdaptersSummary)
	tmp.GetFunc("/net/adapters/summary/detail", controllers.NetInfoAdaptersSummaryDetail)
	tmp.GetFunc("/net/adapters/:interface/summary/detail", controllers.NetInfoAdaptersSummaryDetail)

	// MEM
	tmp.GetFunc("/mem", controllers.MemInfoSummary)
	tmp.GetFunc("/mem/swap/summary", controllers.MemInfoMemSwapSummary)
	tmp.GetFunc("/mem/swap/total", controllers.MemInfoMemSwapTotal)
	tmp.GetFunc("/mem/swap/used", controllers.MemInfoMemSwapUsed)
	tmp.GetFunc("/mem/swap/free", controllers.MemInfoMemSwapFree)
	tmp.GetFunc("/mem/swap/usedpercent", controllers.MemInfoMemSwapUsedPercent)
	tmp.GetFunc("/mem/swap/sin", controllers.MemInfoMemSwapSin)
	tmp.GetFunc("/mem/swap/sout", controllers.MemInfoMemSwapSout)
	tmp.GetFunc("/mem/virtual/summary", controllers.MemInfoMemVirtualSummary)
	tmp.GetFunc("/mem/virtual/total", controllers.MemInfoMemVirtualTotal)
	tmp.GetFunc("/mem/virtual/available", controllers.MemInfoMemVirtualAvailable)
	tmp.GetFunc("/mem/virtual/used", controllers.MemInfoMemVirtualUsed)
	tmp.GetFunc("/mem/virtual/usedpercent", controllers.MemInfoMemVirtualUsedPercent)
	tmp.GetFunc("/mem/virtual/free", controllers.MemInfoMemVirtualFree)
	tmp.GetFunc("/mem/virtual/active", controllers.MemInfoMemVirtualActive)
	tmp.GetFunc("/mem/virtual/inactive", controllers.MemInfoMemVirtualInactive)
	tmp.GetFunc("/mem/virtual/buffers", controllers.MemInfoMemVirtualBuffers)
	tmp.GetFunc("/mem/virtual/cached", controllers.MemInfoMemVirtualCached)
	tmp.GetFunc("/mem/virtual/wired", controllers.MemInfoMemVirtualWired)

	// LOAD
	tmp.GetFunc("/load", controllers.LoadInfoSummary)

	return tmp
}
