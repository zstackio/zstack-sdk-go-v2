// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverResourceSplitBrain 操作ResourceSplitBrain
func (cli *ZSClient) RecoverResourceSplitBrain(uuid string, params param.RecoverResourceSplitBrainParam) (*view.RecoverResourceSplitBrainEventView, error) {
	resp := view.RecoverResourceSplitBrainEventView{}
	if err := cli.Put("v1/primary-storage/mini/{resourceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

