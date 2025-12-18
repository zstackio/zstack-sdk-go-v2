// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateL3NetworksForIpSecConnection 获取CandidateL3NetworksForIpSecConnection详情
func (cli *ZSClient) GetCandidateL3NetworksForIpSecConnection(uuid string) (*view.GetCandidateL3NetworksForIpSecConnectionView, error) {
	var resp view.GetCandidateL3NetworksForIpSecConnectionView
	if err := cli.Get("v1/ipsec/candidatesL3Networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

