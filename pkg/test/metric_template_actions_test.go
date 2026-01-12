// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMetricTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMetricTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMetricTemplate error: %v", err)
		return
	}
	golog.Infof("QueryMetricTemplate result count: %d", len(result))
}

func TestDeleteMetricTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMetricTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMetricTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMetricTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MetricTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMetricTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMetricTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteMetricTemplate succeeded for UUID: %s", list[0].UUID)
}

func TestCreateMetricTemplate(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateMetricTemplate is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateMetricTemplateParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateMetricTemplateParamDetail{
	// 		Name: "test-metrictemplate",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateMetricTemplate(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateMetricTemplate error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateMetricTemplate result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteMetricTemplate(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteMetricTemplate error: %v", err)
	// }
}
