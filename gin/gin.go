package gatekeeper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	hydra "github.com/ory-am/hydra/sdk"
)

// GinGK represents an instance of Gatekeeper for the Gin web framework
type GinGK struct {
	hc *hydra.Client
}

// New creates a new Gatekeeper instance for the Gin web framework
func New(hc *hydra.Client) *GinGK {
	return &GinGK{
		hc: hc,
	}
}

// ScopesRequired verifies if the token is valid and if the scope requirements are met
func (gk *GinGK) ScopesRequired(scopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := gk.hc.Warden.TokenValid(c, gk.hc.Warden.TokenFromRequest(c.Request), scopes...)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// All required scopes are found
		c.Set("hydra", ctx)
		c.Next()
	}
}
