package seeder

import (
	"context"
	"myapp/data_type"
	"myapp/model"
	"myapp/repository"
)

var (
	RoleSuperAdmin = model.Role{
		Id:          "8d01b6df-d26b-4b50-81df-b01167bedf0c",
		Name:        data_type.RoleSuperAdmin,
		Description: nil,
	}
	RoleCashier = model.Role{
		Id:          "bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10",
		Name:        data_type.RoleCashier,
		Description: nil,
	}
	RoleDriver = model.Role{
		Id:          "00c12bd8-7470-40c5-938e-029b1239650c",
		Name:        data_type.RoleDriver,
		Description: nil,
	}
	RoleInventory = model.Role{
		Id:          "73b6799d-ebd6-491f-9b47-272fb0f22914",
		Name:        data_type.RoleInventory,
		Description: nil,
	}
)

func RoleSeeder(repositoryManager repository.RepositoryManager) {
	roleRepository := repositoryManager.RoleRepository()

	count, err := roleRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := roleRepository.InsertMany(context.Background(), getRoleData()); err != nil {
		panic(err)
	}
}

func getRoleData() []model.Role {
	return []model.Role{
		RoleSuperAdmin,
		RoleCashier,
		RoleDriver,
		RoleInventory,
	}
}
