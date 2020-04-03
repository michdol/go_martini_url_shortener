package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/go-martini/martini"
	"url_shortener/db"
	"url_shortener/serializers"
)


func HomeHandler(database db.DB, req *http.Request) string {
	return "Hello world\n"
}

func AllUrlsHandler(database db.DB, req *http.Request) []byte {
	urls := database.GetAll()
	return serializers.SerializeUrls(urls)
}

func GetUrlHandler(database db.DB, params martini.Params) []byte {
	pk, err := strconv.Atoi(params["pk"])
	if err != nil {
		return []byte(fmt.Sprintf("Invalid parameter %s, integer required\n", params["pk"]))
	}
	website := database.Get(pk)
	return serializers.SerializeUrl(website)
}

func AddUrlHandler(database db.DB, payload serializers.Payload) string {
	var url string = payload.Url
	id, _ := database.Add(&db.Website{
		Url: url,
	})
	return fmt.Sprintf("%d\n", id)
}