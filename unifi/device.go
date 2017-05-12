package unifi

import (
	"github.com/ZAP-Quebec/unifi-inform/data"
)

type Device interface {
	Mac() data.MacAddr
	AuthKey() data.Key
	SetManagementConfig(data.ManagementConfig) error
}
