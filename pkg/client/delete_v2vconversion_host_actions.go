// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteV2VConversionHost deletes V2VConversionHost
func (cli *ZSClient) DeleteV2VConversionHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/v2v-conversion-hosts/{uuid}", uuid, string(deleteMode))
}
