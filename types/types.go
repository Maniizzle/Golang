package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func TheStruct() {
	theStruct := myStruct{id: 4, name: "bola"}
	fmt.Println(theStruct.name)
	theStruct.name = "bose"
	fmt.Println(theStruct.name)

}

// struct is copied by value
// it is comparable
type myStruct struct {
	id   int
	name string
}

// mmethod
func (u myStruct) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.name, u.id)
}

type MenuItem struct {
	Name   string
	Prices map[string]float64
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Country() string
	Identifiable
}

// conflict
//type Conflict interface {
//	ID() string
//}

type socialSecurityNumber string

func NewSocialCescurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United states of America"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

//testing conflict type
// type test string
// func NewConflict(value test) Conflict {
// 	return value
// }
// func (ts test) ID() string {
// 	return string(ts)
// }

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	case europeanUnionIdentifier:
		return v
	default:
		panic("using an invalid type to initialize")
	}

}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	firstName string
	lastName  string
}

type Employee struct {
	Name
}

type Person struct {
	name             Name
	twitterHandler   TwitterHandler
	instagramHandler InstagramHandler
	Citizen
	// Conflict
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		name: Name{
			firstName: firstName,
			lastName:  lastName,
		},
		Citizen: citizen,
		// Conflict: conflict,
	}
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

// type alias
// this copies the field and method set. they become that exact type
type TwitterHandler = string

// type declaration
// this copies the fields of a type over to a new type
type InstagramHandler string

//	func (p *Person) ID(handler TwitterHandler) string {
//		return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
//	}
func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

// ppointer base receivers for shared state
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) SetInstagramHandler(handler InstagramHandler) error {
	if len(handler) == 0 {
		p.instagramHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter must start with an @ symbol")
	}
	p.instagramHandler = handler
	return nil
}

func (th InstagramHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://instagram.com/%s", cleanHandler)
}

func (p Person) TwitterHandler() string {
	return p.twitterHandler
}

func (p Person) InstagramHandler() InstagramHandler {
	return p.instagramHandler
}
func (p Person) FullName() string {
	return p.name.firstName + " " + p.name.lastName
}

type Handler struct {
	handle string
	name   string
}

func (h Handler) RandomFunc() {

}

type FacebookHandler Handler
type FacebookHandlerr = Handler

func (fh FacebookHandler) Redirect() string {

	//I can access the randomfunc method of type Handler
	dd := Handler{handle: "koya", name: "bose"}
	dd.RandomFunc()

	//I cannot access the randomfunc method of type Handler even though
	//FacebookHandler is of type Handler
	jj := FacebookHandler{name: "geez", handle: "unknown"}
	jj.Redirect()
	return ""
}

func (fh FacebookHandlerr) RedirectAnother() string {

	//I can access the randomfunc method of type Handler
	dd := Handler{handle: "koya", name: "bose"}
	dd.RandomFunc()

	//I can access the randomfunc method of type Handler
	//because FacebookHandlerr is a type alias so it becomes a Handler
	jj := FacebookHandlerr{name: "geez", handle: "unknown"}
	jj.RandomFunc()
	return ""
}

func TypesMain() {
	//d := types.Person{}
	// p := types.NewPerson("Bola", "Tito", types.NewEuropeanUnionIdentifier("123-4554-244", "Germany"), types.NewConflict("bose"))
	p := NewPerson("Bola", "Tito", NewEuropeanUnionIdentifier("123-4554-244", "Germany"))

	p.SetTwitterHandler("@olamide")
	p.SetInstagramHandler("@olamideinstagram")

	println(p.TwitterHandler())
	println(p.InstagramHandler().RedirectUrl())
	println(p.ID())
	println(p.Country())
	println(p.FullName())

	//comparing interfaces
	eu := NewEuropeanUnionIdentifier("123", "Japan")
	eu2 := NewEuropeanUnionIdentifier(123, "Japan")
	println(eu2.ID())
	println(eu2.Country())

	//interface  without complex structure are comparable
	if eu == eu2 {
		println("We match")
	}
	return

}
