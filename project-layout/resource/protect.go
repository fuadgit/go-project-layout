package resource

import (
	"fmt"
	"net/http"

	log "ibfd.org/app/log4u"
	admt "ibfd.org/app/table/adm"
	"ibfd.org/app/uc/adm"
)

// Protector defines an action protector.
type Protector struct {
	authenticator  adm.Authenticator
	basicAuthRealm string
	actions        map[admt.Action]bool
}

// NewProtector creates an action protector.
func NewProtector(authenticator adm.Authenticator, basicAuthRealm string) *Protector {
	return &Protector{authenticator, basicAuthRealm, make(map[admt.Action]bool)}
}

// Protect determines whether a user has access to the requested action.
func (pr *Protector) Protect(action admt.Action, inner http.Handler) http.HandlerFunc {
	pr.validate(action)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer ServerError(w, r)
		user, passw := basicAuth(r)
		context, err := pr.authenticator.Authenticate(user, passw)
		if err != nil {
			sendISError(w, fmt.Sprintf("failed to authenticate: %s", err.Error()))
			return
		}
		if context.IsAuthenticated() {
			hasAccess, err := context.IsAuthorized(action)
			if err != nil {
				sendISError(w, fmt.Sprintf("failed to authorize: %s", err.Error()))
				return
			}
			if hasAccess {
				log.Debugf("granting %s access to %s\n", user, action)
				inner.ServeHTTP(w, r)
				return
			}
			sendError(w, NewError(http.StatusForbidden, "no access"))
			return
		}
		w.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=%q", pr.basicAuthRealm))
		sendError(w, NewError(http.StatusUnauthorized, "not authorized"))
	})
}

func (pr *Protector) validate(action admt.Action) {
	if !pr.actions[action] {
		pr.actions[action] = true
	} else {
		log.Fatalf("duplicate action detected: %s", action)
	}
	if !pr.authenticator.IsDefined(action) {
		log.Fatalf("no access rules defined yet for action: %s", action)
	}
}

func basicAuth(r *http.Request) (admt.Username, admt.Password) {
	un, pw, _ := r.BasicAuth()
	return admt.Username(un), admt.Password(pw)
}
