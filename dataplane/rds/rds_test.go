package rds

import (
	"strings"
	"testing"

	"github.com/hortonworks/cb-cli/dataplane/api/client/v3_workspace_id_rdsconfigs"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	"github.com/hortonworks/cb-cli/dataplane/types"
	"github.com/hortonworks/dp-cli-common/utils"
)

type mockRdsClient struct {
}

func (*mockRdsClient) ListRdsConfigsByWorkspace(params *v3_workspace_id_rdsconfigs.ListRdsConfigsByWorkspaceParams) (*v3_workspace_id_rdsconfigs.ListRdsConfigsByWorkspaceOK, error) {
	resp := []*model.RDSConfigResponse{
		{
			Name:             &(&types.S{S: "test"}).S,
			Description:      &(&types.S{S: "description"}).S,
			ConnectionURL:    &(&types.S{S: "connectionURL"}).S,
			DatabaseEngine:   &(&types.S{S: "databaseEngine"}).S,
			Type:             &(&types.S{S: "type"}).S,
			ConnectionDriver: nil,
		},
	}
	return &v3_workspace_id_rdsconfigs.ListRdsConfigsByWorkspaceOK{Payload: resp}, nil
}

func (*mockRdsClient) CreateRdsConfigInWorkspace(params *v3_workspace_id_rdsconfigs.CreateRdsConfigInWorkspaceParams) (*v3_workspace_id_rdsconfigs.CreateRdsConfigInWorkspaceOK, error) {
	return nil, nil
}

func (*mockRdsClient) TestRdsConnectionInWorkspace(params *v3_workspace_id_rdsconfigs.TestRdsConnectionInWorkspaceParams) (*v3_workspace_id_rdsconfigs.TestRdsConnectionInWorkspaceOK, error) {
	return nil, nil
}

func TestListRdsImpl(t *testing.T) {
	var rows []utils.Row
	listAllRdsImpl(new(mockRdsClient), func(h []string, r []utils.Row) { rows = r }, int64(2))
	if len(rows) != 1 {
		t.Fatalf("row number doesn't match 1 == %d", len(rows))
	}
	for _, r := range rows {
		expected := "test description connectionURL databaseEngine type  "
		if strings.Join(r.DataAsStringArray(), " ") != expected {
			t.Errorf("row data not match %s == %s", expected, strings.Join(r.DataAsStringArray(), " "))
		}
	}
}
