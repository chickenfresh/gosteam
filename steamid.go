package gosteam

import "strconv"

func (s *SteamID) ToString() string {
	return strconv.FormatUint(uint64(*s), 10)
}
