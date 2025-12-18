// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddNvmeServer adds NvmeServer
func (cli *ZSClient) AddNvmeServer(params param.AddNvmeServerParam) (*view.AddNvmeServerEventView, error) {
	resp := view.AddNvmeServerEventView{}
	if err := cli.Post("v1/storage-devices/nvme/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
