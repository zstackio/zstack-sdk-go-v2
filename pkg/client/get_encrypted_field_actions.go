// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetEncryptedField gets EncryptedField by uuid
func (cli *ZSClient) GetEncryptedField(uuid string) (*view.GetEncryptedFieldView, error) {
	var resp view.GetEncryptedFieldView
	if err := cli.Get("v1/encrypted/fields", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
