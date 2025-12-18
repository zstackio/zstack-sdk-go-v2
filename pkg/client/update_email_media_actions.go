// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEmailMedia updates EmailMedia
func (cli *ZSClient) UpdateEmailMedia(uuid string, params param.UpdateEmailMediaParam) (*view.UpdateEmailMediaEventView, error) {
	resp := view.UpdateEmailMediaEventView{}
	if err := cli.Put("v1/media/emails/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
