package mandatory

import "strconv"

// User is an user
type User struct {
	Name     string
	LastName string
	Age      int
	IsAdult  bool
}

//NewUser creates a new adult user
func NewUser(name, lastName string, age int) *User {
	owca := User{Name: name}
	owca.LastName = lastName
	owca.Age = age
	owca.IsAdult = Is18(age)
	return &owca
}

//ToString makes user a string
func (u *User) ToString() string {
	x := string(u.Name) + string(u.LastName) + strconv.Itoa(u.Age)
	return x
}

//ChangeAge changes the age
func (u *User) ChangeAge(newAge int) {
	if newAge > 120 || newAge < 0 {
		return
	}
	u.Age = newAge
	return
}
