// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSchedulerJob updates SchedulerJob
func (cli *ZSClient) UpdateSchedulerJob(uuid string, params param.UpdateSchedulerJobParam) (*view.UpdateSchedulerJobEventView, error) {
	resp := view.UpdateSchedulerJobEventView{}
	if err := cli.Put("v1/scheduler/jobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
