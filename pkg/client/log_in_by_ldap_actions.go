// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LogInByLdap operates on LogInByLdap
func (cli *ZSClient) LogInByLdap(uuid string, params param.LogInByLdapParam) (*view.LogInByLdapView, error) {
	resp := view.LogInByLdapView{}
	if err := cli.Put("v1/ldap/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
