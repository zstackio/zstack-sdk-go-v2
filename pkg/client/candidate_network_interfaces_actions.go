// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateNetworkInterfaces 获取CandidateNetworkInterfaces详情
func (cli *ZSClient) GetCandidateNetworkInterfaces(uuid string) (*view.GetCandidateNetworkInterfacesView, error) {
	var resp view.GetCandidateNetworkInterfacesView
	if err := cli.Get("v1/cluster/hosts-network-interfaces", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

