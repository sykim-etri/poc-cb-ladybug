package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/cb-ladybug/pkg/core/model"
	"github.com/cloud-barista/cb-ladybug/pkg/core/service"
	"github.com/labstack/echo/v4"
)

func ListApp(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	appInstanceList, err := service.ListAppInstance(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, appInstanceList)
}

func GetApp(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "app"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	appInstance, err := service.GetAppInstance(c.Param("namespace"), c.Param("app"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, appInstance)
}

func CreateApp(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	namespace := c.Param("namespace")

	mcas := model.NewMcas(namespace)

	status, err := mcas.GetStatus()
	if err != nil {
		common.CBLog.Error(err)
		return err
	}

	if status == model.STATUS_MCAS_DISABLED {
		common.CBLog.Infof("MCAS for namespace '%s' is disabled.\n", namespace)
		return nil
	}

	appInstanceReq := &model.AppInstanceReq{}
	reqData := c.FormValue("reqData")
	err = json.Unmarshal([]byte(reqData), &appInstanceReq.Data)
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	fhValues, err := c.FormFile("valuesFile")
	if fhValues != nil {
		fileValues, _ := fhValues.Open()
		defer fileValues.Close()
		bytesValues, err := ioutil.ReadAll(fileValues)
		if err != nil {
			common.CBLog.Error(err)
			return SendMessage(c, http.StatusBadRequest, err.Error())
		}

		appInstanceReq.ValuesYaml = string(bytesValues)
	}

	model.AppInstanceReqDef(appInstanceReq)

	err = model.AppInstanceReqValidate(appInstanceReq)
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	appInstance, err := service.CreateAppInstance(c.Param("namespace"), appInstanceReq)
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusInternalServerError, err.Error())
	}

	return Send(c, http.StatusOK, appInstance)
}

func DeleteApp(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "app"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	status, err := service.DeleteAppInstance(c.Param("namespace"), c.Param("app"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusInternalServerError, err.Error())
	}

	return Send(c, http.StatusOK, status)
}
