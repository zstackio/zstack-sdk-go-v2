// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetMdevDeviceCandidates gets MdevDeviceCandidates by uuid
func (cli *ZSClient) GetMdevDeviceCandidates(uuid string) (*view.GetMdevDeviceCandidatesView, error) {
	var resp view.GetMdevDeviceCandidatesView
	if err := cli.Get("v1/mdev-devices/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
