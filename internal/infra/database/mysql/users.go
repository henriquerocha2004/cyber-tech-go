package mysql

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UserCommandRepository struct {
	mysqlConnection *sqlx.DB
}

func NewUserCommandRepository(mysqlConn *sqlx.DB) *UserCommandRepository {
	return &UserCommandRepository{
		mysqlConnection: mysqlConn,
	}
}

func (u *UserCommandRepository) Create(user entities.User) (int, error) {
	user.CreatedAt = time.Now().Format("2006-01-02 15:04")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04")

	query := `
		INSERT INTO users 
		(first_name, last_name, type, email, document, password, type_person, created_at, updated_at, created_by)
		VALUES 
		(:first_name, :last_name, :type, :email, :document, :password, :type_person, :created_at, :updated_at, :created_by)
	`
	result, err := u.mysqlConnection.NamedExec(query, user)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (u *UserCommandRepository) Update(user entities.User) error {
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04")
	query := `
		UPDATE users SET first_name = :first_name, last_name = :last_name, type = :type, email = :email,
			type_person = :type_person, updated_at = :updated_at WHERE id = :id
	`
	_, err := u.mysqlConnection.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserCommandRepository) Delete(userId int) error {
	query := `DELETE FROM users WHERE id = ?`
	u.mysqlConnection.MustExec(query, userId)
	return nil
}

func (u *UserCommandRepository) ResetPassWord(userId int, password string) error {
	updatedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE users SET password = ?, updated_at = ? WHERE id = ?`
	u.mysqlConnection.MustExec(query, password, updatedAt, userId)
	return nil
}

func (u *UserCommandRepository) CreateAddress(address entities.Address) error {
	query := `
		INSERT INTO addresses 
			(street, city, district, state, country, zip_code, type, user_id)
		VALUES
			(:street, :city, :district, :state, :country, :zip_code, :type, :user_id)
	`
	_, err := u.mysqlConnection.NamedExec(query, address)
	return err
}

func (u *UserCommandRepository) UpdateAddress(address entities.Address) error {
	log.Println(address)
	query := `
		UPDATE addresses SET street = :street, city = :city, district = :district, state = :state, 
			country = :country, zip_code = :zip_code, type = :type, user_id = :user_id
		WHERE id = :id
	`
	_, err := u.mysqlConnection.NamedExec(query, address)
	return err
}

func (u *UserCommandRepository) DeleteAddress(addressId int) error {
	query := `DELETE FROM addresses WHERE id = ?`
	u.mysqlConnection.MustExec(query, addressId)
	return nil
}

func (u *UserCommandRepository) CreateContact(contact entities.Contact) error {
	query := `
		INSERT INTO contacts 
			(type, phone, is_whatsapp, user_id)
		VALUES
			(:type, :phone, :is_whatsapp, :user_id)
	`
	_, err := u.mysqlConnection.NamedExec(query, contact)
	return err
}

func (u *UserCommandRepository) UpdateContact(contact entities.Contact) error {
	query := `
		UPDATE contacts SET type = :type, phone = :phone, is_whatsapp = :is_whatsapp, user_id = :user_id
		WHERE id = :id
	`
	_, err := u.mysqlConnection.NamedExec(query, contact)
	return err
}

func (u *UserCommandRepository) DeleteContact(contactId int) error {
	query := `DELETE FROM contacts WHERE id = ?`
	u.mysqlConnection.MustExec(query, contactId)
	return nil
}

type UserQueryRepository struct {
	mysqlConnection *sqlx.DB
}

func NewUserQueryRepository(mysqlConn *sqlx.DB) *UserQueryRepository {
	return &UserQueryRepository{
		mysqlConnection: mysqlConn,
	}
}

func (u *UserQueryRepository) FindById(userId int) (*entities.User, error) {
	user := entities.User{}
	query := `
		SELECT id, first_name, last_name, type, email, document, type_person, created_at, updated_at, created_by
			FROM users WHERE id = ?
		`
	err := u.mysqlConnection.Get(&user, query, userId)
	if err != nil {
		return nil, err
	}
	if user.Type == entities.Employment {
		return &user, nil
	}

	var usersIds []int
	usersIds = append(usersIds, userId)

	contacts, err := u.FindContactsGroupedByUsers(usersIds)
	if err != nil {
		return nil, err
	}

	if len(contacts) >= 1 {
		user.Contacts = contacts[userId]
	}

	addresses, err := u.FindAddressGroupedByUsers(usersIds)
	if err != nil {
		return nil, err
	}

	if len(addresses) >= 1 {
		user.Address = addresses[userId]
	}

	return &user, nil
}

func (u *UserQueryRepository) FindByEmail(email string) (*entities.User, error) {
	user := entities.User{}
	query := `
		SELECT id, first_name, last_name, type, email, password, created_by
			FROM users WHERE email = ?
		`
	err := u.mysqlConnection.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserQueryRepository) FindContactsGroupedByUsers(usersIds []int) (map[int][]entities.Contact, error) {
	s, _ := json.Marshal(usersIds)
	sIds := strings.Trim(string(s), "[]")
	contacts := []entities.Contact{}
	query := fmt.Sprintf("SELECT * FROM contacts WHERE user_id IN (%s)", sIds)
	err := u.mysqlConnection.Select(&contacts, query)
	if err != nil {
		return nil, err
	}

	groupedContacts := make(map[int][]entities.Contact)

	for _, contact := range contacts {
		groupedContacts[contact.UserId] = append(groupedContacts[contact.UserId], contact)
	}
	return groupedContacts, nil
}

func (u *UserQueryRepository) FindAddressGroupedByUsers(usersIds []int) (map[int][]entities.Address, error) {
	s, _ := json.Marshal(usersIds)
	sIds := strings.Trim(string(s), "[]")
	addresses := []entities.Address{}
	query := fmt.Sprintf("SELECT * FROM addresses WHERE user_id IN (%s)", sIds)
	err := u.mysqlConnection.Select(&addresses, query)
	if err != nil {
		return nil, err
	}

	groupedAddresses := make(map[int][]entities.Address)

	for _, address := range addresses {
		groupedAddresses[address.UserId] = append(groupedAddresses[address.UserId], address)
	}
	return groupedAddresses, nil
}

func (u *UserQueryRepository) FindAll(typeUser string) ([]entities.User, error) {
	var users []entities.User
	query := `SELECT id, first_name, last_name, type, email, document, type_person, created_at, updated_at, created_by FROM users WHERE type = ?`
	err := u.mysqlConnection.Select(&users, query, typeUser)
	if err != nil {
		return nil, err
	}

	if len(users) < 1 {
		return []entities.User{}, nil
	}

	var usersIds []int
	for _, user := range users {
		if user.Type == entities.Client {
			usersIds = append(usersIds, user.Id)
		}
	}

	if len(usersIds) < 1 {
		return users, nil
	}

	contacts, err := u.FindContactsGroupedByUsers(usersIds)
	if err != nil {
		return nil, err
	}

	addresses, err := u.FindAddressGroupedByUsers(usersIds)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.Contacts = contacts[user.Id]
		user.Address = addresses[user.Id]
	}

	return users, nil
}
