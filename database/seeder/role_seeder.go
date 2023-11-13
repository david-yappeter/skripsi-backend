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
	}
}
