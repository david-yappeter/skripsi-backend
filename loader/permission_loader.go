package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type PermissionLoader struct {
	loader dataloader.Loader
}

func (l *PermissionLoader) load(id string) (*model.Permission, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Permission), nil
}

func (l *PermissionLoader) RolePermissionFn(rolePermission *model.RolePermission) func() error {
	return func() error {
		permission, err := l.load(rolePermission.PermissionId)
		if err != nil {
			return err
		}

		rolePermission.Permission = permission

		return nil
	}
}

func NewPermissionLoader(permissionRepository repository.PermissionRepository) *PermissionLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		permissions, err := permissionRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		permissionById := map[string]model.Permission{}
		for _, permission := range permissions {
			permissionById[permission.Id] = permission
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var permission *model.Permission
			if v, ok := permissionById[k.String()]; ok {
				permission = &v
			}

			result := &dataloader.Result{Data: permission, Error: nil}
			if permission == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &PermissionLoader{
		loader: NewDataloader(batchFn),
	}
}
