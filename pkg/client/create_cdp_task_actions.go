// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateCdpTask creates CdpTask
func (cli *ZSClient) CreateCdpTask(params param.CreateCdpTaskParam) (*view.CreateCdpTaskEventView, error) {
	resp := view.CreateCdpTaskEventView{}
	if err := cli.Post("v1/cdp-backup-storage/task", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
