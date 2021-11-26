package gosteam

type SteamID uint64

func (s SteamID) GetAccountID() uint32 {
	return uint32(s.get(0, 0xFFFFFFFF))
}

func (s SteamID) get(offset uint, mask uint64) uint64 {
	return (uint64(s) >> offset) & mask
}
