package auth

import (
	"regexp"
	"time"

	"github.com/julienschmidt/httprouter"
)

func check_uid_legality(uid string) bool {
	// 需要确保uid没有非法字符（如../）。
	re := regexp.MustCompile("^[A-Za-z0-9_]$")
	return re.MatchString(uid)
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
