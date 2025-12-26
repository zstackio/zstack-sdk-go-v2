// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachTagToResources operates on TagToResources
func (cli *ZSClient) AttachTagToResources(params param.AttachTagToResourcesParam) (*view.AttachTagToResourcesEventView, error) {
	resp := view.AttachTagToResourcesEventView{}
	if err := cli.Post("v1/tags/{tagUuid}/resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
