// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachHostFromHostSchedulingRuleGroup 操作HostFromHostSchedulingRuleGroup
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hostSchedulingRuleGroup/{hostGroupUuid}/host", uuid, string(deleteMode))
}

