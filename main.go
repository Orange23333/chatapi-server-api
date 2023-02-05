package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/julienschmidt/httprouter"
)

var ai_model_list = []string{"gpt-j-6b"}

// Error Codes and Error Messages:
// ERC = ERror.Code, EDM = Error.DebugMessage, EPM = Error.PublicMessage
const ERC_APIRequestLostParameter int = 40001
const EDM_APIRequestLostParameter string = "API request lost parameter %s. Please check your code or view the API document."
const EPM_APIRequestLostParameter string = EDM_APIRequestLostParameter

func search_string(a []string, x string) int {
	i := sort.SearchStrings(a, x)
	if a[i] == x {
		return len(a)
	} else {
		return -1
	}
}

func write_all(path string, text string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	//write := bufio.NewWriter(file)
	//write.WriteString(text)
	//write.Flush()

	return nil
}

func check_pass_token(ps httprouter.Params) bool {
	uid := ps.ByName("uid")
	pass_token := ps.ByName("pass_token")
	time_stamp := ps.ByName("time_stamp")

	return true
}

func get_index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<!DOCTYPE HTML><html><head><title>ChatAPI</title></head><body>Welcome to use ChatAPI (Testing)! View <a href=\"https://www.ourorangenet.com/wiki/index.php/ChatAPI\">ChatAPI - 沃社Wiki</a> for help.</body></html>\n")
}

func get_ai_model_list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json_bytes, e := json.Marshal(ai_model_list)
	if e == nil {
		fmt.Fprint(w, string(json_bytes)+"\n")
	} else {
		fmt.Fprint(w, e.Error()+"\n")
	}
}

var new_request_count int64 = 0

func post_new_request(w http.ResponseWriter, r *http.Request, ps httprouter.Params) int64 {

	// 因此需要确保uid没有非法字符（如../）。

	err := write_all(fmt.Sprintf("./requests/%s.%d.request", ps.ByName("uid"), new_request_count), ps.ByName("text"))

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

func main() {
	sort.Strings(ai_model_list)

	router := httprouter.New()

	router.GET("/", get_index)

	router.GET("/login/get-pass-token/:user_id:password_sha256:verfiy_code_answer")
	router.GET("/verfiy-code/get-image/:verfiy_code_token")
	router.DELETE("/logout/:uid:pass_token")

	if search_string(ai_model_list, "gpt-j-6b") < len(ai_model_list) {
		//之后这一步由各个模块自己完成
		router.POST("/modules/gpt-j-6b/requests/new/:text:uid:pass_token", post_new_request)
		router.GET("/modules/gpt-j-6b/requests/status/:request_id:uid:pass_token", get_request_status)
		router.DELETE("/modules/gpt-j-6b/requests/cannel/:request_id:uid:pass_token", delete_cannel_request)
		router.GET("/modules/gpt-j-6b/requests/result/:request_id:uid:pass_token", get_request_result)
	}

	log.Fatal(http.ListenAndServe(":22388", router))
}
