package test

import (
	"orders/internal/dotenv"
	"os"
	"testing"
)

func TestDotenv(t *testing.T) {
	dotenv := dotenv.Config()
	test := dotenv.Get("TEST")
	if test != "test" {
		t.Errorf("Expected test to be test, got %s", test)
	}

	test2 := dotenv.Get("EMPTY")
	if test2 != "" {
		t.Errorf("Expected test2 to be empty, got %s", test2)
	}
}

func TestDotenvOverride(t *testing.T) {
	dotenv := dotenv.Config()
	test := dotenv.Get("OVERRIDE")
	if test != "overridden" {
		t.Errorf("Expected test to be overridden, got %s", test)
	}
}

func TestDotenvOsOverride(t *testing.T) {
	_ = dotenv.Config()
	test := os.Getenv("TEST")
	if test != "test" {
		t.Errorf("Expected test to be test, got %s", test)
	}

	test2 := os.Getenv("EMPTY")
	if test2 != "" {
		t.Errorf("Expected test2 to be empty, got %s", test2)
	}

	override := os.Getenv("OVERRIDE")
	if override != "overridden" {
		t.Errorf("Expected override to be overridden, got %s", override)
	}
}
