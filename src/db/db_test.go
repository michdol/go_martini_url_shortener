package db

import (
	"testing"
)


func TestGet(t *testing.T) {
	var ret *Website
	website := &Website{
		Url: "asd",
	}

	AppDatabase.Add(website)
	ret = AppDatabase.Get(1)
	if ret.Id != 1 || ret.Url != "asd" {
		t.Errorf("Expected Id: 1, got: %d", ret.Id)
	}

	AppDatabase.Add(website)
	ret = AppDatabase.Get(2)
	if ret.Id != 2 || ret.Url != "asd" {
		t.Errorf("Expected Id: 2, got: %d", ret.Id)
	}

	retList := AppDatabase.GetAll()
	l := len(retList)
	if l != 2 {
		t.Errorf("Expected length to be 3, got: %d", l)
	}
}