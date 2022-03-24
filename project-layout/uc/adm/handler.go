package adm

import (
	log "ibfd.org/app/log4u"
	admt "ibfd.org/app/table/adm"
)

// Context provides the authentication context.
type Context interface {
	IsAuthenticated() bool
	IsAuthorized(action admt.Action) (bool, error)
}

// AdminContext holds the current authentication/authorization context.
type adminContext struct {
	actions       []admt.Action
	user          *admt.User
	authenticated bool
}

// Authenticator provides qClass authentication and authorization.
type Authenticator interface {
	IsDefined(action admt.Action) bool
	Authenticate(username admt.Username, password admt.Password) (Context, error)
}

// AdminHandler implements administrative use-cases.
type AdminHandler struct {
	table   *admt.Table
	actions []admt.Action
}

// New creates an admin handler.
func New(t *admt.Table) *AdminHandler {
	actions, err := t.ReadActions()
	if err != nil {
		log.Fatalf("error fetching actions: [%v]", err)
	}
	return &AdminHandler{t, actions}
}

// IsDefined determines whether an action is defined.
func (ah *AdminHandler) IsDefined(action admt.Action) bool {
	return action.In(ah.actions)
}

// Authenticate a user.
func (ah *AdminHandler) Authenticate(username admt.Username, password admt.Password) (Context, error) {
	user, err := ah.table.ReadUser(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &adminContext{authenticated: false}, nil
	}
	actions, err := ah.table.ReadPermissions(username)
	if err != nil {
		return nil, err
	}
	return &adminContext{actions, user, user.IsValid(password)}, nil
}

// IsAuthenticated determines whether a user is authenticated.
func (ac *adminContext) IsAuthenticated() bool {
	return ac.authenticated
}

// IsAuthorized determines whether this context is allowed access to a specific action.
func (ac *adminContext) IsAuthorized(action admt.Action) (bool, error) {
	return ac.authenticated && action.In(ac.actions), nil
}
