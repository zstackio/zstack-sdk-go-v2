// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2VirtualIDAttribute 更新IAM2VirtualIDAttribute
func (cli *ZSClient) UpdateIAM2VirtualIDAttribute(uuid string, params param.UpdateIAM2VirtualIDAttributeParam) (*view.UpdateIAM2VirtualIDAttributeEventView, error) {
	resp := view.UpdateIAM2VirtualIDAttributeEventView{}
	if err := cli.Put("v1/iam2/virtual-ids/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

