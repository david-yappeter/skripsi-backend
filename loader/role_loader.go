package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type RoleLoader struct {
	loader dataloader.Loader
}

func (l *RoleLoader) load(id string) (*model.Role, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Role), nil
}

func (l *RoleLoader) UserRoleFn(userRole *model.UserRole) func() error {
	return func() error {
		role, err := l.load(userRole.RoleId)
		if err != nil {
			return err
		}

		userRole.Role = role

		return nil
	}
}

func NewRoleLoader(roleRepository repository.RoleRepository) *RoleLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		roles, err := roleRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		roleById := map[string]model.Role{}
		for _, role := range roles {
			roleById[role.Id] = role
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var role *model.Role
			if v, ok := roleById[k.String()]; ok {
				role = &v
			}

			result := &dataloader.Result{Data: role, Error: nil}
			if role == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &RoleLoader{
		loader: NewDataloader(batchFn),
	}
}
