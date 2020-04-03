package serializers

import (
	"fmt"
	"encoding/json"
	"url_shortener/db"
)


type Payload struct {
	Url string `json:"url" binding:"required"`
}

func SerializeUrl(website *db.Website) []byte {
	if website.Id == 0 {
		return []byte("Not found\n")
	}
	return []byte(fmt.Sprintf("%s\n", website.Url))
}

func SerializeUrls(urlsMap map[int]*db.Website) []byte {
	websites := make([]*db.Website, len(urlsMap))
	count := 0
	for _, website := range urlsMap {
		websites[count] = website
		count += 1
	}
	ret, err := json.Marshal(websites)
	if err != nil {
		return []byte("Error occurred\n")
	}
	return []byte(fmt.Sprintf("%s\n", ret))
}
