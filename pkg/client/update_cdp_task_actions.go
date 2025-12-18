// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(uuid string, params param.UpdateCdpTaskParam) (*view.UpdateCdpTaskEventView, error) {
	resp := view.UpdateCdpTaskEventView{}
	if err := cli.Put("v1/cdp-backup-storage/task/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
