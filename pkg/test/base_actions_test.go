package test

import (
	"os"
	"testing"

	"github.com/kataras/golog"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/client"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

const (
	// The ZStack Cloud Community Edition only supports login authentication for super admin accounts.
	// The ZStack Cloud Basic Edition supports login authentication for AccessKey, super admin, and sub-accounts.
	// The ZStack Cloud Enterprise Edition supports login authentication for AccessKey, super admin, sub-accounts, and enterprise users.

	accountLoginHostname        = "172.24.249.239" //ZStack Cloud API endpoint IP address
	accountLoginAccountName     = "admin"
	accountLoginAccountPassword = "password"

	accountLoginMasterHostname = "IPOfCloudAPIEndpoint"
	accountLoginSlaveHostname  = "IPOfCloudAPIEndpoint"

	accessKeyAuthHostname        = ""
	accessKeyAuthAccessKeyId     = "3YWXy79yktjCzUaY3Xfz"
	accessKeyAuthAccessKeySecret = "secret"

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
	loginSession, err = accountLoginCli.Login()
	if err != nil {
		golog.Errorf("TestMain err %v", err)
	}
	defer accountLoginCli.Logout()

	m.Run()
	os.Exit(0)
}
