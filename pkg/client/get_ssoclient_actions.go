// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSSOClient gets SSOClient by uuid
func (cli *ZSClient) GetSSOClient(uuid string) (*view.GetSSOClientView, error) {
	var resp view.GetSSOClientView
	if err := cli.Get("v1/get/sso/client", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
