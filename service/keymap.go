package service

import "mousk/infra/keyboardctl"

type KeymapService struct{}

func (i *KeymapService) GetValidKeycodes() map[string]uint32 {
	nkmap := keyboardctl.ExportNameKeycodeMap()
	return nkmap
}
