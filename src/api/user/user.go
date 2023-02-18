package main

import (
	"net/http"

	"../stamping"

	"github.com/julienschmidt/httprouter"
)

// PUT: /profile/modify/password/:uid:old_password
func put_change_password(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if CheckHttpTimeStamp(time.now, )
}
