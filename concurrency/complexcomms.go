package concurrency

import (
	"fmt"
	"sync"
)

type User struct {
	name string
}

type Book struct {
	name string
}

func (u *User) setUser(val string, wg *sync.WaitGroup) {
	u.name = val
	wg.Done()
}

func (u *User) getUser() string {
	return u.name
}

func (b *Book) setBook(val string, wg *sync.WaitGroup) {
	b.name = val
	wg.Done()
}

func (b *Book) getBook() string {
	return b.name
}

func Monitor() {
	user := &User{}
	book := &Book{}
	var wg sync.WaitGroup
	wg.Add(2)
	go user.setUser("amogh", &wg)
	go book.setBook("harry potter", &wg)

	wg.Wait()
	db := user.getUser()
	db2 := book.getBook()
	fmt.Println("db updated", db, db2)
}
