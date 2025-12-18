// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetEncryptedField 获取EncryptedField详情
func (cli *ZSClient) GetEncryptedField(uuid string) (*view.GetEncryptedFieldView, error) {
	var resp view.GetEncryptedFieldView
	if err := cli.Get("v1/encrypted/fields", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

