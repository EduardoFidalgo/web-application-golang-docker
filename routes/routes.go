package routes

import (
	"net/http"

	"github.com/webservice-golang-api/controller"
)

func Routes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/form", controller.StoreView)
	http.HandleFunc("/store", controller.Store)
	http.HandleFunc("/destroy", controller.Destroy)
	http.HandleFunc("/edit", controller.UpdateView)
	http.HandleFunc("/update", controller.Update)
}
