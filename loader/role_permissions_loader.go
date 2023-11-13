package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type RolePermissionsLoader struct {
	loader dataloader.Loader
}

func (l *RolePermissionsLoader) load(id string) ([]model.RolePermission, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.RolePermission), nil
}

func (l *RolePermissionsLoader) RoleFn(role *model.Role) func() error {
	return func() error {
		rolePermissions, err := l.load(role.Id)
		if err != nil {
			return err
		}

		role.RolePermissions = rolePermissions

		return nil
	}
}

func NewRolePermissionsLoader(rolePermissionRepository repository.RolePermissionRepository) *RolePermissionsLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		roleIds := make([]string, len(keys))
		for idx, k := range keys {
			roleIds[idx] = k.String()
		}

		rolePermissions, err := rolePermissionRepository.FetchByRoleIds(ctx, roleIds)
		if err != nil {
			panic(err)
		}

		rolePermissionsByRoleId := map[string][]model.RolePermission{}
		for _, rolePermission := range rolePermissions {
			rolePermissionsByRoleId[rolePermission.RoleId] = append(rolePermissionsByRoleId[rolePermission.RoleId], rolePermission)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			rolePermissions := []model.RolePermission{}
			if v, ok := rolePermissionsByRoleId[k.String()]; ok {
				rolePermissions = v
			}

			results[idx] = &dataloader.Result{Data: rolePermissions, Error: nil}
		}

		return results
	}

	return &RolePermissionsLoader{
		loader: NewDataloader(batchFn),
	}
}
