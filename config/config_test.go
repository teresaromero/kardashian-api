package config_test

import (
	"errors"
	"kardashian_api/config"
	"testing"
)

func TestConfig_NoFileLoaded(t *testing.T) {
	err := config.LoadEnv(".env.example")
	expected := errors.New("open .env.example: no such file or directory")
	if err.Error() != expected.Error() {
		t.Errorf("Expected %v but got %v", expected, err)
	}
}

func TestConfig_FileLoaded(t *testing.T) {
	err := config.LoadEnv("../.env.example")
	if err != nil {
		t.Errorf("Expected %v but got %v", nil, err)
	}
}

func TestConfig_ExportedVars(t *testing.T) {
	err := config.LoadEnv("../.env.example")
	if err != nil {
		t.Errorf("Expected %v but got %v", nil, err)
	}
	actual := config.MongoURI
	expected := "mongodb://host:port"
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
