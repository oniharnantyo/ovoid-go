package ovoid_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
	"html"
	"io/ioutil"
	"net/http"
	"reflect"
)

type (
	Headers struct {
		AppID string `json:"App-id"`
		AppVersion string `json:"App-version"`
		OSName string `json:"OS"`
		OSVersion string `json:"os_version"`
		MACAddress string `json:"mac_address"`
	}

	BodyRequest struct {
		DeviceID string `json:"deviceId, omitempty"`
		MobilePhone string `json:"mobile, omitempty"`
	}

	Requester struct {
		ExtraURL string
		Body BodyRequest
		Headers Headers
		IsJson bool
	}
)

func (r Requester)Post(ctx context.Context, response interface{})(error)  {
	requestBody, err := json.Marshal(r.Body)
	if err != nil {
		return err
	}

	url := cfg.BaseEndpoint + r.ExtraURL
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	val := reflect.ValueOf(r.Headers)
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		req.Header.Set(t.Field(i).Tag.Get("json"), val.Field(i).String())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rBody, response); err != nil {
		return err
	}

	value := reflect.ValueOf(response).Elem()
	for i := 0; i < value.NumField(); i++ {
		filed := value.Field(i)
		if filed.Type() != reflect.TypeOf("") {
			continue
		}
		str := filed.Interface().(string)
		filed.SetString(html.EscapeString(str))
	}

	valid, err := govalidator.ValidateStruct(response)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("Invalid Data")
	}

	return nil
}

func (r Requester)Get(ctx context.Context, response interface{})(error)  {
	url := cfg.BaseEndpoint + r.ExtraURL
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()

	bodyQuery := reflect.ValueOf(r.Body)
	for _, key := range bodyQuery.MapKeys() {
		value := bodyQuery.MapIndex(key)
		q.Add(key.String(), value.String())
	}

	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rBody, response); err != nil {
		return err
	}

	value := reflect.ValueOf(response).Elem()
	for i := 0; i < value.NumField(); i++ {
		filed := value.Field(i)
		if filed.Type() != reflect.TypeOf("") {
			continue
		}
		str := filed.Interface().(string)
		filed.SetString(html.EscapeString(str))
	}

	valid, err := govalidator.ValidateStruct(response)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("Invalid Data")
	}

	return nil
}