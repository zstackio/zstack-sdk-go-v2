// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSignatureServerEncryptPublicKey 获取SignatureServerEncryptPublicKey详情
func (cli *ZSClient) GetSignatureServerEncryptPublicKey(uuid string) (*view.GetSignatureServerEncryptPublicKeyView, error) {
	var resp view.GetSignatureServerEncryptPublicKeyView
	if err := cli.Get("v1/secret-resource-pool-token/signature-server-encrypt-public-key", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

