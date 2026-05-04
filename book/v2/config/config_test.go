package config_test

import (
	"fmt"
	"go18/book/v2/config"
	"os"
	"testing"
)

func TestLoadConfigFromYaml(t *testing.T) {
	err := config.LoadConfigFromYaml(fmt.Sprintf("%sC:\\Users\\flyfl\\Desktop\\go_18\\book\\v2\\application.toml", os.Getenv("workspaceFolder")))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("DATABASE_HOST", "DATABASE_PORT")

	err := config.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}
