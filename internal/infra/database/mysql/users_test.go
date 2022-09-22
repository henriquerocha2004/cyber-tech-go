// +go:building integration
package mysql_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/database/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type TestUserSuit struct {
	suite.Suite
	connection            *sqlx.DB
	userCommandRepository mysql.UserCommandRepository
	userQueryRepository   mysql.UserQueryRepository
	transaction           *sqlx.Tx
}

func setupDatabase() (*mysql.MysqlConfig, error) {
	viper.SetConfigName("env_test")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../../../")
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	database, err := json.Marshal(viper.Get("database"))
	if err != nil {
		return nil, err
	}

	var mysqlConfig mysql.MysqlConfig
	err = json.Unmarshal(database, &mysqlConfig)
	if err != nil {
		return nil, err
	}
	return &mysqlConfig, nil
}

func connectDatabase(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newTestUserSuit() *TestUserSuit {
	return &TestUserSuit{}
}

func TestUserTests(t *testing.T) {
	suite.Run(t, newTestUserSuit())
}

func (s *TestUserSuit) SetupTest() {
	config, err := setupDatabase()
	if err != nil {
		log.Fatal("Error in get database information: " + err.Error())
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	connection, err := connectDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}
	s.connection = connection
	s.userCommandRepository = *mysql.NewUserCommandRepository(s.connection)
	s.userQueryRepository = *mysql.NewUserQueryRepository(s.connection)
}

func (s *TestUserSuit) BeforeTest(suiteName, testName string) {
	s.transaction = s.connection.MustBegin()
}

func (s *TestUserSuit) AfterTest(suiteName, testName string) {
	s.transaction.Rollback()
}

func (s *TestUserSuit) TestUser() {
	s.Run("should create user", func() {
		user := &entities.User{
			FirstName: "Henrique",
			LastName:  "Rocha",
			Email:     "teste@teste.com",
			Type:      "employment",
			CreatedBy: 1,
		}
		user.SetPassword("12345678")

		id, err := s.userCommandRepository.Create(*user)
		s.NoError(err)
		s.GreaterOrEqual(id, 1)
	})
	s.Run("should update user", func() {
		user := &entities.User{
			FirstName: "Henrique",
			LastName:  "Rocha",
			Email:     "teste@teste.com",
			Type:      "employment",
			CreatedBy: 1,
		}
		user.SetPassword("12345678")

		id, err := s.userCommandRepository.Create(*user)
		s.NoError(err)

		user.FirstName = "Jose"
		user.LastName = "Fernando"
		user.Id = id

		err = s.userCommandRepository.Update(*user)
		s.NoError(err)
	})

	s.Run("should delete user", func() {
		user := &entities.User{
			FirstName: "Henrique",
			LastName:  "Souza",
			Email:     "teste@teste.com",
			Type:      "employment",
			CreatedBy: 1,
		}
		user.SetPassword("12345678")

		id, err := s.userCommandRepository.Create(*user)
		s.NoError(err)
		err = s.userCommandRepository.Delete(id)
		s.NoError(err)
	})
	s.Run("should reset password", func() {
		user := &entities.User{
			FirstName: "Henrique",
			LastName:  "Souza",
			Email:     "teste@teste.com",
			Type:      "employment",
			CreatedBy: 1,
		}
		user.SetPassword("12345678")
		id, err := s.userCommandRepository.Create(*user)
		s.NoError(err)
		user.SetPassword("43217899")
		err = s.userCommandRepository.ResetPassWord(id, user.Password)
		s.NoError(err)
	})
	s.Run("should create address", func() {
		user := &entities.User{
			FirstName: "Henrique",
			LastName:  "Souza",
			Email:     "teste@teste.com",
			Type:      "employment",
			CreatedBy: 1,
		}
		user.SetPassword("12345678")
		id, err := s.userCommandRepository.Create(*user)
		s.NoError(err)

		address := entities.Address{
			Street:   "Rua dos Bobos Nº 40",
			City:     "Namekuzei",
			District: "Algum Bairro",
			State:    "BA",
			Country:  "Brazil",
			ZipCode:  "12345678",
			Type:     "personal",
			UserId:   id,
		}

		err = s.userCommandRepository.CreateAddress(address)
		s.NoError(err)
	})

	s.Run("should return user by id", func() {
		user := &entities.User{
			FirstName: "Henrique",
			LastName:  "Souza",
			Email:     "teste@teste.com",
			Type:      "client",
			CreatedBy: 1,
		}
		user.SetPassword("12345678")
		id, err := s.userCommandRepository.Create(*user)
		s.NoError(err)

		address := entities.Address{
			Street:   "Rua dos Bobos Nº 40",
			City:     "Namekuzei",
			District: "Algum Bairro",
			State:    "BA",
			Country:  "Brazil",
			ZipCode:  "12345678",
			Type:     "personal",
			UserId:   id,
		}

		err = s.userCommandRepository.CreateAddress(address)
		s.NoError(err)

		contact := entities.Contact{
			Type:       "personal",
			Phone:      "74145669996",
			IsWhatsApp: false,
			UserId:     id,
		}

		err = s.userCommandRepository.CreateContact(contact)
		s.NoError(err)

		userDb, err := s.userQueryRepository.FindById(id)
		s.NoError(err)
		s.Equal(userDb.FirstName, user.FirstName)
		s.Equal(userDb.LastName, user.LastName)
	})
	s.Run("should return all users", func() {
		users, err := s.userQueryRepository.FindAll("client")
		s.NotEmpty(users)
		s.NoError(err)
	})
}
