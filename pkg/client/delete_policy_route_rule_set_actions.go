// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePolicyRouteRuleSet deletes PolicyRouteRuleSet
func (cli *ZSClient) DeletePolicyRouteRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/ruleSets/{uuid}", uuid, string(deleteMode))
}
