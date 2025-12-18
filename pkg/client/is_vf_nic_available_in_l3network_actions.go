// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// IsVfNicAvailableInL3Network 操作IsVfNicAvailableInL3Network
func (cli *ZSClient) IsVfNicAvailableInL3Network(params param.IsVfNicAvailableInL3NetworkParam) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/hosts/{hostUuid}/vfnicavailable", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

