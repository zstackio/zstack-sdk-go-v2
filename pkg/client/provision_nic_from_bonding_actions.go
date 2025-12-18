// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachProvisionNicFromBonding 操作ProvisionNicFromBonding
func (cli *ZSClient) DetachProvisionNicFromBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/bm-instances/bm2-bondings/{uuid}", uuid, string(deleteMode))
}

