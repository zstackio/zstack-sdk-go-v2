// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachProvisionNicFromBonding operates on ProvisionNicFromBonding
func (cli *ZSClient) DetachProvisionNicFromBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/bm-instances/bm2-bondings/{uuid}", uuid, string(deleteMode))
}
