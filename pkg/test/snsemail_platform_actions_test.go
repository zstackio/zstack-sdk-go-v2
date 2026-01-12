// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEmailPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSEmailPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailPlatform result count: %d", len(result))
}

func TestCreateSNSEmailPlatform(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSEmailPlatform is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSEmailPlatformParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSEmailPlatformParamDetail{
	// 		Name: "test-snsemailplatform",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSEmailPlatform(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSEmailPlatform error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSEmailPlatform result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSEmailPlatform(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSEmailPlatform error: %v", err)
	// }
}

func TestValidateSNSEmailPlatform(t *testing.T) {
	// ValidateSNSEmailPlatform operation
	t.Skip("TestValidateSNSEmailPlatform requires manual implementation")

}
