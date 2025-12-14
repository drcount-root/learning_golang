package user

import (
	"errors"
	"math/rand"
	"time"
)

type User struct {
	id        int
	name      string
	age       int
	createdAt time.Time
}

type ExtraInfo struct {
	portfoilio float64
	profit     float64
}

// struct embedding
type Admin struct {
	email    string
	password string
	// embedding User struct
	User
	// embedding ExtraInfo struct
	ExtraInfo
}

func New(name string, age int) (*User, error) {
	if name == "" || age == -1 || age < 0 || age > 150 {
		return nil, errors.New("invalid input data for user")
	}

	return &User{
		id:        rand.Intn(1000),
		name:      name,
		age:       age,
		createdAt: time.Now(),
	}, nil
}

func (u *User) NewExtraInfo(portfoilio, profit float64) (*ExtraInfo, error) {

	if portfoilio < 0 || profit < 0 || profit == -1 || portfoilio == -1 {
		return nil, errors.New("invalid input data for extra info")
	}

	return &ExtraInfo{
		portfoilio,
		profit,
	}, nil
}

func (u *User) UpdateUserDetails(name string, age int) (*User, error) {
	if name == "" || age == -1 || age < 0 || age > 150 {
		return nil, errors.New("invalid input data for user")
	}

	// (*u).name = name
	// (*u).age = age
	u.name = name
	u.age = age
	return u, nil
}

func (u *User) NewAdmin(ext *ExtraInfo, email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid input data for admin")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			id:        u.id,
			name:      u.name,
			age:       u.age,
			createdAt: time.Now(),
		},
		ExtraInfo: ExtraInfo{
			portfoilio: ext.portfoilio,
			profit:     ext.profit,
		},
	}, nil
}

func (au *Admin) DeleteAdminExtraInfo() {
	au.ExtraInfo = ExtraInfo{}
}