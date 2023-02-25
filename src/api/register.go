package api

import (
	"chatapi/server/api/main"
	"chatapi/server/api/stamping"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PUT: /register/apply-access/test-1/:user_id:password"
func Put_Register_ApplyAccess_Test1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !stamping.TimoutHandler(stamping.DEFAULT_TIMEOUT_SEC, w, r, ps) {
		return
	}

	userName := ps.ByName("uname")
	password := ps.ByName("pswd")
	email := ps.ByName("email")
	comment := ps.ByName("comment")

	if check_uname_legality(userName) {
		w.Write([]byte(main.MessagesToListViewJson([]string{
			"Username is illegal!",
		}, 1)))
		w.WriteHeader(400)
		return
	}
}
