package handlers

import (
	"encoding/json"
	"handbook/datamanager"
	"handbook/entity"
	"log"
	"net/http"
)

// RandomCardsHandler 随机返回5个卡片内容
func RandomCardsHandler(w http.ResponseWriter, r *http.Request) {

	result := entity.RandomCardsOKResult{}
	result.Ret = 0
	result.Cards = append(result.Cards, datamanager.GetRomdonCards()...)

	out := &result
	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(b)
}
