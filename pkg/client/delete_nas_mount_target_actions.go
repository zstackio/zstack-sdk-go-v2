// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteNasMountTarget deletes NasMountTarget
func (cli *ZSClient) DeleteNasMountTarget(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/mount/{uuid}", uuid, string(deleteMode))
}
