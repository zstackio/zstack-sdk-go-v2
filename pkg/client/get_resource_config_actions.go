// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceConfig gets ResourceConfig by uuid
func (cli *ZSClient) GetResourceConfig(uuid string) (*view.GetResourceConfigView, error) {
	var resp view.GetResourceConfigView
	if err := cli.Get("v1/resource-configurations/{resourceUuid}/{category}/{name}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
