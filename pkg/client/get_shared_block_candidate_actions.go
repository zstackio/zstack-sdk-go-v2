// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSharedBlockCandidate gets SharedBlockCandidate by uuid
func (cli *ZSClient) GetSharedBlockCandidate(uuid string) (*view.GetSharedBlockCandidateView, error) {
	var resp view.GetSharedBlockCandidateView
	if err := cli.Get("v1/primary-storage/sharedblockgroup/sharedblock-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
