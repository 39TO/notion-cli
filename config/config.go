package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
)

type env struct {
	NOTION_TOKEN        string
	NOTION_ROOT_PAGE_ID string
}

var _env env

func init() {
	var err error
	var envErrs []error

	if godotenv.Load() != nil {
		fmt.Println("No .env file found")
	}

	if _env.NOTION_TOKEN, err = getEnv("NOTION_TOKEN"); err != nil {
		envErrs = append(envErrs, err)
	}
	if _env.NOTION_ROOT_PAGE_ID, err = getEnv("NOTION_ROOT_PAGE_ID"); err != nil {
		envErrs = append(envErrs, err)
	}

	if len(envErrs) != 0 {
		var errMsg string
		for _, envErr := range envErrs {
			errMsg += fmt.Sprintf("%s\n", envErr.Error())
		}
		panic(errMsg)
	}
}

func GetNotionToken() string {
	return _env.NOTION_TOKEN
}

func GetNotionRootPageID() string {
	return _env.NOTION_ROOT_PAGE_ID
}

func getEnv(key string) (string, error) {
	e := os.Getenv(key)
	if e == "" {
		return "", xerrors.Errorf("getEnv: %s not found", key)
	}
	return e, nil
}
