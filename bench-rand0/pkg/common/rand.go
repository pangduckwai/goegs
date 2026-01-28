package common

func Rand02a(rnd uint64) int {
	if rnd < 9223372036854775807 {
		return 0
	} else {
		return 1
	}
}
func Rand02(rnd uint64) int {
	return int(rnd >> 63)
}

func Rand03(rnd uint64) int {
	if rnd < 6148914691236517205 {
		return 0
	} else if rnd < 12297829382473034410 {
		return 1
	} else {
		return 2
	}
}

func Rand04(rnd uint64) int {
	return int(rnd >> 62)
}

func Rand08(rnd uint64) int {
	return int(rnd >> 61)
}

func Rand10(rnd uint64) int {
	if rnd < 1844674407370955161 {
		return 0
	} else if rnd < 3689348814741910322 {
		return 1
	} else if rnd < 5534023222112865484 {
		return 2
	} else if rnd < 7378697629483820645 {
		return 3
	} else if rnd < 9223372036854775807 {
		return 4
	} else if rnd < 11068046444225730968 {
		return 5
	} else if rnd < 12912720851596686130 {
		return 6
	} else if rnd < 14757395258967641291 {
		return 7
	} else if rnd < 16602069666338596453 {
		return 8
	} else {
		return 9
	}
}

func Rand16(rnd uint64) int {
	return int(rnd >> 60)
}
