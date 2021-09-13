# Clumio GO SDK

## Overview

The `clumiogosdk` GO package provides an object-oriented API which allows developers to
write software using operations which Clumio provides for protecting data. This document provides
information on how to build and use the SDK.

## Requirements

The library requires GO 1.16 and higher. Third-party libraries are also required.

### Install
```
go get github.com/clumio/clumiogosdk
```

### Quick Start
```
import (
	"github.com/clumio/clumiogosdk/config"
	"github.com/clumio/clumiogosdk/controllers"
	"github.com/clumio/clumiogosdk/models"
)

func main() {
	cfg := config.Config{
		Token: <access_token>,
		// BaseUrl can be API URL of any namespace
		BaseUrl: "https://eng-48-cc-1-backend.api.k8s-clumio.com/api",
	}
	handler := controllers.NewPolicyDefinitionsV1(cfg)
	res, cerr := handler.ListPolicyDefinitions(nil, nil)
}
```

## Documentation
For full documentation, including a quick start example, see [SDK documentation](https://probable-telegram-b0cf4eee.pages.github.io)
