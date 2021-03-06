package main

func isPowerOfTwo(n int) bool {
	var hash = map[int]bool{
		1:                   true,
		2:                   true,
		4:                   true,
		8:                   true,
		16:                  true,
		32:                  true,
		64:                  true,
		128:                 true,
		256:                 true,
		512:                 true,
		1024:                true,
		2048:                true,
		4096:                true,
		8192:                true,
		16384:               true,
		32768:               true,
		65536:               true,
		131072:              true,
		262144:              true,
		524288:              true,
		1048576:             true,
		2097152:             true,
		4194304:             true,
		8388608:             true,
		16777216:            true,
		33554432:            true,
		67108864:            true,
		134217728:           true,
		268435456:           true,
		536870912:           true,
		1073741824:          true,
		2147483648:          true,
		4294967296:          true,
		8589934592:          true,
		17179869184:         true,
		34359738368:         true,
		68719476736:         true,
		137438953472:        true,
		274877906944:        true,
		549755813888:        true,
		1099511627776:       true,
		2199023255552:       true,
		4398046511104:       true,
		8796093022208:       true,
		17592186044416:      true,
		35184372088832:      true,
		70368744177664:      true,
		140737488355328:     true,
		281474976710656:     true,
		562949953421312:     true,
		1125899906842624:    true,
		2251799813685248:    true,
		4503599627370496:    true,
		9007199254740992:    true,
		18014398509481984:   true,
		36028797018963968:   true,
		72057594037927936:   true,
		144115188075855872:  true,
		288230376151711744:  true,
		576460752303423488:  true,
		1152921504606846976: true,
		2305843009213693952: true,
		4611686018427387904: true,
	}
	return hash[n]
}
