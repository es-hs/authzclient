package authzclient

import (
	"context"
	"strconv"
	"strings"
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

func AddRolesForUserToDomain(userId uint, shopId uint, roles []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.AddRolesForUserToDomain(ctx, &pb.AddRolesForUserToDomainRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
		Act:    roles,
	})
	if err != nil {
		return false, err
	}
	return r.Result, nil
}

func RemoveRolesFromDomain(userId uint, shopId uint, roles []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.RemoveRolesFromDomain(ctx, &pb.RemoveRolesFromDomainRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
		Act:    roles,
	})
	if err != nil {
		return false, err
	}
	return r.Result, nil
}

func RemoveRoleFromDomain(userId uint, shopId uint, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.RemoveRoleFromDomain(ctx, &pb.RemoveRoleFromDomainRequest{
		UserId: int64(userId),
		ShopId: int64(shopId),
		Act:    role,
	})
	if err != nil {
		return false, err
	}
	return r.Result, err
}

func GetAllUsersByDomain(shopId uint) ([]uint, []uint, error) {
	var userIds []uint
	var partnerIds []uint
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetAllUsersByDomain(ctx, &pb.GetAllUsersByDomainRequest{
		ShopId: int64(shopId),
	})
	if err != nil {
		return nil, nil, err
	}
	for k := range r.UserIds {
		if strings.Contains(r.UserIds[k], "user_") {
			uID, _ := strconv.Atoi(r.UserIds[k][5:])
			userIds = append(userIds, uint(uID))
		}
		if strings.Contains(r.UserIds[k], "partner_") {
			uID, _ := strconv.Atoi(r.UserIds[k][8:])
			partnerIds = append(partnerIds, uint(uID))
		}
	}
	return userIds, partnerIds, err
}

func GetUsersForRoleInDomain(shopId uint, role string) ([]uint, []uint, error) {
	var userIds []uint
	var partnerIds []uint
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetUsersForRoleInDomain(ctx, &pb.GetUsersForRoleInDomainRequest{
		ShopId: int64(shopId),
		Role:   role,
	})
	if err != nil {
		return nil, nil, err
	}
	for k := range r.UserIds {
		if strings.Contains(r.UserIds[k], "user_") {
			uID, _ := strconv.Atoi(r.UserIds[k][5:])
			userIds = append(userIds, uint(uID))
		}
		if strings.Contains(r.UserIds[k], "partner_") {
			uID, _ := strconv.Atoi(r.UserIds[k][8:])
			partnerIds = append(partnerIds, uint(uID))
		}
	}

	return userIds, partnerIds, err
}

func DeleteDomains(shopId uint) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.DeleteDomains(ctx, &pb.DeleteDomainsRequest{
		ShopIds: []int64{int64(shopId)},
	})
	if err != nil {
		return false, err
	}
	return r.Result, nil
}
