// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSignatureServerEncryptPublicKey gets SignatureServerEncryptPublicKey by uuid
func (cli *ZSClient) GetSignatureServerEncryptPublicKey(uuid string) (*view.GetSignatureServerEncryptPublicKeyView, error) {
	var resp view.GetSignatureServerEncryptPublicKeyView
	if err := cli.Get("v1/secret-resource-pool-token/signature-server-encrypt-public-key", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
