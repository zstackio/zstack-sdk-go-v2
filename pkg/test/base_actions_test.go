package test

import (
	"context"
	"os"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/client"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

const (
	// The ZStack Cloud Community Edition only supports login authentication for super admin accounts.
	// The ZStack Cloud Basic Edition supports login authentication for AccessKey, super admin, and sub-accounts.
	// The ZStack Cloud Enterprise Edition supports login authentication for AccessKey, super admin, sub-accounts, and enterprise users.

	accountLoginHostname        = "172.25.16.98" //ZStack Cloud API endpoint IP address
	accountLoginAccountName     = "admin"
	accountLoginAccountPassword = "password"

	accountLoginMasterHostname = "IPOfCloudAPIEndpoint"
	accountLoginSlaveHostname  = "IPOfCloudAPIEndpoint"

	accessKeyAuthHostname        = "172.26.100.254"
	accessKeyAuthAccessKeyId     = "ak"
	accessKeyAuthAccessKeySecret = "sk"

	contextPath = "zstack"

	readOnly = false
	debug    = false
)

var accountLoginCli = client.NewZSClient(
	client.DefaultZSConfig(accountLoginHostname).
		LoginAccount(accountLoginAccountName, accountLoginAccountPassword).
		ReadOnly(readOnly).
		Debug(true),
)

var accessKeyAuthCli = client.NewZSClient(
	client.DefaultZSConfig(accessKeyAuthHostname).
		AccessKey(accessKeyAuthAccessKeyId, accessKeyAuthAccessKeySecret).
		ReadOnly(readOnly).
		Debug(debug),
)

var loginSession *view.SessionInventoryView

func TestMain(m *testing.M) {
	var err error
	loginSession, err = accountLoginCli.Login(context.Background(), )
	if err != nil {
		os.Exit(1) // 登录失败直接退出，不运行测试
	}
	defer accountLoginCli.Logout(context.Background(), )
	code := m.Run()
	os.Exit(code)
}
