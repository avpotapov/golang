package user

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	user := New()
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 0 {
		t.Error("length of subUsers must be 0")
	}
	if user.Name != "" {
		t.Error("user name must be empty")
	}
}

func TestNewWithArg(t *testing.T) {
	name := "new user"

	user := New(name)
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 0 {
		t.Error("length of subUsers must be 0")
	}
	if name != user.Name {
		t.Error("user has not name")
	}
}

func TestNewWithArgs(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
}

func TestEqual(t *testing.T) {
	user1, user2 := New("new user", "new subuser1", "new subuser2"), New("new user", "new subuser1", "new subuser2")
	if !user1.Equal(user2) {
		t.Error("structs is equals")
	}
}

func TestNotEqual(t *testing.T) {
	user1, user2 := New("new user", "new subuser", "new subuser2"), New("new user", "new subuser1", "new subuser2")
	if user1.Equal(user2) {
		t.Error("structs is not equals")
	}
}

func TestClone(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}

	cloned := user.Clone()
	if !user.Equal(cloned) {
		t.Error("structs is equals")
	}
	if user == cloned {
		t.Errorf("struct pointer must be defferent: %v- %v", user, cloned)
	}
	if &user.SubUsers == &cloned.SubUsers {
		t.Errorf("slice pointer must be defferent: %v- %v", &user.SubUsers, &cloned.SubUsers)
	}
}

func TestToString(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}

	fmt.Println(user.String())
}

func TestAddZeroSubUser(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
	if user.AddSubUsers() != 0 {
		t.Error("added subuser must be 0")
	}
}

func TestAddOneSubUser(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
	if user.AddSubUsers("new subuser3") != 1 {
		t.Error("added subuser must be 1")
	}
}

func TestAddManySubUser(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
	if user.AddSubUsers("new subuser3", "new subuser4", "new subuser5") != 3 {
		t.Error("added subusers must be 3")
	}
}

func TestRemoveZeroSubUser(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
	if user.RemoveSubUsers() != 0 {
		t.Error("removed subuser must be 0")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
}

func TestRemoveOneSubUser(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
	if user.RemoveSubUsers("new subuser1") != 1 {
		t.Error("removed subuser must be 1")
	}
	if len(user.SubUsers) != 1 {
		t.Error("length of subUsers must be 1")
	}
}

func TestRemoveManySubUser(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	if user == nil {
		t.Error("user not be nil")
	}
	if user.SubUsers == nil {
		t.Error("subUsers not be nil")
	}
	if len(user.SubUsers) != 2 {
		t.Error("length of subUsers must be 2")
	}
	if user.RemoveSubUsers("new subuser1", "new subuser2") != 2 {
		t.Error("removed subuser must be 2")
	}
	if len(user.SubUsers) != 0 {
		t.Error("length of subUsers must be 0")
	}
}

func TestCloneAndModify(t *testing.T) {
	user := New("new user", "new subuser1", "new subuser2")
	cloned := user.Clone()
	cloned.Name = "cloned name"
	cloned.AddSubUsers("new subuser3", "new subuser4", "new subuser5", "new subuser6")
	cloned.RemoveSubUsers("new subuser1", "new subuser2", "new subuser6")
	fmt.Println(user.String())
	fmt.Println(cloned.String())

}
