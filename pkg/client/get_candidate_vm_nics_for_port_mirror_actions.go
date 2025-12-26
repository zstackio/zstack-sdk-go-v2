// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateVmNicsForPortMirror gets CandidateVmNicsForPortMirror by uuid
func (cli *ZSClient) GetCandidateVmNicsForPortMirror(uuid string) (*view.GetCandidateVmNicsForPortMirrorView, error) {
	var resp view.GetCandidateVmNicsForPortMirrorView
	if err := cli.Get("v1/port-mirrors/{portMirrorUuid}/vm-instances/candidate-nics/{type}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
