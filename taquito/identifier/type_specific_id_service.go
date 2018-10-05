package identifier

import (
	"github.com/PenguinParadigm/samvera18apis/taquito/datautils"
)

type TypeSpecificIDService struct {
	localService  Service
	remoteService Service
}

// return a DRUID if this is a collection or DRO, otherwise a UUID
func (d *TypeSpecificIDService) Mint(resource *datautils.Resource) (string, error) {
	if resource.IsObject() || resource.IsCollection() {
		// Objects and collections get DRUIDS
		return d.remoteService.Mint(resource)
	}

	// other types gets UUIDs
	return d.localService.Mint(resource)
}
