// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExecuteDRSScheduling 操作ExecuteDRSScheduling
func (cli *ZSClient) ExecuteDRSScheduling(uuid string, params param.ExecuteDRSSchedulingParam) (*view.ExecuteDRSSchedulingEventView, error) {
	resp := view.ExecuteDRSSchedulingEventView{}
	if err := cli.Put("v1/clusters/drs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

