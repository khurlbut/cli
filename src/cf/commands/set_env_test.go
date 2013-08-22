package commands_test

import (
	"cf"
	"cf/api"
	. "cf/commands"
	"github.com/stretchr/testify/assert"
	"testhelpers"
	"testing"
)

func TestRunWhenApplicationExists(t *testing.T) {
	app := cf.Application{Name: "my-app", Guid: "my-app-guid"}
	appRepo := &testhelpers.FakeApplicationRepository{AppByName: app}

	args := []string{"my-app", "DATABASE_URL", "mysql://example.com/my-db"}
	ui := callSetEnv(args, appRepo)

	assert.Contains(t, ui.Outputs[0], "my-app")
	assert.Contains(t, ui.Outputs[0], "DATABASE_URL")
	assert.Contains(t, ui.Outputs[1], "OK")

	assert.Equal(t, appRepo.AppName, "my-app")
	assert.Equal(t, appRepo.SetEnvApp, app)
	assert.Equal(t, appRepo.SetEnvName, "DATABASE_URL")
	assert.Equal(t, appRepo.SetEnvValue, "mysql://example.com/my-db")
}

func TestRunWhenAppDoesNotExist(t *testing.T) {
	appRepo := &testhelpers.FakeApplicationRepository{AppByNameErr: true}

	args := []string{"does-not-exist", "DATABASE_URL", "mysql://example.com/my-db"}
	ui := callSetEnv(args, appRepo)

	assert.Contains(t, ui.Outputs[0], "FAILED")
	assert.Contains(t, ui.Outputs[1], "App does not exist.")
}

func TestRunWhenSettingTheEnvFails(t *testing.T) {
	app := cf.Application{Name: "my-app", Guid: "my-app-guid"}
	appRepo := &testhelpers.FakeApplicationRepository{
		AppByName: app,
		SetEnvErr: true,
	}

	args := []string{"does-not-exist", "DATABASE_URL", "mysql://example.com/my-db"}
	ui := callSetEnv(args, appRepo)

	assert.Contains(t, ui.Outputs[0], "Updating env variable")
	assert.Contains(t, ui.Outputs[1], "FAILED")
	assert.Contains(t, ui.Outputs[2], "Failed setting env")
}

func callSetEnv(args []string, appRepo api.ApplicationRepository) (ui *testhelpers.FakeUI) {
	context := testhelpers.NewContext(2, args)
	ui = new(testhelpers.FakeUI)
	se := NewSetEnv(ui, appRepo)
	se.Run(context)

	return
}
