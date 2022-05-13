package bdd

import (
	"errors"
	"github.com/cucumber/godog"
	"go-postgres/postgres"
)

func iHaveAccount(login string) error {
	return nil
}

func iMakeLogInWithPassword(password string) error {
	return nil
}

func iHavePermissions(permissions string, login string, password string) error {
	user, err := postgres.GetUserWithNameAndPassword(login, password)
	if err != nil {
		return err
	}
	group, err := postgres.GetGroupWithID(user.Role)
	if err != nil {
		return err
	}
	if group.Permissions != permissions {
		return errors.New("invalid account permissions")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have an administrator account with login "([^"]*)"$`, iHaveAccount)
	ctx.Step(`^I make log in with password "([^"]*)"$`, iMakeLogInWithPassword)
	ctx.Step(`^I have "([^"]*)" with credentials "([^"]*)":"([^"]*)"$`, iHavePermissions)
	ctx.Step(`^I have an reader account with login "([^"]*)"$`, iHaveAccount)
	ctx.Step(`^I make log in with password "([^"]*)"$`, iMakeLogInWithPassword)
	ctx.Step(`^I have "([^"]*)" with credentials "([^"]*)":"([^"]*)"$`, iHavePermissions)
	ctx.Step(`^I have an redactor account with login "([^"]*)"$`, iHaveAccount)
	ctx.Step(`^I make log in with password "([^"]*)"$`, iMakeLogInWithPassword)
	ctx.Step(`^I have "([^"]*)" with credentials "([^"]*)":"([^"]*)"$`, iHavePermissions)
}
