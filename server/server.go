package server

import (
	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log      *logrus.Entry
	echo     *echo.Echo
	accounts controllers.AccountController
}

func New(ctlr controllers.AccountController) Server {
	e := echo.New()
	e.HideBanner = true
	logger := logrus.NewEntry(logrus.New())
	serv := Server{log: logger, echo: e, accounts: ctlr}
	e.Use(serv.LogRequest)
	return serv
}

func (s *Server) Start(port string) {
	s.setRoutes()
	s.echo.Logger.Fatal(s.echo.Start(port))
}

func (s *Server) setRoutes() {
	s.echo.POST("/api/campaigns", s.createCampaign)
	s.echo.PUT("/api/campaigns", s.updateCampaign)
	s.echo.POST("/api/campaigns/action/:campaignID", s.createAction)
	s.echo.DELETE("/api/campaigns/action/", s.removeAction)
	s.echo.GET("/api/campaigns/action", s.getAllActions)
	s.echo.GET("/api/campaigns/:id", s.getCampaignByID)
	s.echo.GET("/api/campaigns/:title", s.getCampaignByTitle)
	s.echo.POST("/api/campaigns/accounts/:userID/:campaignID", s.enrollUserInCampaign)
	s.echo.DELETE("/api/campaigns/accounts/:userID/:campaignID", s.removeUserFromCampaign)
	s.echo.GET("/api/campaigns", s.getAllCampaigns)
}
