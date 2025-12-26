// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSystemTag creates SystemTag
func (cli *ZSClient) CreateSystemTag(params param.CreateSystemTagParam) (*view.CreateSystemTagEventView, error) {
	resp := view.CreateSystemTagEventView{}
	if err := cli.Post("v1/system-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
