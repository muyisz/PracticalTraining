package util

import (
	"encoding/json"
	"log"

	"github.com/muyisz/PracticalTraining/constant"
)

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
