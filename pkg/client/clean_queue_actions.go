// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CleanQueue operates on CleanQueue
func (cli *ZSClient) CleanQueue(uuid string, params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.Put("v1/clean/queue", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
