// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResumeVmInstance 操作ResumeVmInstance
func (cli *ZSClient) ResumeVmInstance(uuid string, params param.ResumeVmInstanceParam) (*view.ResumeVmInstanceEventView, error) {
	resp := view.ResumeVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

