// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteEcsSecurityGroupRuleRemote 删除EcsSecurityGroupRuleRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRuleRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group-rule/remote/{uuid}", uuid, string(deleteMode))
}

