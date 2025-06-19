package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
	ban  bool
	role string
}

type sysAdmin struct {
	*User
	accessLevel int
}

type Describer interface {
	Description() string
}

func main() {

	//функции

	Adminable := registerUser("Balbes", 19)
	//Adminable.Ban()
	//Adminable.UnBan()
	Adminable.promoteToAdmin()
	fmt.Println(Adminable)

	fmt.Println(Adminable.userString())

	AccessPromoter := sysAdmin{
		User:        nil,
		accessLevel: 0,
	}
	AccessPromoter.increaseAccessLvl(10)
	fmt.Println(AccessPromoter)

	//fmt.Println(f.fioAdmin(c))

	user := User{age: 15}
	//sys := sysAdmin{}

	fmt.Println(user)
	String(user)

}

func registerUser(inName string, age int) *User {
	if age <= 13 {
		fmt.Println("ИДИ НАХУЙ, КНИЖКИ ЧИТАЙ, ШКОЛОТА")
		return nil
	}
	return &User{
		name: inName,
		age:  age,
		ban:  false,
		role: "user",
	}
}

//func (u User) Ban() {
//	u.ban = true
//}
//
//func (u User) UnBan() {
//	u.ban = false
//}

func (us *User) promoteToAdmin() {
	if us.age <= 18 {
		fmt.Println("возраст недостаточный")
		return
	}
	us.role = "admin"
}

func (us *User) userString() string {
	usString := fmt.Sprintf("Имя: %s , лет: %d , Бан: %v , Роль: %s", us.name, us.age, us.ban, us.role)
	return usString
}

func (a *sysAdmin) increaseAccessLvl(b int) {
	a.accessLevel += b

}

//func (a *sysAdmin) fioAdmin(c string) string {
//	c = a.userString()
//	c += a.increaseAccessLvl()
//	return c
//}

//func (a sysAdmin) makeObject() string {
//	if a.age <= 18 {
//		return "сисадминит"
//	}
//	return "юзерит"
//}

func (us User) Description() string {
	if us.age > 18 {
		return "сисадминит"
	}
	return "юзерит"
}

func String(o Describer) {
	fmt.Print(o.Description())
}
