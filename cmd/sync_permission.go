//go:build devtools

package cmd

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/global"
	"myapp/loader"
	"myapp/manager"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func init() {
	rootCmd.AddCommand(newSyncPermissionCommand())
}

func newSyncPermissionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync-permission",
		Short: "Sync code permission with database permission",
		Long:  `This command will auto sync permission only (permissions, role_permissions and clinic_type_permissions). Developer need to manually adjust roles and clinic_types data and then run this command to sync the latest permission.`,
		Run: func(_ *cobra.Command, _ []string) {
			global.DisableDebug()

			container := manager.NewContainer(manager.LoadDefault)
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			ctx := context.Background()

			repositoryManager := container.RepositoryManager()

			syncPermissions(ctx, repositoryManager)
			syncRolePermissions(ctx, repositoryManager)
			syncClinicTypePermissions(ctx, repositoryManager)

			fmt.Println("Sync permission successful")
		},
	}

	return cmd
}

func syncPermissions(ctx context.Context, repositoryManager repository.RepositoryManager) {
	permissionRepository := repositoryManager.PermissionRepository()
	rolePermissionRepository := repositoryManager.RolePermissionRepository()

	existingPermissions, err := permissionRepository.Fetch(ctx)
	if err != nil {
		panic(err)
	}

	existingPermissionMap := map[data_type.Permission]model.Permission{}
	toBeDeletedPermissionArr := []model.Permission{}
	for _, existingPermission := range existingPermissions {
		if existingPermission.Title.IsValid() {
			existingPermissionMap[existingPermission.Title] = existingPermission
		} else {
			toBeDeletedPermissionArr = append(toBeDeletedPermissionArr, existingPermission)
		}
	}

	newPermissions := []model.Permission{}
	for _, permissionEnum := range data_type.ListPermission() {
		if existingPermission, exist := existingPermissionMap[permissionEnum]; exist {
			// if existingPermission.Type == permissionEnum.PermissionType() {
			// 	// keep existing permission
			// 	continue
			// }

			// delete old permission before create new permission
			toBeDeletedPermissionArr = append(toBeDeletedPermissionArr, existingPermission)
		}

		newPermissions = append(
			newPermissions,
			model.Permission{
				Id:          util.NewUuid(),
				Title:       permissionEnum,
				Description: strings.ToTitle(strings.ToLower(strings.Join(strings.Split(permissionEnum.String(), "_"), " "))),
				IsActive:    true,
				// Type:        permissionEnum.PermissionType(),
			},
		)
	}

	if len(toBeDeletedPermissionArr) > 0 {
		for _, toBeDeletedPermission := range toBeDeletedPermissionArr {
			if err := rolePermissionRepository.DeleteByPermissionId(ctx, toBeDeletedPermission.Id); err != nil {
				panic(err)
			}

			if err := permissionRepository.Delete(ctx, &toBeDeletedPermission, data_type.RepositoryOptionDisableAuditLog); err != nil {
				panic(err)
			}
		}
	}

	if len(newPermissions) > 0 {
		if err := permissionRepository.InsertMany(
			ctx,
			newPermissions,
			data_type.RepositoryOptionDisableAuditLog,
		); err != nil {
			panic(err)
		}
	}
}

func syncRolePermissions(ctx context.Context, repositoryManager repository.RepositoryManager) {
	permissionRepository := repositoryManager.PermissionRepository()
	roleRepository := repositoryManager.RoleRepository()
	rolePermissionRepository := repositoryManager.RolePermissionRepository()

	roles, err := roleRepository.Fetch(ctx)
	if err != nil {
		panic(err)
	}

	rolePermissionLoader := loader.NewRolePermissionsLoader(rolePermissionRepository)

	if err := util.Await(func(group *errgroup.Group) {
		for i := range roles {
			group.Go(rolePermissionLoader.RoleFn(&roles[i]))
		}
	}); err != nil {
		panic(err)
	}

	permissionLoader := loader.NewPermissionLoader(permissionRepository)

	if err := util.Await(func(group *errgroup.Group) {
		for i := range roles {
			for j := range roles[i].RolePermissions {
				group.Go(permissionLoader.RolePermissionFn(&roles[i].RolePermissions[j]))
			}
		}
	}); err != nil {
		panic(err)
	}

	for _, role := range roles {
		existingRolePermissionMap := map[data_type.Permission]model.RolePermission{}
		for _, rolePermission := range role.RolePermissions {
			existingRolePermissionMap[rolePermission.Permission.Title] = rolePermission
		}

		newPermissions := []data_type.Permission{}
		for _, permissionEnum := range role.Name.Permissions() {
			if _, exist := existingRolePermissionMap[permissionEnum]; exist {
				// keep existing permission by removing it from map
				delete(existingRolePermissionMap, permissionEnum)
				continue
			}

			newPermissions = append(newPermissions, permissionEnum)
		}

		if len(existingRolePermissionMap) > 0 {
			for _, rolePermission := range existingRolePermissionMap {
				if err := rolePermissionRepository.Delete(ctx, &rolePermission, data_type.RepositoryOptionDisableAuditLog); err != nil {
					panic(err)
				}
			}
		}

		if len(newPermissions) > 0 {
			permissions, err := permissionRepository.FetchByTitles(ctx, newPermissions)
			if err != nil {
				panic(err)
			}

			rolePermissions := []model.RolePermission{}
			for _, permission := range permissions {
				rolePermissions = append(
					rolePermissions,
					model.RolePermission{
						RoleId:       role.Id,
						PermissionId: permission.Id,
					},
				)
			}

			if len(rolePermissions) > 0 {
				if err := rolePermissionRepository.InsertMany(ctx, rolePermissions, data_type.RepositoryOptionDisableAuditLog); err != nil {
					panic(err)
				}
			}
		}
	}
}

func syncClinicTypePermissions(ctx context.Context, repositoryManager repository.RepositoryManager) {
	// permissionRepository := repositoryManager.PermissionRepository()
	// clinicTypeRepository := repositoryManager.ClinicTypeRepository()
	// clinicTypePermissionRepository := repositoryManager.ClinicTypePermissionRepository()

	// clinicTypes, err := clinicTypeRepository.Fetch(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// clinicTypePermissionLoader := loader.NewClinicTypePermissionsLoader(clinicTypePermissionRepository)

	// if err := util.Await(func(group *errgroup.Group) {
	// 	for i := range clinicTypes {
	// 		group.Go(clinicTypePermissionLoader.ClinicTypeFn(&clinicTypes[i]))
	// 	}
	// }); err != nil {
	// 	panic(err)
	// }

	// permissionLoader := loader.NewPermissionLoader(permissionRepository)

	// if err := util.Await(func(group *errgroup.Group) {
	// for i := range clinicTypes {
	// 	for j := range clinicTypes[i].ClinicTypePermissions {
	// 		group.Go(permissionLoader.ClinicTypePermissionFn(&clinicTypes[i].ClinicTypePermissions[j]))
	// 	}
	// }
	// }); err != nil {
	// 	panic(err)
	// }

	// for _, clinicType := range clinicTypes {
	// 	existingClinicTypePermissionMap := map[data_type.Permission]model.ClinicTypePermission{}
	// 	for _, clinicTypePermission := range clinicType.ClinicTypePermissions {
	// 		existingClinicTypePermissionMap[clinicTypePermission.Permission.Title] = clinicTypePermission
	// 	}

	// 	newPermissions := []data_type.Permission{}
	// 	for _, permissionEnum := range clinicType.Name.Permissions() {
	// 		if _, exist := existingClinicTypePermissionMap[permissionEnum]; exist {
	// 			// keep existing permission by removing it from map
	// 			delete(existingClinicTypePermissionMap, permissionEnum)
	// 			continue
	// 		}

	// 		newPermissions = append(newPermissions, permissionEnum)
	// 	}

	// 	if len(existingClinicTypePermissionMap) > 0 {
	// 		for _, clinicTypePermission := range existingClinicTypePermissionMap {
	// 			if err := clinicTypePermissionRepository.Delete(
	// 				ctx,
	// 				&clinicTypePermission,
	// 				data_type.RepositoryOptionDisableAuditLog,
	// 			); err != nil {
	// 				panic(err)
	// 			}
	// 		}
	// 	}

	// if len(newPermissions) > 0 {
	// 	permissions, err := permissionRepository.FetchByTitle(ctx, newPermissions)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	clinicTypePermissions := []model.ClinicTypePermission{}
	// 	for _, permission := range permissions {
	// 		clinicTypePermissions = append(
	// 			clinicTypePermissions,
	// 			model.ClinicTypePermission{
	// 				ClinicTypeId: clinicType.Id,
	// 				PermissionId: permission.Id,
	// 			},
	// 		)
	// 	}

	// 	if len(clinicTypePermissions) > 0 {
	// 		if err := clinicTypePermissionRepository.InsertMany(
	// 			ctx,
	// 			clinicTypePermissions,
	// 			data_type.RepositoryOptionDisableAuditLog,
	// 		); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
	// }
}
