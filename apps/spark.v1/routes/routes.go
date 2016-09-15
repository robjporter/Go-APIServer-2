package routes

import (
	"../controllers"
	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/spark"
}

func Routes() *bone.Mux {
	tmp := bone.New()

	tmp.GetFunc("/", controllers.SparkHome)
	tmp.GetFunc("/display", controllers.SparkDisplay)

	// ROOMS
	tmp.GetFunc("/rooms/list/all", controllers.SparkRoomsListAll)
	tmp.GetFunc("/rooms/name/:name", controllers.SparkRoomsCheckRoomExists)
	tmp.GetFunc("/rooms/id/:id", controllers.SparkRoomsGetRoomNameFromID)
	tmp.GetFunc("/rooms/create/:name", controllers.SparkRoomsCreateRoomByName)
	tmp.GetFunc("/rooms/change/:name/:name2", controllers.SparkRoomsChangeName)
	tmp.GetFunc("/rooms/delete/:name", controllers.SparkRoomsDeleteRoom)

	// PEOPLE
	tmp.GetFunc("/people/find/me", controllers.SparkPeopleFindMe)
	tmp.GetFunc("/people/find/email/:email", controllers.SparkPeopleFindEmail)
	tmp.GetFunc("/people/find/name/:name", controllers.SparkPeopleFindName)
	tmp.GetFunc("/people/find/email/:email/id", controllers.SparkPeopleFindEmailReturnID)
	tmp.GetFunc("/people/find/name/:name/id", controllers.SparkPeopleFindNameReturnID)
	tmp.GetFunc("/people/find/email/:email/emails", controllers.SparkPeopleFindEmailReturnEmails)
	tmp.GetFunc("/people/find/name/:name/emails", controllers.SparkPeopleFindNameReturnEmails)
	tmp.GetFunc("/people/find/email/:email/name", controllers.SparkPeopleFindEmailReturnName)
	tmp.GetFunc("/people/find/name/:name/name", controllers.SparkPeopleFindNameReturnName)
	tmp.GetFunc("/people/find/email/:email/avatar", controllers.SparkPeopleFindEmailReturnAvatar)
	tmp.GetFunc("/people/find/name/:name/avatar", controllers.SparkPeopleFindNameReturnAvatar)
	tmp.GetFunc("/people/find/email/:email/created", controllers.SparkPeopleFindEmailReturnCreated)
	tmp.GetFunc("/people/find/name/:name/created", controllers.SparkPeopleFindNameReturnCreated)

	// MEMBERSHIP
	tmp.GetFunc("/membership/find/room/:id", controllers.SparkMembershipByRoomID)
	tmp.GetFunc("/membership/find/person/:id", controllers.SparkMembershipByPersonID)
	tmp.GetFunc("/membership/find/email/:id", controllers.SparkMembershipByEmail)

	return tmp
}
