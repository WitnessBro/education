package test

import (
	"os"
	"testing"

	"github.com/WitnessBro/education/internal/config"
	testlib "github.com/WitnessBro/education/test/test_lib"
)

var TestApp *testlib.TestApp

func TestMain(m *testing.M) {
	config, _ := config.LoadConfig("configs/config.yaml")
	TestApp = testlib.TestInitApp(config)
	testValue := m.Run()
	os.Exit(testValue)
}
