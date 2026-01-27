//go:build !linux

package device

import (
	"github.com/giangv11/wireguard-go/conn"
	"github.com/giangv11/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
