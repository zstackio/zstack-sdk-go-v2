// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetEncryptedField gets EncryptedField by uuid
func (cli *ZSClient) GetEncryptedField() (*view.GetEncryptedFieldView, error) {
	var resp view.GetEncryptedFieldView
	if err := cli.GetWithRespKey("v1/encrypted/fields", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
