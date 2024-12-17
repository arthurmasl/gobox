package main

import (
	"fmt"
	"time"
	"unsafe"
)

const COUNT = 3

type User struct {
	name    string
	balance int
	store   [1024 * 10]byte
}

var store [1024 * 10]byte

func NewUserP(name string, balance int) *User {
	return &User{name, balance, store}
}

func NewUserV(name string, balance int) User {
	return User{name, balance, store}
}

func main() {
	// u1 := NewUserP("bob", 100)
	// v1 := NewUserV("bob", 100)

	// fmt.Printf("%#v\n%T\n%p\n%v\n", u1, u1, u1, unsafe.Sizeof(u1))
	// fmt.Printf("%#v\n%T\n%v\n", v1, v1, unsafe.Sizeof(v1))

	usersP := make([]*User, 0, COUNT)
	for range COUNT {
		usersP = append(usersP, NewUserP("bob", 100))
	}

	// fmt.Printf("%v\n%T\n%v\n", usersP, usersP, unsafe.Sizeof(usersP))

	usersV := make([]User, 0, COUNT)
	for range COUNT {
		usersV = append(usersV, NewUserV("bob", 100))
	}

	// fmt.Printf("%v\n%T\n%v\n", usersV, usersV, unsafe.Sizeof(usersV))

	t1 := time.Now()
	for _, user := range usersP {
		user.balance += 100
	}
	fmt.Println(time.Since(t1))
	fmt.Println(unsafe.Sizeof(usersP[0].store))

	// t2 := time.Now()
	// for i := range usersV {
	// 	usersV[i].balance += 200
	// }
	// fmt.Println(time.Since(t2))
}
