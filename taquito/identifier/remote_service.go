package identifier

import (
	"log"

	"github.com/sul-dlss-labs/identifier-service/generated/client"
	"github.com/PenguinParadigm/samvera18apis/taquito/config"
	"github.com/PenguinParadigm/samvera18apis/taquito/datautils"
)

type remoteIdentifierService struct {
	TransportConfig *client.TransportConfig
}

// NewRemoteIdentifierService creates a new instance of the identifier service
func NewRemoteIdentifierService(config *config.Config) Service {
	host := config.IdentifierServiceHost
	return &remoteIdentifierService{
		TransportConfig: client.DefaultTransportConfig().WithHost(host),
	}
}

func (d *remoteIdentifierService) Mint(resource *datautils.Resource) (string, error) {
	c := client.NewHTTPClientWithConfig(nil, d.TransportConfig)
	ok, err := c.Operations.MintNewDRUIDS(nil)
	if err != nil {
		log.Printf("[ERROR] Unable to get an identifier from the remote service.")
		return "", err
	}
	return string(ok.Payload[0]), nil
}
