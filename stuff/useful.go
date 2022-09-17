package usefulthings

type UniswapV2pair struct {
	TokenA    Address
	TokenB    Address
	Pool      Address
	ReservesA Big.Int
	ReservesB Big.Int
}
