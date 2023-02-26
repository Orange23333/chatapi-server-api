package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var AiModelList = []string{"gpt-j-6b"}

// GET: /ai-models/list
func GetAiModelList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json_bytes, e := json.Marshal(AiModelList)

	if e == nil {
		fmt.Fprint(w, string(json_bytes)+"\n")

		w.WriteHeader(200)
	} else {
		fmt.Fprint(w, e.Error()+"\n")

		w.WriteHeader(500)
	}
}
