package adm

import "strings"

// Username defines a user name
type Username string

// Password defines a password
type Password string

// User defines a ??? user.
type User struct {
	Username Username
	Password Password
}

// Action defines an area of functionality used for authorization purposes.
type Action string

// AdminStore defines the interface for administration storage.
type AdminStore interface {
	ReadUser(Username) (*User, error)
	ReadActions() ([]Action, error)
	ReadPermissions(Username) ([]Action, error)
}

// Table provides implementation of admin store
type Table struct {
	store AdminStore
}

// NewTable creates an Admin table
func NewTable(s AdminStore) *Table {
	return &Table{s}
}

// ReadUser fetches a user by username
func (t *Table) ReadUser(uname Username) (*User, error) {
	return t.store.ReadUser(uname)
}

// ReadActions fetches all actions
func (t *Table) ReadActions() ([]Action, error) {
	return t.store.ReadActions()
}

// ReadPermissions fetches all permissions for a user
func (t *Table) ReadPermissions(uname Username) ([]Action, error) {
	return t.store.ReadPermissions(uname)
}

// In returns true if the value in the action presents in given actions
func (action Action) In(actions []Action) bool {
	for _, a := range actions {
		if a.Equals(action) {
			return true
		}
	}
	return false
}

// Equals determines whether two actions compare equal.
func (action Action) Equals(other Action) bool {
	return strings.EqualFold(string(action), string(other))
}

func (action Action) String() string {
	return string(action)
}

// IsValid checks if the password is valid or not
func (u *User) IsValid(pw Password) bool {
	return len(pw) > 0 && u.Password.Equals(pw)
}

func (un Username) String() string {
	return string(un)
}

// Equals determines whether passwords compare equal.
func (pw Password) Equals(other Password) bool {
	return string(pw) == string(other)
}

// Equals determines whether usernames compare equal.
func (un Username) Equals(other Username) bool {
	return strings.EqualFold(un.String(), other.String())
}

func (pw Password) String() string {
	return "•••••"
}
