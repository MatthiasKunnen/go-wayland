package client

import "math"

// From wayland/wayland-util.h

func fixedToFloat64(f int32) float64 {
	u_i := int64((1023+44)<<52) + (1 << 51) + int64(f)
	u_d := math.Float64frombits(uint64(u_i))
	return u_d - (3 << 43)
}

func float64ToFixed(d float64) int32 {
	u_d := d + (3 << (51 - 8))
	u_i := int64(math.Float64bits(u_d))
	return int32(u_i)
}