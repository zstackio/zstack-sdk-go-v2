// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangePrimaryStorageState changes PrimaryStorageState
func (cli *ZSClient) ChangePrimaryStorageState(uuid string, params param.ChangePrimaryStorageStateParam) (*view.ChangePrimaryStorageStateEventView, error) {
	resp := view.ChangePrimaryStorageStateEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
