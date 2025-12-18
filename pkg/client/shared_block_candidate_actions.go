// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSharedBlockCandidate 获取SharedBlockCandidate详情
func (cli *ZSClient) GetSharedBlockCandidate(uuid string) (*view.GetSharedBlockCandidateView, error) {
	var resp view.GetSharedBlockCandidateView
	if err := cli.Get("v1/primary-storage/sharedblockgroup/sharedblock-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

