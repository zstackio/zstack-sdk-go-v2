// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmStartingCandidateClustersHosts 获取VmStartingCandidateClustersHosts详情
func (cli *ZSClient) GetVmStartingCandidateClustersHosts(uuid string) (*view.GetVmStartingCandidateClustersHostsView, error) {
	var resp view.GetVmStartingCandidateClustersHostsView
	if err := cli.Get("v1/vm-instances/{uuid}/starting-target-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

