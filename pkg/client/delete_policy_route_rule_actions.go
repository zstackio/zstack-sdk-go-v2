// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePolicyRouteRule deletes PolicyRouteRule
func (cli *ZSClient) DeletePolicyRouteRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rules/{uuid}", uuid, string(deleteMode))
}
