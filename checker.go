package authzclient

import (
	"context"
	"strconv"
	"strings"
	"time"

	pb "github.com/es-hs/erpc/authz"
)

func CheckPermission(userId uint64, shopId uint64, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.CheckPermission(ctx, &pb.CheckPermissionRequest{
		UserId: userId,
		ShopId: shopId,
		Act:    role,
	})
	if err != nil {
		return false, err
	}
	return r.Result, err
}

func AddRoleToDomain(userId uint64, shopId uint64, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.AddRoleToDomain(ctx, &pb.AddRoleToDomainRequest{
		UserId: userId,
		ShopId: shopId,
		Act:    role,
	})
	if err != nil {
		return false, err
	}
	return r.Result, err
}

func GetRolesInDomain(userId uint64, shopId uint64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetRolesInDomain(ctx, &pb.GetRolesInDomainRequest{
		UserId: userId,
		ShopId: shopId,
	})
	if err != nil {
		return nil, err
	}
	return r.Roles, err
}

func GetImplicitRolesInDomain(userId uint64, shopId uint64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetImplicitRolesInDomain(ctx, &pb.GetImplicitRolesInDomainRequest{
		UserId: userId,
		ShopId: shopId,
	})
	if err != nil {
		return nil, err
	}
	return r.Roles, err
}

func GenerateOwnerRole(userId uint64, shopId uint64) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GenerateOwnerRole(ctx, &pb.GenerateOwnerRoleRequest{
		UserId: userId,
		ShopId: shopId,
	})
	if err != nil {
		return 1, err
	}
	return uint64(r.Code), nil
}

func AddRolesForUserToDomain(userId uint64, shopId uint64, roles []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.AddRolesForUserToDomain(ctx, &pb.AddRolesForUserToDomainRequest{
		UserId: userId,
		ShopId: shopId,
		Act:    roles,
	})
	if err != nil {
		return false, err
	}
	return r.Result, nil
}

func RemoveRolesFromDomain(userId uint64, shopId uint64, roles []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.RemoveRolesFromDomain(ctx, &pb.RemoveRolesFromDomainRequest{
		UserId: userId,
		ShopId: shopId,
		Act:    roles,
	})
	if err != nil {
		return false, err
	}
	return r.Result, nil
}

func RemoveRoleFromDomain(userId uint64, shopId uint64, role string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.RemoveRoleFromDomain(ctx, &pb.RemoveRoleFromDomainRequest{
		UserId: userId,
		ShopId: shopId,
		Act:    role,
	})
	if err != nil {
		return false, err
	}
	return r.Result, err
}

func GetAllUsersByDomain(shopId uint64) ([]uint64, []uint64, error) {
	var userIds []uint64
	var partnerIds []uint64
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetAllUsersByDomain(ctx, &pb.GetAllUsersByDomainRequest{
		ShopId: shopId,
	})
	if err != nil {
		return nil, nil, err
	}
	for k := range r.UserIds {
		if strings.Contains(r.UserIds[k], "user_") {
			uID, _ := strconv.Atoi(r.UserIds[k][5:])
			userIds = append(userIds, uint64(uID))
		}
		if strings.Contains(r.UserIds[k], "partner_") {
			uID, _ := strconv.Atoi(r.UserIds[k][8:])
			partnerIds = append(partnerIds, uint64(uID))
		}
	}
	return userIds, partnerIds, err
}

func GetUsersForRoleInDomain(shopId uint64, role string) ([]uint64, []uint64, error) {
	var userIds []uint64
	var partnerIds []uint64
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.GetUsersForRoleInDomain(ctx, &pb.GetUsersForRoleInDomainRequest{
		ShopId: shopId,
		Role:   role,
	})
	if err != nil {
		return nil, nil, err
	}
	for k := range r.UserIds {
		if strings.Contains(r.UserIds[k], "user_") {
			uID, _ := strconv.Atoi(r.UserIds[k][5:])
			userIds = append(userIds, uint64(uID))
		}
		if strings.Contains(r.UserIds[k], "partner_") {
			uID, _ := strconv.Atoi(r.UserIds[k][8:])
			partnerIds = append(partnerIds, uint64(uID))
		}
	}

	return userIds, partnerIds, err
}

func DeleteDomains(shopId uint64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := C.DeleteDomains(ctx, &pb.DeleteDomainsRequest{
		ShopIds: []uint64{shopId},
	})
	if err != nil {
		return false, err
	}
	return r.Result, nil
}
