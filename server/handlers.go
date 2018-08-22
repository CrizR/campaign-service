package server

import (
	"net/http"

	"github.com/ecclesia-dev/account-service/models"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

const (
	indent = "  "
)

func (s *Server) createCampaign(ctx echo.Context) error {

	var account models.Account
	if err := ctx.Bind(&account); err != nil {
		log.Error(err)
	}
	// Add validation call here
	if err := s.accounts.CreateAccount(account); err != nil {
		log.Error(err)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) getCampaignByID(ctx echo.Context) error {
	accounts, err := s.accounts.GetAllAccounts()
	if err != nil {
		log.Error(err)
	}
	return ctx.JSONPretty(http.StatusOK, accounts, indent)
}

func (s *Server) getCampaignByTitle(ctx echo.Context) error {
	account, err := s.accounts.GetAccountByEmail(ctx.Param("email"))
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSONPretty(http.StatusOK, account, indent)
}

func (s *Server) getAccountByID(ctx echo.Context) error {
	account, err := s.accounts.GetAccountByID(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSONPretty(http.StatusOK, &account, indent)
}

func (s *Server) getAllCampaigns(ctx echo.Context) error {
	err := s.accounts.RemoveAccount(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) removeCampaign(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}


func (s *Server) updateCampaign(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) createAction(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) removeAction(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) getAllActions(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) enrollUserInCampaign(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (s *Server) removeUserFromCampaign(ctx echo.Context) error {

	var update map[string]interface{}
	if err := ctx.Bind(&update); err != nil {
		log.Fatal(err)
	}
	if err := s.accounts.UpdateAccount(ctx.Param("id"), nil); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, nil)
}