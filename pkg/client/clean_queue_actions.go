// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanQueue operates on CleanQueue
func (cli *ZSClient) CleanQueue(uuid string, params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.Put("v1/clean/queue", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
