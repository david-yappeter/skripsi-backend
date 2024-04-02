package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	UserSuperAdminOne = model.User{
		Id:       "53e457f9-7d11-4e24-8bb5-a5134be9e2e7",
		Username: "super.admin.one",
		Name:     "Super Admin One",
		Password: "$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG",
		IsActive: true,
	}

	UserInventoryOne = model.User{
		Id:       "1e7a5eb3-c5c4-4fa3-b971-168765bfc413",
		Username: "inventory.one",
		Name:     "Inventory One",
		Password: "$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG",
		IsActive: true,
	}

	UserCashierOne = model.User{
		Id:       "68ed7124-fa08-4720-b741-9fe4fa697c21",
		Username: "cashier.one",
		Name:     "Cashier One",
		Password: "$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG",
		IsActive: true,
	}

	UserDriverOne = model.User{
		Id:       "fbfcbc34-77b9-4901-82f1-e5fa78d5aa48",
		Username: "driver.one",
		Name:     "Driver One",
		Password: "$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG",
		IsActive: true,
	}
)

func UserSeeder(repositoryManager repository.RepositoryManager) {
	userRepository := repositoryManager.UserRepository()

	count, err := userRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := userRepository.InsertMany(context.Background(), getUserData()); err != nil {
		panic(err)
	}
}

func getUserData() []model.User {
	return []model.User{
		UserSuperAdminOne,
		UserInventoryOne,
		UserCashierOne,
		UserDriverOne,
	}
}
