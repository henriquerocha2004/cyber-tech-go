package di

import (
	"github.com/henriquerocha2004/cyber-tech-go/cmd/api/handlers"
	"github.com/henriquerocha2004/cyber-tech-go/internal/actions"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/auth"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/database/mysql"
	"github.com/jmoiron/sqlx"
)

type DependencyContainer struct {
	connection                *sqlx.DB
	userHandler               *handlers.UserHandler
	authHandler               *handlers.HandlerAuth
	serviceHandler            *handlers.ServiceHandler
	productCategoryHandler    *handlers.ProductCategoryHandler
	login                     auth.Login
	userActions               actions.UserAction
	userQueryRepository       entities.UserQueryRepository
	userCommandRepository     entities.UserCommandRepository
	serviceRepository         entities.ServiceRepository
	productCategoryRepository entities.ProductCategoryRepository
}

func (d *DependencyContainer) GetProductCategoryHandler() *handlers.ProductCategoryHandler {
	if d.productCategoryHandler == nil {
		d.productCategoryHandler = handlers.NewProductCategoryHandler(
			d.GetProductCategoryRepository(),
		)
	}
	return d.productCategoryHandler
}

func (d *DependencyContainer) GetProductCategoryRepository() entities.ProductCategoryRepository {
	d.productCategoryRepository = mysql.NewProductCategoryRepository(
		d.GetDatabaseConnection(),
	)
	return d.productCategoryRepository
}

func (d *DependencyContainer) GetServiceHandler() *handlers.ServiceHandler {
	if d.serviceHandler == nil {
		d.serviceHandler = handlers.NewServiceHandler(
			d.GetServiceRepository(),
		)
	}
	return d.serviceHandler
}

func (d *DependencyContainer) GetServiceRepository() entities.ServiceRepository {
	d.serviceRepository = mysql.NewServiceRepository(
		d.GetDatabaseConnection(),
	)
	return d.serviceRepository
}

func (d *DependencyContainer) GetAuthHandler() *handlers.HandlerAuth {
	if d.authHandler == nil {
		d.authHandler = handlers.NewHandlerAuth(
			d.GetLogin(),
		)
	}
	return d.authHandler
}

func (d *DependencyContainer) GetLogin() auth.Login {
	d.login = *auth.NewLogin(
		d.GetUserQueryRepository(),
	)
	return d.login
}

func (d *DependencyContainer) GetUserHandler() *handlers.UserHandler {
	d.userHandler = handlers.NewUserHandler(
		d.GetUserActions(),
	)
	return d.userHandler
}

func (d *DependencyContainer) GetUserActions() actions.UserAction {
	d.userActions = *actions.NewUserAction(
		d.GetUserCommandRepository(),
		d.GetUserQueryRepository(),
	)

	return d.userActions
}

func (d *DependencyContainer) GetUserQueryRepository() entities.UserQueryRepository {
	if d.userQueryRepository == nil {
		d.userQueryRepository = mysql.NewUserQueryRepository(
			d.GetDatabaseConnection(),
		)
	}
	return d.userQueryRepository
}

func (d *DependencyContainer) GetUserCommandRepository() entities.UserCommandRepository {
	if d.userCommandRepository == nil {
		d.userCommandRepository = mysql.NewUserCommandRepository(
			d.GetDatabaseConnection(),
		)
	}
	return d.userCommandRepository
}

func (d *DependencyContainer) GetDatabaseConnection() *sqlx.DB {
	if d.connection == nil {
		d.connection = mysql.NewMysqlConnection()
	}
	return d.connection
}
