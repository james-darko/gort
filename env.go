package gort

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
)

func Env(name string) (string, bool) {
	v := os.Getenv(name)
	if v == "" {
		return "", false
	}
	return v, true
}

func MustEnv(name string) string {
	v, ok := Env(name)
	if !ok {
		panic("required env var " + name + " is not set")
	}
	return v
}

func MustJsonEnv(name string, v any) {
	vStr, ok := Env(name)
	if !ok {
		panic("required env var " + name + " is not set")
	}
	switch vStr[0] {
	case '{', '[':
		err := json.Unmarshal([]byte(vStr), v)
		if err != nil {
			panic("could not unmarshal required config " + name + ": " + err.Error())
		}
	default:
		r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(vStr))
		err := json.NewDecoder(r).Decode(v)
		if err != nil {
			panic("could not unmarshal required config " + name + ": " + err.Error())
		}
	}
}

func AssertEnv(name string) string {
	v, ok := Env(name)
	Assertf(ok, "%s env var is required", name)
	return v
}

func AssertJsonEnv(name string, v any) {
	vStr, ok := Env(name)
	Assertf(ok, "%s env var is required", name)
	switch vStr[0] {
	case '{', '[':
		err := json.Unmarshal([]byte(vStr), v)
		Assertf(err == nil, "could not unmarshal required config %s: %v", name, err)
	default:
		r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(vStr))
		err := json.NewDecoder(r).Decode(v)
		Assertf(err == nil, "could not unmarshal required config %s: %v", name, err)
	}
}
