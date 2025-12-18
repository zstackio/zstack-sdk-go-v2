// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunNasAccessGroup deletes AliyunNasAccessGroup
func (cli *ZSClient) DeleteAliyunNasAccessGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nas/access/{uuid}", uuid, string(deleteMode))
}
