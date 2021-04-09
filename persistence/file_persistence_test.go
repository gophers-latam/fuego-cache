package persistence

import (
	"testing"

	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
)

var filePersistence = FilePersistence{
	File:     "./test.csv",
	InMemory: true,
}

func TestFilePersistence_Get(t *testing.T) {
	key := internal.ApplyHash("123")
	filePersistence.Save(key, "test")

	res, err := filePersistence.Get("123")

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	logs.Info(res)
}

func TestFilePersistence_Update(t *testing.T) {
	key := internal.ApplyHash("1234")

	filePersistence.Save(key, "test1")
	filePersistence.Update(key, "test2")

	res, err := filePersistence.Get("1234")

	if err != nil {
		t.Error(err)
		t.Fail()
	} else if res != "test2" {
		t.Fail()
	}

	logs.Info(res)
}
