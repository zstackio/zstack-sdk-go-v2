// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetElaborationCategories 获取ElaborationCategories详情
func (cli *ZSClient) GetElaborationCategories(uuid string) (*view.GetElaborationCategoriesView, error) {
	var resp view.GetElaborationCategoriesView
	if err := cli.Get("v1/errorcode/elaborations/categories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

