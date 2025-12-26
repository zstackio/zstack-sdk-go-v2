// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePolicyRouteRuleSet deletes PolicyRouteRuleSet
func (cli *ZSClient) DeletePolicyRouteRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/ruleSets/{uuid}", uuid, string(deleteMode))
}
