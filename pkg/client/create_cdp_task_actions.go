// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateCdpTask creates CdpTask
func (cli *ZSClient) CreateCdpTask(params param.CreateCdpTaskParam) (*view.CreateCdpTaskEventView, error) {
	resp := view.CreateCdpTaskEventView{}
	if err := cli.Post("v1/cdp-backup-storage/task", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
