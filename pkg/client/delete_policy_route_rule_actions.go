// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePolicyRouteRule deletes PolicyRouteRule
func (cli *ZSClient) DeletePolicyRouteRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rules/{uuid}", uuid, string(deleteMode))
}
