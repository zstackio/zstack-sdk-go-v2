// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetNoTriggerSchedulerJobs 获取NoTriggerSchedulerJobs详情
func (cli *ZSClient) GetNoTriggerSchedulerJobs(uuid string) (*view.GetNoTriggerSchedulerJobsView, error) {
	var resp view.GetNoTriggerSchedulerJobsView
	if err := cli.Get("v1/scheduler/jobs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

