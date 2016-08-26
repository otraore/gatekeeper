package gatekeeper

import (
	"net/http"

	"github.com/goadesign/goa"
	hydra "github.com/ory-am/hydra/sdk"
	"golang.org/x/net/context"
)

var errUnAuthorized = goa.NewErrorClass("Unauthorized", http.StatusUnauthorized)

// GoaGK represents an instance of Gatekeeper for the Echo web framework
type GoaGK struct {
	hc *hydra.Client
}

// NewGoa creates a new Gatekeeper instance for the Echo web framework
func NewGoa(hc *hydra.Client) *GoaGK {
	return &GoaGK{
		hc: hc,
	}
}

// ScopesRequired verifies if the token is valid and if the scope requirements are met
func (gk *GoaGK) ScopesRequired(scopes ...string) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			hctx, err := gk.hc.Warden.TokenValid(ctx, gk.hc.Warden.TokenFromRequest(req), scopes...)
			if err != nil {
				return errUnAuthorized(err)
			}
			// All required scopes are found
			context.WithValue(ctx, "hydra", hctx)
			return h(ctx, rw, req)
		}
	}
}
