// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveIAM2ProjectLoginExpired removes IAM2ProjectLoginExpired
func (cli *ZSClient) RemoveIAM2ProjectLoginExpired(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/remove/login-expired/{uuid}/actions", uuid, string(deleteMode))
}
