package config

import "strings"

type IHttpConfig interface {
	Port() string
	PathPrefix() string
	IgnoreLogUrls() []string
}

type HttpConfig struct{}

func (*HttpConfig) Port() string {
	return GetEnv("PORT", "8080")
}

func (*HttpConfig) PathPrefix() string {
	return GetEnv("PATH_PREFIX", "/")
}

func (*HttpConfig) IgnoreLogUrls() []string {
	ignoreLogsUrl := GetEnv("IGNORE_LOG_URLS", "")
	if ignoreLogsUrl == "" {
		return []string{}
	}

	return strings.Split(ignoreLogsUrl, ",")
}
