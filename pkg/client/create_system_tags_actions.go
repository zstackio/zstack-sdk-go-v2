// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSystemTags creates SystemTags
func (cli *ZSClient) CreateSystemTags(params param.CreateSystemTagsParam) (*view.CreateSystemTagsEventView, error) {
	resp := view.CreateSystemTagsEventView{}
	if err := cli.Post("v1/system-tags/{resourceUuid}/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
