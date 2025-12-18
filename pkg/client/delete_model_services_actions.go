// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteModelServices deletes ModelServices
func (cli *ZSClient) DeleteModelServices(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/", uuid, string(deleteMode))
}
