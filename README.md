# Clumio GO SDK

## Overview

The `clumiogosdk` GO package provides an object-oriented API which allows developers to
write software using operations which Clumio provides for protecting data. This document provides
information on how to build and use the SDK.

## Requirements

The library requires GO 1.16 and higher. Third-party libraries are also required.

### Install
```
go get github.com/clumio-code/clumiogosdk
```

### Quick Start
```
import (
	"github.com/clumio-code/clumiogosdk/config"
	"github.com/clumio-code/clumiogosdk/controllers"
	"github.com/clumio-code/clumiogosdk/models"
)

func main() {
	cfg := config.Config{
		Token: <access_token>,
		// BaseUrl can be API URL of any namespace
		BaseUrl: <base_url>,
	}
	handler := controllers.NewPolicyDefinitionsV1(cfg)
	res, cerr := handler.ListPolicyDefinitions(nil, nil)
}
```
The REST API documentation describes all the available APIs and can be accessed from the help section in the top right corner of the Clumio UI.