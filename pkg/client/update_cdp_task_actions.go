// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(uuid string, params param.UpdateCdpTaskParam) (*view.UpdateCdpTaskEventView, error) {
	resp := view.UpdateCdpTaskEventView{}
	if err := cli.Put("v1/cdp-backup-storage/task/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
