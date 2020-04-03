package serializers

import (
	"bytes"
	"testing"
	"url_shortener/db"
)


func TestSerializeUrl(t *testing.T) {
	var ret []byte
	website := &db.Website{
		Id: 1,
		Url: "asd",
	}

	ret = SerializeUrl(website)
	expected := []byte("asd\n")

	if !bytes.Equal(ret, expected) {
		t.Errorf("Failed to serialize website, expected: %s, got: %s", expected, ret)
	}
}

func TestSerializeUrls(t *testing.T) {
	var ret []byte
	websites := map[int]*db.Website{
		1: &db.Website{
			Id: 1,
			Url: "asd",
		},
	}

	ret = SerializeUrls(websites)
	expected := []byte("[{\"Id\":1,\"Url\":\"asd\"}]\n")

	if !bytes.Equal(ret, expected) {
		t.Errorf("Failed to serialize website, expected: %s, got: %s", expected, ret)
	}
}
