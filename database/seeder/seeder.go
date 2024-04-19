package seeder

import (
	"fmt"
	"myapp/model"
	"myapp/repository"
)

var Seeders = map[string]func(repositoryManager repository.RepositoryManager){
	model.CustomerTypeTableName: CustomerTypeSeeder,
	model.CustomerTableName:     CustomerSeeder,
	model.FileTableName:         FileSeeder,
	model.ProductTableName:      ProductSeeder,
	model.ProductStockTableName: ProductStockSeeder,
	model.ProductUnitTableName:  ProductUnitSeeder,
	model.RoleTableName:         RoleSeeder,
	model.ShopeeConfigTableName: ShopeeConfigSeeder,
	model.SupplierTypeTableName: SupplierTypeSeeder,
	model.SupplierTableName:     SupplierSeeder,
	model.TiktokConfigTableName: TiktokConfigSeeder,
	model.UserTableName:         UserSeeder,
	model.UserRoleTableName:     UserRoleSeeder,
	model.UnitTableName:         UnitSeeder,
}

func Seed(repositoryManager repository.RepositoryManager, tableName string) {
	if seed, exist := Seeders[tableName]; exist {
		seed(repositoryManager)
	} else {
		fmt.Printf("Seeder for table `%s` not found\n", tableName)
	}
}

func SeedAll(repositoryManager repository.RepositoryManager) {
	seedOrders := []string{
		model.RoleTableName,
		model.UserTableName,
		model.UserRoleTableName,

		model.TiktokConfigTableName,
		model.ShopeeConfigTableName,

		model.FileTableName,

		model.UnitTableName,
		model.ProductTableName,
		model.ProductUnitTableName,
		model.ProductStockTableName,

		model.CustomerTypeTableName,
		model.CustomerTableName,

		model.SupplierTypeTableName,
		model.SupplierTableName,
	}

	for _, tableName := range seedOrders {
		seed, ok := Seeders[tableName]
		if !ok {
			panic(fmt.Errorf("table name %s not found", tableName))
		}

		seed(repositoryManager)
	}
}
