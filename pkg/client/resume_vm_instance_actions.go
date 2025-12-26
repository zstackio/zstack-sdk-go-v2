// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResumeVmInstance operates on ResumeVmInstance
func (cli *ZSClient) ResumeVmInstance(uuid string, params param.ResumeVmInstanceParam) (*view.ResumeVmInstanceEventView, error) {
	resp := view.ResumeVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
