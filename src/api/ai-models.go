package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Ai_Model_List = []string{"gpt-j-6b"}

func get_ai_model_list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json_bytes, e := json.Marshal(Ai_Model_List)

	if e == nil {
		fmt.Fprint(w, string(json_bytes)+"\n")

		w.WriteHeader(200)
	} else {
		fmt.Fprint(w, e.Error()+"\n")

		w.WriteHeader(500)
	}
}
