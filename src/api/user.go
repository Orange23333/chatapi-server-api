package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PUT: /profile/modify/password/:uid:old_pswd:new_pswd:time_stamp
func Put_ChangePassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
