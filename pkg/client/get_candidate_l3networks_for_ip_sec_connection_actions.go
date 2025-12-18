// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateL3NetworksForIpSecConnection gets CandidateL3NetworksForIpSecConnection by uuid
func (cli *ZSClient) GetCandidateL3NetworksForIpSecConnection(uuid string) (*view.GetCandidateL3NetworksForIpSecConnectionView, error) {
	var resp view.GetCandidateL3NetworksForIpSecConnectionView
	if err := cli.Get("v1/ipsec/candidatesL3Networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
