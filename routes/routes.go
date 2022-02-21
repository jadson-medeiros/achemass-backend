package routes

import (
	"net/http"

	"github.com/jadson-medeiros/achemass-backend/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
}
