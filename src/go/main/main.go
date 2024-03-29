package main

import (
	"bufio"
	"chatapi/servers/api-server/api"
	"chatapi/servers/api-server/stamping"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/julienschmidt/httprouter"
)

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
	origin_pos, _ := file.Seek(0, os.SEEK_CUR)
	end_pos, _ := file.Seek(0, os.SEEK_END)
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

func ErrorsToListViewJson(errs []error, n int) string {
	r := "{\"view\":\"list\",\"length\":" + string(n) + ",\"items\":["
	if n > 0 {
		r += "\"" + errs[0].Error() + "\""
	}
	for i := 1; i < n; i++ {
		r += ",\"" + errs[i].Error() + "\""
	}
	r += "]}"
	return r
}

func MessagesToListViewJson(messages []string, n int) string {
	r := "{\"view\":\"list\",\"length\":" + string(n) + ",\"items\":["
	if n > 0 {
		r += "\"" + messages[0] + "\""
	}
	for i := 1; i < n; i++ {
		r += ",\"" + messages[i] + "\""
	}
	r += "]}"
	return r
}

func main() {
	os.MkdirAll(api.UserProfilesDir, 0755)
	os.MkdirAll("./data/access-application/test-1", 0755)
	os.MkdirAll("./data/")

	router := httprouter.New()
	stamping := stamping.New(router)

	router.GET("/", get_index) //Use `router.ServeFiles` insteads!

	router.GET("/ai-models/list", api.GetAiModelList)

	if find_string(api.AiModelList, "gpt-j-6b") < len(api.AiModelList) {
		//之后这一步由各个模块自己完成
		router.POST("/modules/gpt-j-6b/requests/new/:text:uid:pass_token:time_stamp", post_new_request)
		router.GET("/modules/gpt-j-6b/requests/status/:request_id:uid:pass_token:time_stamp", get_request_status)
		router.DELETE("/modules/gpt-j-6b/requests/cannel/:request_id:uid:pass_token:time_stamp", delete_cannel_request)
		router.GET("/modules/gpt-j-6b/requests/result/:request_id:uid:pass_token:time_stamp", get_request_result)
	}

	log.Fatal(http.ListenAndServe(":22388", stamping))
}
