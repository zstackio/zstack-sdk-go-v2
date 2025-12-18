// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// TriggerGCJob 操作TriggerGCJob
func (cli *ZSClient) TriggerGCJob(uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.Put("v1/gc-jobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

