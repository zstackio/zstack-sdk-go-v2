// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSSOClient gets SSOClient by uuid
func (cli *ZSClient) GetSSOClient(uuid string) (*view.GetSSOClientView, error) {
	var resp view.GetSSOClientView
	if err := cli.Get("v1/get/sso/client", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
