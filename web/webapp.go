package web

import "net/http"

func Webapp() {

	http.ListenAndServe(":8000", nil)

}
