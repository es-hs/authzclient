/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"log"
	"time"

	"github.com/es-hs/authzclient"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// address     = "3.0.95.112:50051"
	// address     = "authz.gempages.xyz:50051"
	defaultName = "world"
)

func main() {
	err := authzclient.InitAuthClient(address, 1*time.Second, grpc.WithInsecure(), grpc.WithBlock())
	defer authzclient.Conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	// Contact the server and print out its response.
	t1 := time.Now()
	// r, err := c.AddRoleToDomain(ctx, &pb.AddRoleToDomainRequest{
	// 	UserId: 2009,
	// 	ShopId: 2011,
	// 	Act:    authzclient.OWNER_ROLE,
	// })
	r, err := authzclient.AddRoleToDomain(2009, 2011, authzclient.OWNER_ROLE)
	log.Println(time.Since(t1))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Result ", r)
	//check roles list
	// r2, err := c.GetRolesInDomain(ctx, &pb.GetRolesInDomainRequest{
	// 	UserId: 2009,
	// 	ShopId: 2011,
	// })
	roles, err := authzclient.GetRolesInDomain(2009, 2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(roles)
	log.Println(time.Since(t1))
	//check roles detail list
	// r3, err := c.GetImplicitRolesInDomain(ctx, &pb.GetImplicitRolesInDomainRequest{
	// 	UserId: 2009,
	// 	ShopId: 2011,
	// })
	roles, err = authzclient.GetImplicitRolesInDomain(2009, 2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(roles)
	log.Println(time.Since(t1))

	//check roles detail list
	// r4, err := c.CheckPermission(ctx, &pb.CheckPermissionRequest{
	// 	UserId: 2009,
	// 	ShopId: 2011,
	// 	Act:    authzclient.LOGIN_PERMISSION,
	// })
	r, err = authzclient.CheckPermission(2009, 2011, authzclient.LOGIN_PERMISSION)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r)
	log.Println(time.Since(t1))
	//gemerate role for

	// r5, err := c.GenerateOwnerRole(ctx, &pb.GenerateOwnerRoleRequest{
	// 	UserId: 2009,
	// 	ShopId: 2011,
	// })
	_, _ = authzclient.GenerateOwnerRole(2008, 2011)
	r5, err := authzclient.GenerateOwnerRole(2009, 2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r5)
	log.Println(time.Since(t1))

	r6, err := authzclient.AddRolesForUserToDomain(2009, 2011, []string{authzclient.ADMIN_ROLE, authzclient.PRODUCT_READ})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r6)
	log.Println(time.Since(t1))

	r6, err = authzclient.AddRolesForUserToDomain(2010, 2011, []string{authzclient.ADMIN_ROLE, authzclient.PRODUCT_READ})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r6)
	log.Println(time.Since(t1))

	//show
	log.Println("list role of 2009 in domain 2011")
	roles, err = authzclient.GetRolesInDomain(2009, 2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(roles)
	log.Println(time.Since(t1))
	//
	r7, err := authzclient.RemoveRoleFromDomain(2009, 2011, authzclient.OWNER_ROLE)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r7)
	log.Println(time.Since(t1))

	//show
	roles, err = authzclient.GetRolesInDomain(2009, 2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(roles)
	log.Println(time.Since(t1))
	//

	// r8, err := authzclient.RemoveRolesFromDomain(2009, 2011, []string{authzclient.ADMIN_ROLE, authzclient.PRODUCT_READ})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Println(r8)
	// log.Println(time.Since(t1))

	//show
	roles, err = authzclient.GetRolesInDomain(2009, 2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(roles)
	log.Println(time.Since(t1))
	//

	//show partner list and role list
	log.Println("list user ")
	users, partners, err := authzclient.GetAllUsersByDomain(2011)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(users)
	log.Println(partners)
	log.Println(time.Since(t1))

	//show partner list and role list by role
	log.Println("list user with role ", authzclient.ADMIN_ROLE)
	users, partners, err = authzclient.GetUsersForRoleInDomain(2011, authzclient.ADMIN_ROLE)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(users)
	log.Println(partners)

	log.Println("list user with role ", authzclient.OWNER_ROLE)
	users, partners, err = authzclient.GetUsersForRoleInDomain(2011, authzclient.OWNER_ROLE)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(users)
	log.Println(partners)
	log.Println(time.Since(t1))
}
