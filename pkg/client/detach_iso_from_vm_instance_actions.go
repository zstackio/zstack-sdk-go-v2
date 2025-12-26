// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachIsoFromVmInstance operates on IsoFromVmInstance
func (cli *ZSClient) DetachIsoFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/iso", uuid, string(deleteMode))
}
