package persistence

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
)

type Persist interface {
	Save(k int, value string)
	Get(key string) (string, error)
	Update(k int, value string)
}

var notFoundErr = errors.New("key not found")

const (
	intro          = "\n"
	comma          = ","
	filePermission = 0666
)

type FilePersistence struct {
	File     string
	InMemory bool
}

func write(fileLocation, record string) error {
	//read a file if already exists, or create a new one
	file, err := os.OpenFile(filepath.Join(fileLocation), os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
		return err
	}

	defer internal.OnCloseError(file.Close)

	_, err = file.WriteString(record)

	return err
}

func updateValue(bytes []byte, k int, value, fileLocation string) ([]string, error) {
	pairs := parseBytes(bytes)
	var (
		found   bool
		entries []string
	)

	for _, kv := range pairs {
		values := strings.Split(kv, comma)
		hashedKey, _ := strconv.Atoi(values[0])

		if k == hashedKey {
			found = true
			entries = append(entries, fmt.Sprintf("%d,%s"+intro, k, value))
		}
	}

	if !found {
		return nil, notFoundErr
	}

	return entries, nil
}

func (f *FilePersistence) Update(k int, value string) {
	bytes, err := ioutil.ReadFile(f.File)
	file, _ := os.OpenFile(filepath.Join(f.File), os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)

	if err != nil {
		logs.Error("cannot update the value " + err.Error())
		return
	}

	res, err := updateValue(bytes, k, value, f.File)

	for _, kv := range res {
		file.WriteString(kv)
	}
}

func (f *FilePersistence) Save(k int, value string) {
	//read a file if already exists, or create a new one
	file, err := os.OpenFile(filepath.Join(f.File), os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
		return
	}

	defer internal.OnCloseError(file.Close)

	record := buildRecord(k, value, f.InMemory)

	_, err = file.WriteString(record)

	if err != nil {
		logs.LogError(err)
		// no error returned, just a shame
	}
}

func (f *FilePersistence) Get(key string) (string, error) {
	bytes, err := ioutil.ReadFile(f.File)

	if err != nil {
		return "", err
	}

	return getValue(bytes, key, 0, false)
}

func parseBytes(bytes []byte) []string {
	text := string(bytes)
	return strings.Split(text, intro)
}

func getValue(bytes []byte, strSearchKey string, hashedSearchKey int, hashed bool) (string, error) {
	pairs := parseBytes(bytes)

	for i, j := 0, len(pairs)-1; i < j; i, j = i+1, j-1 {
		pairs[i], pairs[j] = pairs[j], pairs[i]
	}

	x, pairs := pairs[0], pairs[1:]
	pairs = append(pairs, x)

	for _, kv := range pairs {
		values := strings.Split(kv, comma)
		hashedKey := values[0]

		// cannot read empty keys, means EOF
		if hashedKey == "" {
			break
		}

		i, err := strconv.Atoi(hashedKey)

		if err != nil {
			logs.Error("cannot parse key into INT type. " + err.Error())
			return "", nil
		}

		// we can receive the search eky as string in plain text or the hashcode
		var searchKey int
		if hashed {
			searchKey = hashedSearchKey
		} else {
			searchKey = internal.ApplyHash(strSearchKey)
		}

		if i == searchKey {
			return values[1], nil
		}
	}

	return "", notFoundErr
}

func buildRecord(k int, value string, inMemory bool) string {
	if inMemory {
		return fmt.Sprintf("%d,%s"+intro, k, value)
	} else {
		return fmt.Sprintf("%d,%s,%d"+intro, k, value, time.Now().Unix())
	}
}
