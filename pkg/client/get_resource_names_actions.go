// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceNames gets ResourceNames by uuid
func (cli *ZSClient) GetResourceNames(uuid string) (*view.GetResourceNamesView, error) {
	var resp view.GetResourceNamesView
	if err := cli.Get("v1/resources/names", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
