package main

import (
	"fmt"
	"log"
	"math/big"

	UniswapQuery "ethclient/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func marketsoftoken(token common.Address, marketdata [5]UniswapV2pair) []UniswapV2pair {
	var d []UniswapV2pair
	for i := 0; i < len(marketdata); i++ {
		if marketdata[i].TokenA == token || marketdata[i].TokenB == token {
			// fmt.println(marketdata[i].Pool)
			d = append(d, marketdata[i])

		}
	}
	return d
}

func swap(howmuch *big.Int, ofthis common.Address, pooldata UniswapV2pair) Swapstruct {
	var f int64 = 997
	var h int64 = 1000
	g := big.NewInt(h)
	fee := big.NewInt(f)
	inwithfee := fee.Mul(fee, howmuch)
	if ofthis == pooldata.TokenA {
		var numer *big.Int
		var denom *big.Int
		var a *big.Int
		var quot *big.Int
		numer.Mul(pooldata.ReserveB, inwithfee)
		a.Mul(pooldata.ReserveA, g)
		denom.Add(a, inwithfee)
		quot.Div(numer, denom)
		return Swapstruct{amountOut: quot, tokenOut: pooldata.TokenB}

	} else {
		var numer *big.Int
		var denom *big.Int
		var a *big.Int
		var quot *big.Int
		numer.Mul(pooldata.ReserveA, inwithfee)
		a.Mul(pooldata.ReserveB, g)
		denom.Add(a, inwithfee)
		quot.Div(numer, denom)

		return Swapstruct{amountOut: quot, tokenOut: pooldata.TokenA}
	}

}

type UniswapV2pair struct {
	TokenA        common.Address
	TokenB        common.Address
	Pool          common.Address
	ReserveA      *big.Int
	ReserveB      *big.Int
	Infofromblock *big.Int
}

type Swapstruct struct {
	amountOut *big.Int
	tokenOut  common.Address
}

type beststuff struct {
	initialToken common.Address
	gain         *big.Int
	path         []common.Address
	poolpath     []UniswapV2pair
}

func main() {
	Flashable := [26]string{"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", "0x6B175474E89094C44Da98b954EedeAC495271d0F", "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e", "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", "0xdAC17F958D2ee523a2206206994597C13D831ec7", "0xE41d2489571d322189246DaFA5ebDe1F4699F498", "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9", "0x0D8775F648430679A709E98d2b0Cb6250d2887EF", "0x4Fabb145d64652a948d72533023f6E7A623C7C53", "0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c", "0xdd974D5C2e2928deA5F71b9825b8b646686BD200", "0x514910771AF9Ca656af840dff83E8264EcF986CA", "0x0F5D2fB29fb7d3CFeE444a200298f468908cC942", "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2", "0x408e41876cCCDC0F92210600ef50372656052a38", "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F", "0x57Ab1ec28D129707052df4dF418D58a2D46d5f51", "0x0000000000085d4780B73119b644AE5ecd22b376", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", "0xD533a949740bb3306d119CC777fa900bA034cd52", "0x056Fd409E1d7A124BD7017459dFEa2F387b6d5Cd", "0xba100000625a3754423978a60c9317c58a424e3D", "0x8798249c2E607446EfB7Ad49eC89dD1865Ff4272", "0xD5147bc8e386d91Cc5DBE72099DAC6C9b99276F5", "0x03ab458634910AaD20eF5f1C8ee96F1D6ac54919"}
	// fmt.Println(Flashable)
	var flashaddr [26]common.Address
	for i := 0; i < len(Flashable); i++ {
		flashaddr[i] = common.HexToAddress(Flashable[i])
	}
	// fmt.Println(flashaddr)

	// var blacklist [9]common.Address
	blstr := [9]string{"0x02ba7b2026D26896bc1368e6bEf4349d2f595B68", "0x566113069683Ce664958A784f18336B4020a1350", "0xa16001DD47f505B7B7c5639c710A52209E4e8904", "0x8707dEDB55f4AB2d17Ae849061bF034334d1c8a0", "0x86FADb80d8D2cff3C3680819E4da99C10232Ba0F", "0x4D13d624a87baa278733c068A174412AfA9ca6C8", "0x905D3237dC71F7D8f604778e8b78f0c3ccFF9377", "0xDADA00A9C23390112D08a1377cc59f7d03D9df55", "0x0Ba45A8b5d5575935B8158a88C631E9F9C95a2e5"}
	var blacklist [9]common.Address
	for i := 0; i < 9; i++ {
		blacklist[i] = common.HexToAddress(blstr[i])
	}
	// fmt.Println(blacklist)

	// }
	// blacklist := [9]common.Address{common.HexToAddress("0x02ba7b2026D26896bc1368e6bEf4349d2f595B68"), common.HexToAddress("0x566113069683Ce664958A784f18336B4020a1350"), common.HexToAddress("0xa16001DD47f505B7B7c5639c710A52209E4e8904"), common.HexToAddress("0x8707dEDB55f4AB2d17Ae849061bF034334d1c8a0"), common.HexToAddress("0x86FADb80d8D2cff3C3680819E4da99C10232Ba0F"), common.HexToAddress("0x4D13d624a87baa278733c068A174412AfA9ca6C8"), common.HexToAddress("0x905D3237dC71F7D8f604778e8b78f0c3ccFF9377"), common.HexToAddress("0xDADA00A9C23390112D08a1377cc59f7d03D9df55"), common.HexToAddress("0x0Ba45A8b5d5575935B8158a88C631E9F9C95a2e5")}
	// fmt.Println(blacklist)
	// badtokens := common.HexToAddress(Blacklist)
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/cb7a603124c4411ba12877599e494814")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")

	uniswapv2query := common.HexToAddress("0x5EF1009b9FCD4fec3094a5564047e190D72Bd511")
	instance, err := UniswapQuery.NewUniswapFlashQuery(uniswapv2query, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("loaded contract 1")

	facaddr := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("loaded contract 2")

	const n int64 = 5

	startindex := big.NewInt(0)
	endindex := big.NewInt(n)
	fmt.Println("we got big ints")

	query, err := instance.GetPairsByIndexRange(&bind.CallOpts{Pending: false}, facaddr, startindex, endindex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("query succsess")

	var pairlist = make([]common.Address, n)
	var i int64
	for i = 0; i < n; i++ {
		pairlist[i] = query[i][2]
	}
	// fmt.Println(pairlist)

	reserves, err := instance.GetReservesByPairs(&bind.CallOpts{Pending: false}, pairlist)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(reserves[0][0])
	// fmt.Println(reserves[0][1])
	// fmt.Println(reserves[0][2])

	var pools [n]UniswapV2pair

	for i = 0; i < n; i++ {
		pools[i] = UniswapV2pair{TokenA: query[i][0], TokenB: query[i][1], Pool: query[i][2], ReserveA: reserves[i][0], ReserveB: reserves[i][1], Infofromblock: reserves[i][2]}
	}
	fmt.Println(pools[0].TokenA)

	fmt.Println(marketsoftoken(flashaddr[0], pools))

	amounts64 := [4]int64{1, 2, 5, 10}
	amounts := [4]*big.Int{big.NewInt(amounts64[0]), big.NewInt(amounts64[1]), big.NewInt(amounts64[2]), big.NewInt(amounts64[3])}
	// for q, e := range amounts64 {
	// amounts[q] = big.NewInt(*e)
	// }
	var zero64 int64 = 0
	bigzero := big.NewInt(zero64)

	var best []beststuff
	for _, token := range flashaddr {
		for _, amount := range amounts {
			pool1 := marketsoftoken(token, pools)
			for j := 0; j < len(pool1); j++ {
				swap1 := swap(amount, token, pool1[j])
				pool2 := marketsoftoken(swap1.tokenOut, pools)
				for k := 0; k < len(pool2); k++ {
					swap2 := swap(swap1.amountOut, swap1.tokenOut, pool2[k])
					if swap2.tokenOut == token {
						var profit *big.Int
						profit.Sub(swap2.amountOut, amount)
						if profit.Cmp(bigzero) == 1 {
							best = append(best, beststuff{initialToken: token, gain: profit, path: []common.Address{token, swap1.tokenOut, swap2.tokenOut}, poolpath: []UniswapV2pair{pool1[j], pool2[k]}})
						}
					} else {
						pool3 := marketsoftoken(swap2.tokenOut, pools)
						for u := 0; u < len(pool3); u++ {
							swap3 := swap(swap2.amountOut, swap2.tokenOut, pool3[u])
							if swap3.tokenOut == token {
								var profit *big.Int
								profit.Sub(swap3.amountOut, amount)
								if profit.Cmp(bigzero) == 1 {
									best = append(best, beststuff{initialToken: token, gain: profit, path: []common.Address{token, swap1.tokenOut, swap2.tokenOut, swap3.tokenOut}, poolpath: []UniswapV2pair{pool1[j], pool2[k], pool3[u]}})
								}
							}
						}
					}
				}
			}
		}
	}
}

// fmt.Println(pools)
// fmt.Println(pools[0].TokenA)
// fmt.Println(pools[0].TokenB)
// fmt.Println(pools[0].Pool)
// fmt.Println(pools[0].ReserveA)
// fmt.Println(pools[0].ReserveB)
// fmt.Println(pools[0].infofromblock)

// var pooldepth = map[common.Address]

// Pools := new([1000]common.Address)
// Pools := new([n][3]common.Address)
// for i, j := range query {
// 	for _, b := range blacklist {
// 		if j[0] == b || j[1] == b {
// 			fmt.Println("bad token found", b)
// 		} else {
// 			Pools[i] = j

// 		}
//
// 	}
// }
// fmt.Println(Pools[0])
// fmt.Println(Pools)

// fmt.Println(i)
// fmt.Println(i, j)
