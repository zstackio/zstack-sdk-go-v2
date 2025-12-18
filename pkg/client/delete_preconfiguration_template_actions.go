// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePreconfigurationTemplate deletes PreconfigurationTemplate
func (cli *ZSClient) DeletePreconfigurationTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/preconfigurations/{uuid}", uuid, string(deleteMode))
}
