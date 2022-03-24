package mem

import admt "ibfd.org/app/table/adm"

// FakeAdminDB implements a read-only in memory user/permissions database.
type FakeAdminDB struct {
	upwDB    map[admt.Username]admt.Password
	accessDB map[admt.Action][]admt.Username
}

// NewFakeAdminDB creates a fake administration database.
func NewFakeAdminDB() *FakeAdminDB {
	upwDB := make(map[admt.Username]admt.Password)
	upwDB["app"] = "app" // TODO update username and password for basic-auth

	// TODO update access permissions
	accessDB := make(map[admt.Action][]admt.Username)
	accessDB["SafeHome"] = []admt.Username{"app"}

	return &FakeAdminDB{upwDB, accessDB}
}

// ReadUser reads user data by username.
func (fad *FakeAdminDB) ReadUser(username admt.Username) (*admt.User, error) {
	for key, value := range fad.upwDB {
		if username.Equals(key) {
			return &admt.User{Username: key, Password: value}, nil
		}
	}
	return nil, nil
}

// ReadActions fetches all actions.
func (fad *FakeAdminDB) ReadActions() ([]admt.Action, error) {
	actions := make([]admt.Action, 0, len(fad.accessDB))
	for action := range fad.accessDB {
		actions = append(actions, action)
	}
	return actions, nil
}

// ReadPermissions fetches all authorized actions for a user.
func (fad *FakeAdminDB) ReadPermissions(username admt.Username) ([]admt.Action, error) {
	allowed := make([]admt.Action, 0, len(fad.accessDB))
	for action, usernames := range fad.accessDB {
		for _, un := range usernames {
			if username.Equals(un) {
				allowed = append(allowed, action)
			}
		}
	}
	return allowed, nil
}
