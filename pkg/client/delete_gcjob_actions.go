// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteGCJob deletes GCJob
func (cli *ZSClient) DeleteGCJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/gc-jobs/{uuid}", uuid, string(deleteMode))
}
