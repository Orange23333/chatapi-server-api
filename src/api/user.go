package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PUT: /profile/modify/password/:uid:old_pswd:new_pswd:time_stamp
func Put_ChangePassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// PUT: /register/apply-access/test-1/:user_id:password"
func Put_Register_ApplyAccess_Test1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
