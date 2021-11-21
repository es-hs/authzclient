package authzclient

import (
	"context"
	"time"

	pb "github.com/es-hs/erpc/authz"
)

func CheckPermission(userId uint, shopId uint, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.CheckPermission(ctx, &pb.CheckPermissionRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
		Act:    role,
	})
	if err != nil {
		return false, err
	}
	return r.Result, err
}

func AddRoleToDomain(userId uint, shopId uint, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.AddRoleToDomain(ctx, &pb.AddRoleToDomainRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
		Act:    role,
	})
	if err != nil {
		return false, err
	}
	return r.Result, err
}

func GetRolesInDomain(userId uint, shopId uint) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetRolesInDomain(ctx, &pb.GetRolesInDomainRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
	})
	if err != nil {
		return nil, err
	}
	return r.Roles, err
}

func GetImplicitRolesInDomain(userId uint, shopId uint) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetImplicitRolesInDomain(ctx, &pb.GetImplicitRolesInDomainRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
	})
	if err != nil {
		return nil, err
	}
	return r.Roles, err
}

func GenerateOwnerRole(userId uint, shopId uint) (uint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GenerateOwnerRole(ctx, &pb.GenerateOwnerRoleRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
	})
	if err != nil {
		return 1, err
	}
	return uint(r.Code), nil
}
