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

// New creates a new Gatekeeper instance for the Echo web framework
func New(hc *hydra.Client) *GoaGK {
	return &GoaGK{
		hc: hc,
	}
}

// ScopesRequired verifies if the token is valid and if the scope requirements
// within the goa context are met
func (gk *GoaGK) ScopesRequired() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
			token := gk.hc.Warden.TokenFromRequest(req)
			scopes := goa.ContextRequiredScopes(ctx)
			hydraCtx, err := gk.hc.Warden.TokenValid(ctx, token, scopes...)
			if err != nil {
				return goa.ErrUnauthorized(err)
			}
			context.WithValue(ctx, "hydra", hydraCtx)
			return h(ctx, w, req)
		}
	}
}
