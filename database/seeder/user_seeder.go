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
	}
}
