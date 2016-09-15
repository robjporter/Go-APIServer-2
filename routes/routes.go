package routes

import (
	authRoutes "../apps/auth.v1/routes"
	bingRoutes "../apps/bing.v1/routes"
	coreRoutes "../apps/core.v1/routes"
	funRoutes "../apps/fun.v1/routes"
	sparkRoutes "../apps/spark.v1/routes"
	statsRoutes "../apps/stats.v1/routes"
	systemRoutes "../apps/system.v1/routes"
	"github.com/go-zoo/bone"
)

func Include(r *bone.Mux) *bone.Mux {
	// Core app
	core := coreRoutes.Routes()
	coreRoute := coreRoutes.GetRouteURL()
	coreRoutes.Prefetch()

	// Fun App
	fun := funRoutes.Routes()
	funRoute := funRoutes.GetRouteURL()
	funRoutes.Prefetch()

	// System app
	sys := systemRoutes.Routes()
	sysRoute := systemRoutes.GetRouteURL()
	systemRoutes.Prefetch()

	// Bing app
	bing := bingRoutes.Routes()
	bingRoute := bingRoutes.GetRouteURL()
	bingRoutes.Prefetch()

	// Spark app
	spark := sparkRoutes.Routes()
	sparkRoute := sparkRoutes.GetRouteURL()
	sparkRoutes.Prefetch()

	// Spark app
	auth := authRoutes.Routes()
	authRoute := authRoutes.GetRouteURL()
	authRoutes.Prefetch()

	// Stats app
	stats := statsRoutes.Routes()
	statsRoute := statsRoutes.GetRouteURL()
	statsRoutes.Prefetch()

	r.SubRoute(coreRoute, core)
	r.SubRoute(funRoute, fun)
	r.SubRoute(sysRoute, sys)
	r.SubRoute(bingRoute, bing)
	r.SubRoute(sparkRoute, spark)
	r.SubRoute(authRoute, auth)
	r.SubRoute(statsRoute, stats)
	return r
}
