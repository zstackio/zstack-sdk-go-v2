// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangePrimaryStorageState 操作PrimaryStorageState
func (cli *ZSClient) ChangePrimaryStorageState(uuid string, params param.ChangePrimaryStorageStateParam) (*view.ChangePrimaryStorageStateEventView, error) {
	resp := view.ChangePrimaryStorageStateEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

