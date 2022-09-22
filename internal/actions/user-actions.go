package actions

import (
	"errors"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type UserInput struct {
	Id              int            `json:"id,omitempty"`
	FirstName       string         `json:"first_name" validate:"required"`
	LastName        string         `json:"last_name" validate:"required"`
	Type            string         `json:"type" validate:"required"`
	Email           string         `json:"email" validate:"required,email"`
	Document        string         `json:"document"`
	Password        string         `json:"password"`
	PasswordConfirm string         `json:"password_confirm"`
	TypePerson      string         `json:"type_person"`
	LastLogin       string         `json:"last_login"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
	Addresses       []AddressInput `json:"addresses,omitempty"`
	Contacts        []ContactInput `json:"contacts,omitempty"`
	CreatedBy       int            `json:"created_by"`
}

type AddressInput struct {
	Id       int    `json:"id,omitempty"`
	Street   string `json:"street" validate:"required"`
	City     string `json:"city" validate:"required"`
	District string `json:"district" validate:"required"`
	State    string `json:"state" validate:"required"`
	Country  string `json:"country" validate:"required"`
	ZipCode  string `json:"zip_code" validate:"required"`
	Type     string `json:"type" validate:"required"`
	UserId   int    `json:"user_id" validate:"required"`
}

type ContactInput struct {
	Id         int    `json:"id,omitempty"`
	Type       string `json:"type" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	IsWhatsApp bool   `json:"is_whatsapp"`
	UserId     int    `json:"user_id" validate:"required"`
}

type UserOutput struct {
	Error   bool   `json:"error"`
	Message string `json:"Message"`
	Data    any    `json:"data"`
}

type UserAction struct {
	userCommandRepository entities.UserCommandRepository
	userQueryRepository   entities.UserQueryRepository
}

func NewUserAction(usrComRepo entities.UserCommandRepository, usrQryRepo entities.UserQueryRepository) *UserAction {
	return &UserAction{
		userCommandRepository: usrComRepo,
		userQueryRepository:   usrQryRepo,
	}
}

func (u *UserAction) Create(userInput UserInput) UserOutput {
	output := UserOutput{}
	user := &entities.User{
		FirstName:  userInput.FirstName,
		LastName:   userInput.LastName,
		Email:      userInput.Email,
		Type:       userInput.Type,
		Document:   userInput.Document,
		TypePerson: userInput.TypePerson,
	}

	if user.Type == entities.Employment {
		if !user.PasswordAreEqual(userInput.Password, userInput.PasswordConfirm) {
			output.Error = true
			output.Message = "Error in create user: " + errors.New("passwords not equals").Error()
			return output
		}
		user.SetPassword(userInput.Password)
	}

	userId, err := u.userCommandRepository.Create(*user)
	if err != nil {
		output.Error = true
		output.Message = "Error in create user: " + err.Error()
		return output
	}

	if len(userInput.Addresses) >= 1 {
		for _, address := range userInput.Addresses {
			address.UserId = userId
			output := u.CreateAddress(address)
			if output.Error {
				return output
			}
		}
	}

	if len(userInput.Contacts) >= 1 {
		for _, contact := range userInput.Contacts {
			contact.UserId = userId
			output := u.CreateContact(contact)
			if output.Error {
				return output
			}
		}
	}

	output.Error = false
	output.Message = "user created successfully"
	return output
}

func (u *UserAction) ChangePassword(userId int, newPassword string) UserOutput {
	output := UserOutput{}
	user, err := u.userQueryRepository.FindById(userId)
	if err != nil {
		output.Error = true
		output.Message = "Error in change password: " + err.Error()
		return output
	}

	err = user.SetPassword(newPassword)
	if err != nil {
		output.Error = true
		output.Message = "Error in change password: " + err.Error()
		return output
	}

	err = u.userCommandRepository.Update(*user)
	if err != nil {
		output.Error = true
		output.Message = "Error in change password: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "Password changed successfully"
	return output
}

func (u *UserAction) Update(userInput UserInput) UserOutput {
	output := UserOutput{}
	user := entities.User{
		Id:         userInput.Id,
		FirstName:  userInput.FirstName,
		LastName:   userInput.LastName,
		Email:      userInput.Email,
		Type:       userInput.Type,
		Document:   userInput.Document,
		TypePerson: userInput.TypePerson,
	}

	err := u.userCommandRepository.Update(user)
	if err != nil {
		output.Error = true
		output.Message = "Error in update user: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "User updated successfully"
	return output
}

func (u *UserAction) Delete(userId int) UserOutput {
	output := UserOutput{}
	err := u.userCommandRepository.Delete(userId)

	if err != nil {
		output.Error = true
		output.Message = "Error in delete user: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "User deleted successfully"
	return output
}

func (u *UserAction) CreateAddress(addressInput AddressInput) UserOutput {
	output := UserOutput{}
	address := &entities.Address{
		Street:   addressInput.Street,
		City:     addressInput.City,
		District: addressInput.District,
		State:    addressInput.District,
		Country:  addressInput.Country,
		ZipCode:  addressInput.ZipCode,
		Type:     addressInput.Type,
		UserId:   addressInput.UserId,
	}

	err := u.userCommandRepository.CreateAddress(*address)
	if err != nil {
		output.Error = true
		output.Message = "Failed to create address: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "address created successfully"
	return output
}

func (u *UserAction) UpdateAddress(addressInput AddressInput) UserOutput {
	output := UserOutput{}
	address := &entities.Address{
		Id:       addressInput.Id,
		Street:   addressInput.Street,
		City:     addressInput.City,
		District: addressInput.District,
		State:    addressInput.State,
		Country:  addressInput.Country,
		ZipCode:  addressInput.ZipCode,
		Type:     addressInput.Type,
		UserId:   addressInput.UserId,
	}

	err := u.userCommandRepository.UpdateAddress(*address)
	if err != nil {
		output.Error = true
		output.Message = "Failed to update address: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "address updated successfully"
	return output
}

func (u *UserAction) DeleteAddress(addressId int) UserOutput {
	output := UserOutput{}
	err := u.userCommandRepository.DeleteAddress(addressId)
	if err != nil {
		output.Error = true
		output.Message = "Failed to delete address: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "address updated successfully"
	return output
}

func (u *UserAction) CreateContact(contactInput ContactInput) UserOutput {
	output := UserOutput{}
	contact := &entities.Contact{
		Type:       contactInput.Type,
		Phone:      contactInput.Phone,
		IsWhatsApp: contactInput.IsWhatsApp,
		UserId:     contactInput.UserId,
	}

	err := u.userCommandRepository.CreateContact(*contact)

	if err != nil {
		output.Error = true
		output.Message = "Failed to delete contact: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "contact created successfully"
	return output
}

func (u *UserAction) UpdateContact(contactInput ContactInput) UserOutput {
	output := UserOutput{}
	contact := &entities.Contact{
		Id:         contactInput.Id,
		Type:       contactInput.Type,
		Phone:      contactInput.Phone,
		IsWhatsApp: contactInput.IsWhatsApp,
		UserId:     contactInput.UserId,
	}

	err := u.userCommandRepository.UpdateContact(*contact)

	if err != nil {
		output.Error = true
		output.Message = "Failed to updated contact: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "contact updated successfully"
	return output
}

func (u *UserAction) DeleteContact(contactId int) UserOutput {
	output := UserOutput{}
	err := u.userCommandRepository.DeleteContact(contactId)

	if err != nil {
		output.Error = true
		output.Message = "Failed to deleted contact: " + err.Error()
		return output
	}

	output.Error = false
	output.Message = "contact deleted successfully"
	return output

}

func (u *UserAction) FindUsers(typeUser string) UserOutput {
	output := UserOutput{}
	users, err := u.userQueryRepository.FindAll(typeUser)
	if err != nil {
		output.Error = true
		output.Message = "Failed to get users: " + err.Error()
		return output
	}

	output.Error = false
	output.Data = users
	return output
}

func (u *UserAction) FindById(userId int) UserOutput {
	output := UserOutput{}
	user, err := u.userQueryRepository.FindById(userId)
	if err != nil {
		output.Error = true
		output.Message = "Failed to get user: " + err.Error()
		return output
	}

	output.Error = false
	output.Data = user
	return output
}

func (u *UserAction) FindAll(typeUser string) UserOutput {
	output := UserOutput{}
	users, err := u.userQueryRepository.FindAll(typeUser)
	if err != nil {
		output.Error = true
		output.Message = "Failed to get users: " + err.Error()
		return output
	}

	output.Error = false
	output.Data = users
	return output
}
