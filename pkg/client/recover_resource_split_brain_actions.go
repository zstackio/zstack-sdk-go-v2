// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverResourceSplitBrain operates on ResourceSplitBrain
func (cli *ZSClient) RecoverResourceSplitBrain(uuid string, params param.RecoverResourceSplitBrainParam) (*view.RecoverResourceSplitBrainEventView, error) {
	resp := view.RecoverResourceSplitBrainEventView{}
	if err := cli.Put("v1/primary-storage/mini/{resourceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
