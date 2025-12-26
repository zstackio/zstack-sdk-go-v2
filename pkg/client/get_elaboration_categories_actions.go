// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetElaborationCategories gets ElaborationCategories by uuid
func (cli *ZSClient) GetElaborationCategories(uuid string) (*view.GetElaborationCategoriesView, error) {
	var resp view.GetElaborationCategoriesView
	if err := cli.Get("v1/errorcode/elaborations/categories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
