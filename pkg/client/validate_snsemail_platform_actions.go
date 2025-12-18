// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateSNSEmailPlatform operates on ValidateSNSEmailPlatform
func (cli *ZSClient) ValidateSNSEmailPlatform(uuid string, params param.ValidateSNSEmailPlatformParam) (*view.ValidateSNSEmailPlatformEventView, error) {
	resp := view.ValidateSNSEmailPlatformEventView{}
	if err := cli.Put("v1/sns/application-platforms/email/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
