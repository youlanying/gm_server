package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func DBSnapshot(action string, token string, tag int64, dbname string) (bool, error) {
	v := url.Values{}
	v.Set("token", token)
	v.Set("tag", fmt.Sprintf("%d", tag))
	v.Set("dbname", dbname)
	resp, err := http.PostForm(action, v)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(respBody) <= 0 {
		return false, err
	}

	type dumpRes struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}
	res := dumpRes{}
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return false, err
	}
	if res.Code == 1 {
		return true, nil
	}
	return false, errors.New(res.Error)
}
