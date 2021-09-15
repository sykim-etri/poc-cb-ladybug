package mcks

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
	"github.com/cloud-barista/cb-mcas/pkg/utils/config"
	"github.com/go-resty/resty/v2"

	"github.com/sirupsen/logrus"
)

type (
	McksStatus struct {
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Message string `json:"message"`
	}
	Mcks struct {
		Model
	}
)

func NewMcks(namespace string) *Mcks {
	return &Mcks{
		Model: Model{namespace: namespace},
	}
}

type Model struct {
	namespace string
}

func (self *Model) execute(method string, url string, body interface{}, result interface{}) (bool, error) {

	// validation
	if err := self.validate(validation.Validation{}); err != nil {
		return false, err
	}

	resp, err := app.ExecuteHTTP(method, *config.Config.McksUrl+url, body, result)
	if err != nil {
		return false, err
	}

	// response check
	if resp.StatusCode() > 300 && resp.StatusCode() != http.StatusNotFound {
		logrus.Warnf("MCKS: statusCode=%d, url=%s, body=%s", resp.StatusCode(), resp.Request.URL, resp)
		status := McksStatus{}
		json.Unmarshal(resp.Body(), &status)
		/*
			// message > message 로 리턴되는 경우가 있어서 한번더 unmarshal 작업
			if json.Valid([]byte(status.Message)) {
				json.Unmarshal([]byte(status.Message), &status)
			}
		*/
		return false, errors.New(status.Message)
	}

	if method == http.MethodGet && resp.StatusCode() == http.StatusNotFound {
		logrus.Infof("Not found data (status=404, method=%s, url=%s)", method, url)
		return false, nil
	}

	return true, nil
}

func (self *Model) validate(valid validation.Validation) error {
	valid.Required(self.namespace, "namespace")
	//valid.Required(self.Name, "name")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errors.New(fmt.Sprintf("[%s]%s", err.Key, err.Error()))
		}
	}
	return nil
}

// status :200 , body = {message: "Cannot find ..." }  형태의 response 에러처리
func (self *Model) hasResponseMessage(resp *resty.Response) error {
	var d map[string]interface{}
	json.Unmarshal(resp.Body(), &d)
	if d["message"] != nil {
		return errors.New(fmt.Sprintf("%s", d["message"]))
	}
	return nil
}
