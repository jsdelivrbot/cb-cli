package cli

import (
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/utils"
	proxyConfig "github.com/hortonworks/cb-cli/client_cloudbreak/v3_organization_id_proxyconfigs"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

type proxyClient interface {
	CreateProxyconfigInOrganization(params *proxyConfig.CreateProxyconfigInOrganizationParams) (*proxyConfig.CreateProxyconfigInOrganizationOK, error)
	ListProxyconfigsByOrganization(params *proxyConfig.ListProxyconfigsByOrganizationParams) (*proxyConfig.ListProxyconfigsByOrganizationOK, error)
}

var ProxyHeader = []string{"Name", "Host", "Port", "Protocol", "User"}

type proxy struct {
	Name     string `json:"Name" yaml:"Name"`
	Host     string `json:"Host" yaml:"Host"`
	Port     string `json:"Port" yaml:"Port"`
	Protocol string `json:"Protocol" yaml:"Protocol"`
	User     string `json:"User" yaml:"User"`
}

func (p *proxy) DataAsStringArray() []string {
	return []string{p.Name, p.Host, p.Port, p.Protocol, p.User}
}

func CreateProxy(c *cli.Context) error {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "create proxy")

	orgID := c.Int64(FlOrganizationOptional.Name)
	name := c.String(FlName.Name)
	host := c.String(FlProxyHost.Name)
	port := c.String(FlProxyPort.Name)
	protocol := c.String(FlProxyProtocol.Name)
	user := c.String(FlProxyUser.Name)
	password := c.String(FlProxyPassword.Name)

	if protocol != "http" && protocol != "https" {
		utils.LogErrorMessageAndExit("Proxy protocol must be either http or https")
	}
	serverPort, _ := strconv.Atoi(port)

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	return createProxy(cbClient.Cloudbreak.V3OrganizationIDProxyconfigs, orgID, name, host, int32(serverPort), protocol, user, password)
}

func createProxy(proxyClient proxyClient, orgID int64, name, host string, port int32, protocol, user, password string) error {
	proxyRequest := &models_cloudbreak.ProxyConfigRequest{
		Name:       &name,
		ServerHost: &host,
		ServerPort: &port,
		Protocol:   &protocol,
		UserName:   user,
		Password:   password,
	}

	log.Infof("[createProxy] create proxy with name: %s", name)
	var proxy *models_cloudbreak.ProxyConfigResponse
	resp, err := proxyClient.CreateProxyconfigInOrganization(proxyConfig.NewCreateProxyconfigInOrganizationParams().WithOrganizationID(orgID).WithBody(proxyRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	proxy = resp.Payload

	log.Infof("[createProxy] proxy created with name: %s, id: %d", name, proxy.ID)
	return nil
}

func ListProxies(c *cli.Context) error {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list proxies")

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	orgID := c.Int64(FlOrganizationOptional.Name)
	return listProxiesImpl(cbClient.Cloudbreak.V3OrganizationIDProxyconfigs, output.WriteList, orgID)
}

func listProxiesImpl(proxyClient proxyClient, writer func([]string, []utils.Row), orgID int64) error {
	resp, err := proxyClient.ListProxyconfigsByOrganization(proxyConfig.NewListProxyconfigsByOrganizationParams().WithOrganizationID(orgID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, p := range resp.Payload {
		row := &proxy{
			Name:     *p.Name,
			Host:     *p.ServerHost,
			Port:     strconv.Itoa(int(*p.ServerPort)),
			Protocol: *p.Protocol,
			User:     p.UserName,
		}
		tableRows = append(tableRows, row)
	}

	writer(ProxyHeader, tableRows)
	return nil
}

func DeleteProxy(c *cli.Context) error {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "delete a proxy")

	orgID := c.Int64(FlOrganizationOptional.Name)
	proxyName := c.String(FlName.Name)
	log.Infof("[DeleteProxy] delete proxy config by name: %s", proxyName)

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	if _, err := cbClient.Cloudbreak.V3OrganizationIDProxyconfigs.DeleteProxyconfigInOrganization(proxyConfig.NewDeleteProxyconfigInOrganizationParams().WithOrganizationID(orgID).WithName(proxyName)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteProxy] proxy config deleted: %s", proxyName)
	return nil
}