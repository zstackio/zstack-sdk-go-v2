// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeImageState changes ImageState
func (cli *ZSClient) ChangeImageState(uuid string, params param.ChangeImageStateParam) (*view.ChangeImageStateEventView, error) {
	resp := view.ChangeImageStateEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
