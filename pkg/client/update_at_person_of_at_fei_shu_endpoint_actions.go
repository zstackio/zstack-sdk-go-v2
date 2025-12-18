// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAtPersonOfAtFeiShuEndpoint updates AtPersonOfAtFeiShuEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtFeiShuEndpoint(uuid string, params param.UpdateAtPersonOfAtFeiShuEndpointParam) (*view.UpdateAtPersonOfFeiShuEndpointEventView, error) {
	resp := view.UpdateAtPersonOfFeiShuEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/feishu/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
