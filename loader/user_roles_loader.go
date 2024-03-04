package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type UserRolesLoader struct {
	loaderByUserId dataloader.Loader
}

func (l *UserRolesLoader) loadByUserId(id string) ([]model.UserRole, error) {
	thunk := l.loaderByUserId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.UserRole), nil
}

func (l *UserRolesLoader) UserFn(user *model.User) func() error {
	return func() error {
		if user != nil {
			userRoles, err := l.loadByUserId(user.Id)
			if err != nil {
				return err
			}

			user.UserRoles = userRoles
		}

		return nil
	}
}

func NewUserRolesLoader(userRoleRepository repository.UserRoleRepository) *UserRolesLoader {
	batchByUserIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceives, err := userRoleRepository.FetchByUserIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		userRoleByUserId := map[string][]model.UserRole{}
		for _, userRole := range productReceives {
			userRoleByUserId[userRole.UserId] = append(userRoleByUserId[userRole.UserId], userRole)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var userRoles []model.UserRole
			if v, ok := userRoleByUserId[k.String()]; ok {
				userRoles = v
			}

			result := &dataloader.Result{Data: userRoles, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &UserRolesLoader{
		loaderByUserId: NewDataloader(batchByUserIdFn),
	}
}
