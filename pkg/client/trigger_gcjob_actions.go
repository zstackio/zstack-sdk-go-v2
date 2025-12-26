// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// TriggerGCJob operates on TriggerGCJob
func (cli *ZSClient) TriggerGCJob(uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.Put("v1/gc-jobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
