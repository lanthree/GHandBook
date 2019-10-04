package handlers

import (
	"encoding/json"
	"handbook/entity"
	"log"
	"net/http"
)

// NotFoundHandler 是命中不到其他路由的兜底展示
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	out := &entity.ErrorResult{Ret: 404, Reason: "Not Found"}
	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(b)
}
