// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StartDataProtection starts DataProtection
func (cli *ZSClient) StartDataProtection(params param.StartDataProtectionParam) (*view.StartDataProtectionEventView, error) {
	resp := view.StartDataProtectionEventView{}
	if err := cli.Post("v1/start/data/protection/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
