// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch deletes ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/connections/{uuid}", uuid, string(deleteMode))
}
