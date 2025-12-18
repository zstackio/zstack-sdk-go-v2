// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAtPersonOfAtWeComEndpoint updates AtPersonOfAtWeComEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtWeComEndpoint(uuid string, params param.UpdateAtPersonOfAtWeComEndpointParam) (*view.UpdateAtPersonOfWeComEndpointEventView, error) {
	resp := view.UpdateAtPersonOfWeComEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/we-com/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
