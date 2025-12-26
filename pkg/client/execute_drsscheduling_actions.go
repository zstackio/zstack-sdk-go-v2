// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExecuteDRSScheduling operates on ExecuteDRSScheduling
func (cli *ZSClient) ExecuteDRSScheduling(uuid string, params param.ExecuteDRSSchedulingParam) (*view.ExecuteDRSSchedulingEventView, error) {
	resp := view.ExecuteDRSSchedulingEventView{}
	if err := cli.Put("v1/clusters/drs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
