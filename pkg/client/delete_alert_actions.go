// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAlert deletes Alert
func (cli *ZSClient) DeleteAlert(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/alerts", uuid, string(deleteMode))
}
