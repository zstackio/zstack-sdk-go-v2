// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachTagToResources 操作TagToResources
func (cli *ZSClient) AttachTagToResources(params param.AttachTagToResourcesParam) (*view.AttachTagToResourcesEventView, error) {
	resp := view.AttachTagToResourcesEventView{}
	if err := cli.Post("v1/tags/{tagUuid}/resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

