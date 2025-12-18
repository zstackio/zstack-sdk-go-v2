// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateVmNicsForPortMirror 获取CandidateVmNicsForPortMirror详情
func (cli *ZSClient) GetCandidateVmNicsForPortMirror(uuid string) (*view.GetCandidateVmNicsForPortMirrorView, error) {
	var resp view.GetCandidateVmNicsForPortMirrorView
	if err := cli.Get("v1/port-mirrors/{portMirrorUuid}/vm-instances/candidate-nics/{type}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

