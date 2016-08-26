/*
Package gatekeeper is a multi-framework Hydra middleware library for golang
This package implements support for the Echo web framework
*/
package gatekeeper

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	hydra "github.com/ory-am/hydra/sdk"
)

// EchoGK represents an instance of Gatekeeper for the Echo web framework
type EchoGK struct {
	hc *hydra.Client
}

// New creates a new Gatekeeper instance for the Echo web framework
func New(hc *hydra.Client) *EchoGK {
	return &EchoGK{
		hc: hc,
	}
}

// ScopesRequired verifies if the token is valid and if the scope requirements are met
func (gk *EchoGK) ScopesRequired(scopes ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, err := gk.hc.Warden.TokenValid(c, gk.hc.Warden.TokenFromRequest(c.Request().(*standard.Request).Request), scopes...)
			if err != nil {
				c.Error(err)
				return echo.ErrUnauthorized
			}
			// All required scopes are found
			c.Set("hydra", ctx)
			return next(c)
		}
	}
}
