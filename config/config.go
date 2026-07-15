// Copyright (c) 2023 Clumio All Rights Reserved

// Package config has the types required for configuration
package config

import "time"

type Config struct {
	Token                     string
	BaseUrl                   string
	OrganizationalUnitContext string
	CustomHeaders             map[string]string
	Timeout                   time.Duration
}
