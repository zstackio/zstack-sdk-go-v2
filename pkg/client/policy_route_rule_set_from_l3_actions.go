// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPolicyRouteRuleSetFromL3 操作PolicyRouteRuleSetFromL3
func (cli *ZSClient) DetachPolicyRouteRuleSetFromL3(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", uuid, string(deleteMode))
}

