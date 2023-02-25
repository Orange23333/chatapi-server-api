package api

import (
	"net/http"
	"regexp"
	"time"

	"github.com/julienschmidt/httprouter"
)

func check_uname_legality(userName string) bool {
	// 需要确保UserName没有非法字符（如../）。
	re := regexp.MustCompile("^[A-Za-z0-9_]*$")
	return re.MatchString(userName)
}

func check_uid_legality(userId string) bool {
	// 需要确保UserId没有非法字符（如../）。
	re := regexp.MustCompile("^[1-9][0-9]*$|^0$")
	if !re.MatchString(userId) {
		return false
	}
	
	value :=
	return value >= 0
}

func check_pass_token(ps httprouter.Params, receive_time time.Time) bool {
	uid := ps.ByName("uid")
	pass_token := ps.ByName("pass_token")
	time_stamp := ps.ByName("time_stamp")

	loc, _ := time.LoadLocation("UTC")
	time_stamp_time, err := time.ParseInLocation("2006-01-02 15:04:05", time_stamp, loc)
	if err != nil {
		return false
	}

	if !check_time_stamp(time_stamp_time, receive_time, 10) {
		return false
	}

	if !check_uid_legality(uid) {
		return false
	}

	if uid == "test" {
		if pass_token == "7thjo9876thjko" {
			return true
		}
	}

	return false
}

// GET: /login/get-pass-token/:user_id:password_sha256:verfiy_code_answer:time_stamp
func Get_Login_GetPassToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// DETELE: /logout/destory-pass-token/:uid:pass_token:time_stamp
func Get_Logout_DestoryPassToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// GET: /verfiy-code/get-one/:verfiy_code_token:time_stamp"
func Get_VerfiyCode_GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//Return a html block
}
