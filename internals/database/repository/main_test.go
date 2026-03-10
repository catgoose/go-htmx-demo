package repository

import (
	"os"
	"testing"

	"catgoose/go-htmx-demo/internals/logger"
)

func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}
