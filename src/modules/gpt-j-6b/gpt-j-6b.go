package gpt_j_6b

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var new_request_count int64 = 0

func post_new_request(w http.ResponseWriter, r *http.Request, ps httprouter.Params) int64 {
	text := ps.text
	uid := ps.ByName("uid")
	pass_token := ps.ByName("pass_token")

	if check_pass_token(ps, time.Now()) {
		err_json := map[string]string{}
		fmt.Fprint(w, string(json)+"\n")
	}

	err := write_all(fmt.Sprintf("./requests/%s.%d.request", ps.ByName("uid"), new_request_count), ps.ByName("text"))
	if err != nil {
		w.WriteHeader(500)
	}

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func get_request_status(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func delete_cannel_request(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

func get_request_result(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}
