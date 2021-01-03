package basic

import "fmt"

type User struct {
	name string
}

type College struct {
	name string
}

type Entityer interface {
	comBehOne()
}

func (u *User) comBehOne() {
	u.name = "amogh"
	fmt.Println("thi is common behavior in user and college", u.name)
}

func (c *College) comBehOne() {
	c.name = "pes"
	fmt.Println("thi is common behavior in college and user", c.name)
}

func (u *User) alagForUser() {
	fmt.Println("sub behavior user")
}
func (c *College) alagForCollege() {
	fmt.Println("sub behavior college")
}

func SubB(e Entityer) {
	e.comBehOne()
	// e.alagForUser()
}

func RunAll() {
	u := &User{}
	SubB(u)
}
