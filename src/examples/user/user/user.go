// user project user.go
package user

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type SubUsers []*User

type User struct {
	Name     string
	SubUsers SubUsers
}

func New(names ...string) *User {
	user := User{Name: "", SubUsers: make(SubUsers, 0, 10)}
	for i, name := range names {
		if i == 0 {
			user.Name = name
		} else {
			user.SubUsers = append(user.SubUsers, New(name))
		}
	}
	return &user
}

func (self *User) Equal(other *User) bool {
	return reflect.DeepEqual(self, other)
}

func (self *User) Clone() *User {
	cloned := New()
	*cloned = *self
	cloned.SubUsers = make(SubUsers, 0, 10)
	for _, subUser := range self.SubUsers {
		cloned.SubUsers = append(cloned.SubUsers, subUser.Clone())
	}
	return cloned
}

func (self *User) String() string {
	out, err := json.MarshalIndent(self, "", "\t")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return string(out)
}

func (self *User) AddSubUsers(names ...string) int {
	var count int = 0
	for _, name := range names {
		self.SubUsers = append(self.SubUsers, New(name))
		count++
	}
	return count
}

func (self *User) RemoveSubUsers(names ...string) int {
	var count int = 0
	var subUsers map[string]int
	isCreateIndex := true

	for _, name := range names {
		if isCreateIndex {
			subUsers = make(map[string]int, len(self.SubUsers))
			for i, subUser := range self.SubUsers {
				subUsers[subUser.Name] = i
			}
			isCreateIndex = false
		}
		if idx, ok := subUsers[name]; ok {
			self.SubUsers = append(self.SubUsers[:idx], self.SubUsers[idx+1:]...)
			count++
			isCreateIndex = true
		}
	}
	return count
}
