// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SsoClientPushData operates on SsoClientPushData
func (cli *ZSClient) SsoClientPushData(uuid string, params param.SsoClientPushDataParam) (*view.SsoClientPushDataEventView, error) {
	resp := view.SsoClientPushDataEventView{}
	if err := cli.Put("v1/sso/resource/data/push", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
