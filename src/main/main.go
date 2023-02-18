package main

import (
	"bufio"
	"chatapi/server/api/src/stamping"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/julienschmidt/httprouter"
)

var ai_model_list = []string{"gpt-j-6b"}

// Error Codes and Error Messages:
// ERC = ERror.Code, EDM = Error.DebugMessage, EPM = Error.PublicMessage
const ERC_APIRequestLostParameter int = 40001
const EDM_APIRequestLostParameter string = "API request lost parameter %s. Please check your code or view the API document."
const EPM_APIRequestLostParameter string = EDM_APIRequestLostParameter

func find_string(a []string, x string) int {
	i := sort.SearchStrings(a, x)
	if a[i] == x {
		return len(a)
	} else {
		return -1
	}
}

func get_file_size(file *os.File) int64 {
	origin_pos := file.Seek(0, os.SEEK_CUR)
	end_pos := file.Seek(0, os.SEEK_END)
	file.Seek(origin_pos, os.SEEK_SET)
	return end_pos
}

func read_all(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file_size := get_file_size(file)

	read := bufio.NewReader(file)
	content, err := read.ReadString(file_size)
	if err != nil {
		return nil, err
	}

	return content, nil
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

func get_index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content, err := read_all("/pages/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, content)
}

func get_ai_model_list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json_bytes, e := json.Marshal(ai_model_list)
	if e == nil {
		fmt.Fprint(w, string(json_bytes)+"\n")
	} else {
		fmt.Fprint(w, e.Error()+"\n")
	}
	w.WriteHeader()
}

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

func main() {
	sort.Strings(ai_model_list)

	stampingHandler := stamping.New()
	router := httprouter.New()
	stampingHandler.subHandler = router

	router.GET("/", get_index)

	router.POST("/register/apply-access/test-1/:user_id:password_sha")

	router.GET("/login/get-pass-token/:user_id:password_sha256:verfiy_code_answer:time_stamp")
	router.GET("/verfiy-code/get-one/:verfiy_code_token:time_stamp") //Return a html block
	router.DELETE("/logout/:uid:pass_token:time_stamp")

	router.GET("/ai-models/list", get_ai_model_list)

	if find_string(ai_model_list, "gpt-j-6b") < len(ai_model_list) {
		//之后这一步由各个模块自己完成
		router.POST("/modules/gpt-j-6b/requests/new/:text:uid:pass_token:time_stamp", post_new_request)
		router.GET("/modules/gpt-j-6b/requests/status/:request_id:uid:pass_token:time_stamp", get_request_status)
		router.DELETE("/modules/gpt-j-6b/requests/cannel/:request_id:uid:pass_token:time_stamp", delete_cannel_request)
		router.GET("/modules/gpt-j-6b/requests/result/:request_id:uid:pass_token:time_stamp", get_request_result)
	}

	log.Fatal(http.ListenAndServe(":22388", stampingHandler))
}
