// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteResourceConfig deletes ResourceConfig
func (cli *ZSClient) DeleteResourceConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/resource-configurations/{category}/{name}/{resourceUuid}", uuid, string(deleteMode))
}
