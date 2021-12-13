# CB-Ladybug :beetle:
> Multi-Cloud Application Management Framework

[![Go Report Card](https://goreportcard.com/badge/github.com/cloud-barista/poc-cb-ladybug)](https://goreportcard.com/report/github.com/cloud-barista/poc-cb-ladybug)
[![Build](https://img.shields.io/github/workflow/status/cloud-barista/poc-cb-ladybug/Build%20amd64%20container%20image)](https://github.com/cloud-barista/poc-cb-ladybug/actions?query=workflow%3A%22Build+amd64+container+image%22)
[![Top Language](https://img.shields.io/github/languages/top/cloud-barista/poc-cb-ladybug)](https://github.com/cloud-barista/poc-cb-ladybug/search?l=go)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cloud-barista/poc-cb-ladybug?label=go.mod)](https://github.com/cloud-barista/poc-cb-ladybug/blob/master/go.mod)
[![Repo Size](https://img.shields.io/github/repo-size/cloud-barista/poc-cb-ladybug)](#)
[![GoDoc](https://godoc.org/github.com/cloud-barista/poc-cb-ladybug?status.svg)](https://pkg.go.dev/github.com/cloud-barista/poc-cb-ladybug@master)
[![Release Version](https://img.shields.io/github/v/release/cloud-barista/poc-cb-ladybug?color=blue)](https://github.com/cloud-barista/poc-cb-ladybug/releases/latest)
[![License](https://img.shields.io/github/license/cloud-barista/poc-cb-ladybug?color=blue)](https://github.com/cloud-barista/poc-cb-ladybug/blob/master/LICENSE)

```
[NOTE]
CB-Ladybug is currently under development. (the latest release is none) 
So, we do not recommend using the current release in production.
Please note that the functionalities of CB-Ladybug are not stable and secure yet.
If you have any difficulties in using CB-Ladybug, please let us know.
(Open an issue or Join the cloud-barista Slack)
```

## Getting started

### Preparation

- Golang 1.16.+ ([Download and install](https://golang.org/doc/install))

### Dependencies

- CB-MCKS [v0.4.5](https://github.com/cloud-barista/cb-mcks/releases/tag/v0.4.5)
- CB-Tumblebug [v0.4.17](https://github.com/cloud-barista/cb-tumblebug/releases/tag/v0.4.17)
- CB-Spider [v0.4.18](https://github.com/cloud-barista/cb-spider/releases/tag/v0.4.18)


### Clone

- Clone source code
  ```bash
  git clone https://github.com/cloud-barista/poc-cb-ladybug.git cb-ladybug
  cd cb-ladybug
  go get -v all
  ```

### Run 

- Setup environment variables

  Check and modify environment variables in `cb-ladybug/conf/setup.env`
  ```bash
  cat conf/setup.env
  source conf/setup.env
   ```

- Run CB-Ladybug
  ```bash
  go run cmd/cb-ladybug/main.go
   ```

### Build and Execute

- Build CB-Ladybug
  ```bash
  cd $(APP_ROOT)
  make
  ```

- Execute CB-Ladybug
  ```bash
  make run
  ```

### API documentation

- Under construction

## Documents

- Under construction


## Contribution
Learn how to start contribution on the [Contributing Guideline](https://github.com/cloud-barista/docs/tree/master/contributing) and [Style Guideline](https://github.com/cloud-barista/poc-cb-ladybug/blob/master/STYLE_GUIDE.md)

### Test

- Check health status

  ```bash
  ./scripts/get-health.sh

  [INFO]
  - Ladybug URL is 'http://localhost:1592/ladybug'

  ------------------------------------------------------------------------------
  cloud-barista cb-ladybug is alived
  ```

- Run CB-Spider & CB-Tumblebug

  Reference to https://github.com/cloud-barista/cb-tumblebug#3-cb-tumblebug-%EC%8B%A4%ED%96%89

- Run CB-MCKS

  Reference to https://github.com/cloud-barista/cb-mcks#run

- Run chartmuseum in local
  ```
  cd $APP_ROOT/scripts
  ./run_chartmuseum.sh
  ```
  You can access your chartmusem at http://localhost:38080

  If you cannot download a chartmuseum package,
  you can get the package in https://github.com/helm/chartmuseum/releases

- Register Cloud Connection Info

  Reference to https://github.com/cloud-barista/cb-mcks/tree/master/docs/test#cloud-connection-info-%EB%93%B1%EB%A1%9D

  Update hard-coded cloud connection configs as your cloud connection configs in `cb-ladybug/pkg/core/service/mcas.go:makeClusterReq()`

- Create `lb-ns` namespace

  ```bash
  cd <CB-MCKS directory>/docs/test
  ./ns-create.sh lb-ns
  ```

- Enable MCAS
  ```bash
  cd $APP_ROOT/scripts
  ./enable-mcas.sh lb-ns
  ```

- Upload sample packages
  ```bash
  cd $APP_ROOT/scripts
  ./upload_package.sh lb-ns ./sample/chartmuseum-3.1.0.tgz
  ./upload_package.sh lb-ns ./sample/chartmuseum-3.2.0.tgz
  ./upload_package.sh lb-ns ./sample/nginx-9.5.12.tgz
  ```

- Install an app(ex. nginx) instance
  ```bash
  ./install-app-instance.sh lb-ns nginx-01 nginx ./sample/nginx-values.yaml
  ```

- Uninstall the app instance
  ```bash
  ./uninstall-app-instance.sh lb-ns nginx-01
  ```

- Delete the package
  ```bash
  ./delete-package.sh lb-ns nginx 9.5.12
  ```

- Disable MCAS
  ```bash
  cd $APP_ROOT/scripts
  ./disable-mcas.sh lb-ns
  ```
