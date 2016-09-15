package routes

import (
	"../controllers"
	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/bing"
}

func Routes() *bone.Mux {
	tmp := bone.New()
	tmp.GetFunc("/", controllers.BingHome)
	tmp.GetFunc("/home", controllers.BingHome)
	tmp.GetFunc("/about", controllers.BingAbout)
	tmp.GetFunc("/daily/photo", controllers.BingDailyPhoto)
	tmp.GetFunc("/daily/photo/embed", controllers.BingDailyPhotoEmbed)
	tmp.GetFunc("/daily/photo/raw", controllers.BingDailyPhotoRaw)

	tmp.GetFunc("/daily/photo/svga", controllers.BingDailyPhotoSVGA)   //800x600
	tmp.GetFunc("/daily/photo/xga", controllers.BingDailyPhotoXGA)     //1024x768
	tmp.GetFunc("/daily/photo/wxga", controllers.BingDailyPhotoWXGA)   //1280x720
	tmp.GetFunc("/daily/photo/hd", controllers.BingDailyPhotoHD)       //1366x768
	tmp.GetFunc("/daily/photo/fhd", controllers.BingDailyPhotoFHD)     //1920x1080
	tmp.GetFunc("/daily/photo/qhd", controllers.BingDailyPhotoQHD)     //2560x1440
	tmp.GetFunc("/daily/photo/wqxga", controllers.BingDailyPhotoWQXGA) //2560x1600
	tmp.GetFunc("/daily/photo/uhd", controllers.BingDailyPhotoUHD)     //3840x2160

	tmp.GetFunc("/daily/photo/sized/:x/:y", controllers.BingDailyPhotoSized)
	return tmp
}
