// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateMiniHosts 获取CandidateMiniHosts详情
func (cli *ZSClient) GetCandidateMiniHosts(uuid string) (*view.GetCandidateMiniHostsView, error) {
	var resp view.GetCandidateMiniHostsView
	if err := cli.Get("v1/mini-clusters/candidate-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

