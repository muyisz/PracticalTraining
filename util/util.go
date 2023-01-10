package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/muyisz/PracticalTraining/constant"
)

func ReceptionHTMLUrl(name string) string {
	url := fmt.Sprintf("%s%s", constant.ReceptionStorageDirectory, name)
	log.Printf("name:%+v,url:%+v", name, url)
	return url
}

func BackstageHTMLUrl(name string) string {
	url := fmt.Sprintf("%s%s", constant.BackstageStorageDirectory, name)
	log.Printf("name:%+v,url:%+v", name, url)
	return url
}

func PostForm(data []byte, body interface{}) error {
	err := json.Unmarshal(data, body)
	if err != nil {
		log.Printf("[PostForm] Unmarshal err,data:%+v,err:%+v", data, err)
		return err
	}
	return nil
}

func DBQueryErr(err error) bool {
	if err == nil || err.Error() == constant.ErrRecordNotFound {
		return true
	}
	return false
}
