package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	userRoleSuperAdminOne = model.UserRole{
		UserId: UserSuperAdminOne.Id,
		RoleId: RoleSuperAdmin.Id,
	}
	userRoleSuperAdminTwo = model.UserRole{
		UserId: UserSuperAdminOne.Id,
		RoleId: RoleInventory.Id,
	}
	userRoleInventoryOne = model.UserRole{
		UserId: UserInventoryOne.Id,
		RoleId: RoleInventory.Id,
	}
	userRoleCashierOne = model.UserRole{
		UserId: UserCashierOne.Id,
		RoleId: RoleCashier.Id,
	}
	userRoleDriverOne = model.UserRole{
		UserId: UserDriverOne.Id,
		RoleId: RoleDriver.Id,
	}
)

func UserRoleSeeder(repositoryManager repository.RepositoryManager) {
	userRoleRepository := repositoryManager.UserRoleRepository()

	count, err := userRoleRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := userRoleRepository.InsertMany(context.Background(), getUserRoleData()); err != nil {
		panic(err)
	}
}

func getUserRoleData() []model.UserRole {
	return []model.UserRole{
		userRoleSuperAdminOne,
		userRoleSuperAdminTwo,
		userRoleInventoryOne,
		userRoleCashierOne,
		userRoleDriverOne,
	}
}
