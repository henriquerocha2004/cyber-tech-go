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
	connection             *sqlx.DB
	
	userHandler            *handlers.UserHandler
	authHandler            *handlers.HandlerAuth
	serviceHandler         *handlers.ServiceHandler
	productCategoryHandler *handlers.ProductCategoryHandler
	productHandler         *handlers.ProductHandler
	supplierHandler        *handlers.SupplierHandler
	stockHandler           *handlers.StockHandler
	orderStatusHandler     *handlers.OrderServiceStatusHandler
	serviceOrderHandler    *handlers.OrderServiceHandler

	login               auth.Login
	userActions         actions.UserAction
	stockActions        actions.StockActions
	orderServiceActions actions.ServiceOrderActions

	userQueryRepository           entities.UserQueryRepository
	userCommandRepository         entities.UserCommandRepository
	serviceRepository             entities.ServiceRepository
	productCategoryRepository     entities.ProductCategoryRepository
	productRepository             entities.ProductRepository
	supplierRepository            entities.SupplierRepository
	stockRepository               entities.StockRepository
	orderStatusRepository         entities.OrderServiceStatusRepository
	orderServiceCommandRepository entities.OrderServiceCommandRepository
	orderServiceQueryRepository   entities.OrderServiceQueryRepository
}

func (d *DependencyContainer) GetOrderServiceHandler() *handlers.OrderServiceHandler {
	if d.serviceOrderHandler == nil {
		d.serviceOrderHandler = handlers.NewOrderServiceHandler(
			d.GetOrderServiceActions(),
		)
	}
	return d.serviceOrderHandler
}

func (d *DependencyContainer) GetOrderServiceActions() actions.ServiceOrderActions {
	d.orderServiceActions = *actions.NewServiceOrderActions(
		d.GetOrderServiceCommandRepository(),
		d.GetOrderServiceQueryRepository(),
		d.GetOrderStatusRepository(),
		d.GetProductRepository(),
	)
	return d.orderServiceActions
}

func (d *DependencyContainer) GetOrderServiceQueryRepository() entities.OrderServiceQueryRepository {
	d.orderServiceQueryRepository = mysql.NewOrderServiceQueryRepository(
		d.GetDatabaseConnection(),
	)
	return d.orderServiceQueryRepository
}

func (d *DependencyContainer) GetOrderServiceCommandRepository() entities.OrderServiceCommandRepository {
	d.orderServiceCommandRepository = mysql.NewOrderServiceCommandRepository(
		d.GetDatabaseConnection(),
	)
	return d.orderServiceCommandRepository
}

func (d *DependencyContainer) GetOrderStatusHandler() *handlers.OrderServiceStatusHandler {
	if d.orderStatusHandler == nil {
		d.orderStatusHandler = handlers.NewOrderServiceStatusHandler(
			d.GetOrderStatusRepository(),
		)
	}
	return d.orderStatusHandler
}

func (d *DependencyContainer) GetOrderStatusRepository() entities.OrderServiceStatusRepository {
	d.orderStatusRepository = mysql.NewOrderServiceStatusRepository(
		d.GetDatabaseConnection(),
	)
	return d.orderStatusRepository
}

func (d *DependencyContainer) GetStockHandler() *handlers.StockHandler {
	if d.stockHandler == nil {
		d.stockHandler = handlers.NewStockHandler(
			d.GetStockActions(),
		)
	}
	return d.stockHandler
}

func (d *DependencyContainer) GetStockActions() actions.StockActions {
	d.stockActions = *actions.NewStockActions(
		d.GetStockRepository(),
	)
	return d.stockActions
}

func (d *DependencyContainer) GetStockRepository() entities.StockRepository {
	d.stockRepository = mysql.NewStockRepository(
		d.GetDatabaseConnection(),
	)
	return d.stockRepository
}

func (d *DependencyContainer) GetSupplierHandler() *handlers.SupplierHandler {
	if d.supplierHandler == nil {
		d.supplierHandler = handlers.NewSupplierHandler(
			d.GetSupplierRepository(),
		)
	}
	return d.supplierHandler
}

func (d *DependencyContainer) GetSupplierRepository() entities.SupplierRepository {
	d.supplierRepository = mysql.NewSupplierRepository(
		d.GetDatabaseConnection(),
	)
	return d.supplierRepository
}

func (d *DependencyContainer) GetProductHandler() *handlers.ProductHandler {
	if d.productHandler == nil {
		d.productHandler = handlers.NewProductHandler(
			d.GetProductRepository(),
		)
	}
	return d.productHandler
}

func (d *DependencyContainer) GetProductRepository() entities.ProductRepository {
	d.productRepository = mysql.NewProductRepository(
		d.GetDatabaseConnection(),
	)
	return d.productRepository
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
