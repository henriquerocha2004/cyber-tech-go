package entities

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	Employment string = "employment"
	Client     string = "client"
)

// UserCommandRepository command contract if database will implement
type UserCommandRepository interface {
	Create(user User) (int, error)
	Update(user User) error
	Delete(userId int) error
	ResetPassWord(userId int, password string) error
	CreateAddress(address Address) error
	UpdateAddress(address Address) error
	DeleteAddress(addressId int) error
	CreateContact(contact Contact) error
	UpdateContact(contact Contact) error
	DeleteContact(contactId int) error
}

// UserQueryRepository query contract if database will implement
type UserQueryRepository interface {
	FindById(userId int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll(typeUser string) ([]User, error)
	FindContactsGroupedByUsers(usersIds []int) (map[int][]Contact, error)
	FindAddressGroupedByUsers(usersIds []int) (map[int][]Address, error)
}

type User struct {
	Id         int       `json:"id,omitempty" db:"id,omitempty"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Type       string    `json:"type,omitempty" db:"type"`
	Email      string    `json:"email,omitempty" db:"email"`
	Document   string    `json:"document,omitempty" db:"document"`
	Password   string    `json:"password,omitempty" db:"password"`
	TypePerson string    `json:"type_person,omitempty" db:"type_person"`
	LastLogin  string    `json:"last_login,omitempty" db:"last_login,omitempty"`
	CreatedAt  string    `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  string    `json:"updated_at,omitempty" db:"updated_at"`
	Address    []Address `json:"address,omitempty"`
	Contacts   []Contact `json:"contacts,omitempty"`
	CreatedBy  int       `json:"created_by,omitempty" db:"created_by"`
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) PasswordAreEqual(password, confirm string) bool {
	return password == confirm
}

func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}

func (u *User) AddAddress(address Address) {
	u.Address = append(u.Address, address)
}

func (u *User) AddContact(contact Contact) {
	u.Contacts = append(u.Contacts, contact)
}

func (u *User) SetType(typePerson string) error {
	switch typePerson {
	case Employment, Client:
		u.TypePerson = typePerson
		return nil
	default:
		return errors.New("invalid type")
	}
}
