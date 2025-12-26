// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPolicyRouteRuleSetFromL3 operates on PolicyRouteRuleSetFromL3
func (cli *ZSClient) DetachPolicyRouteRuleSetFromL3(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", uuid, string(deleteMode))
}
