package memory

import "github.com/henriquerocha2004/cyber-tech-go/internal/entities"

var (
	users     []entities.User
	addresses []entities.Address
	contacts  []entities.Contact
)

type MemoryUserCommandRepository struct{}

func (m *MemoryUserCommandRepository) Create(user entities.User) (int, error) {
	user.Id = len(users) + 1
	users = append(users, user)
	return user.Id, nil
}

func (m *MemoryUserCommandRepository) Update(user entities.User) error {
	for index, userSaved := range users {
		if userSaved.Id == user.Id {
			users[index] = user
		}
	}
	return nil
}

func (m *MemoryUserCommandRepository) Delete(userId int) error {

	for index, userSaved := range users {
		if userSaved.Id == userId {
			users = append(users[:index], users[index+1:]...)
		}
	}
	return nil
}

func (m *MemoryUserCommandRepository) ResetPassWord(userId int, password string) error {
	for index, userSaved := range users {
		if userSaved.Id == userId {
			users[index].SetPassword(password)
		}
	}
	return nil
}

func (m *MemoryUserCommandRepository) CreateAddress(address entities.Address) error {
	address.Id = len(addresses) + 1
	addresses = append(addresses, address)
	return nil
}

func (m *MemoryUserCommandRepository) UpdateAddress(address entities.Address) error {
	for index, addressSaved := range addresses {
		if addressSaved.Id == address.Id {
			addresses[index] = address
		}
	}
	return nil
}

func (m *MemoryUserCommandRepository) DeleteAddress(addressId int) error {
	for index, addressSaved := range addresses {
		if addressSaved.Id == addressId {
			addresses = append(addresses[:index], addresses[index+1:]...)
		}
	}
	return nil
}

func (m *MemoryUserCommandRepository) CreateContact(contact entities.Contact) error {
	contact.Id = len(contacts) + 1
	contacts = append(contacts, contact)
	return nil
}

func (m *MemoryUserCommandRepository) UpdateContact(contact entities.Contact) error {
	for index, contactSaved := range contacts {
		if contactSaved.Id == contact.Id {
			contacts[index] = contact
		}
	}
	return nil
}

func (m *MemoryUserCommandRepository) DeleteContact(contactId int) error {
	for index, contactSaved := range contacts {
		if contactSaved.Id == contactId {
			contacts = append(contacts[:index], contacts[index+1:]...)
		}
	}
	return nil
}

type MemoryUserQueryRepository struct{}

func (m *MemoryUserQueryRepository) FindById(userId int) (*entities.User, error) {
	for _, userSaved := range users {
		if userSaved.Id == userId {

			userContacts, err := m.FindContactsByUser(userSaved.Id)
			if err != nil {
				return nil, err
			}
			userSaved.Contacts = append(userSaved.Contacts, userContacts...)

			userAddresses, err := m.FindAddressByUser(userSaved.Id)
			if err != nil {
				return nil, err
			}
			userSaved.Address = append(userSaved.Address, userAddresses...)

			return &userSaved, nil
		}
	}
	return nil, nil
}

func (m *MemoryUserQueryRepository) FindAll(typeUser string) ([]entities.User, error) {
	var u []entities.User

	for _, user := range users {
		userContacts, err := m.FindContactsByUser(user.Id)
		if err != nil {
			return nil, err
		}
		user.Contacts = append(user.Contacts, userContacts...)

		userAddresses, err := m.FindAddressByUser(user.Id)
		if err != nil {
			return nil, err
		}
		user.Address = append(user.Address, userAddresses...)

		u = append(u, user)
	}
	return u, nil
}

func (m *MemoryUserQueryRepository) FindContactById(contactId int) (*entities.Contact, error) {
	for _, contact := range contacts {
		if contact.Id == contactId {
			return &contact, nil
		}
	}
	return nil, nil
}

func (m *MemoryUserQueryRepository) FindContactsByUser(userId int) ([]entities.Contact, error) {
	var userContacts []entities.Contact

	for _, contact := range contacts {
		if contact.UserId == userId {
			userContacts = append(userContacts, contact)
		}
	}
	return userContacts, nil
}

func (m *MemoryUserQueryRepository) FindAddressById(addressId int) (*entities.Address, error) {
	for _, address := range addresses {
		if address.Id == addressId {
			return &address, nil
		}
	}
	return nil, nil
}

func (m *MemoryUserQueryRepository) FindAddressByUser(userId int) ([]entities.Address, error) {
	var userAddress []entities.Address

	for _, address := range addresses {
		if address.UserId == userId {
			userAddress = append(userAddress, address)
		}
	}
	return userAddress, nil
}
