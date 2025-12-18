// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMaaSUsage gets MaaSUsage by uuid
func (cli *ZSClient) GetMaaSUsage(uuid string) (*view.GetMaaSUsageView, error) {
	var resp view.GetMaaSUsageView
	if err := cli.Get("v1/maas/usage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
