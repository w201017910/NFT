package util

import (
	"fmt"
	"nft/database"
	"strconv"
	"time"
)

var weekTime int = 604800
var mouthTime int = 2592000

func Week(t []database.Transaction) float64 {
	if len(t) == 0 {
		return 0
	}
	limitTime := int(time.Now().Unix()) - weekTime
	var endPrice int
	for _, v := range t {
		if limitTime < v.Times {
			newPrice, _ := strconv.Atoi(v.TokenPrice)
			endPrice = newPrice
		}

	}

	price1, _ := strconv.Atoi(t[0].TokenPrice)

	a := float64(endPrice) / float64(price1)
	fmt.Println(len(t))
	if a >= 1 {
		a = a*100 - 100
	} else {
		a = -(100 - a*100)
	}
	return a
}
func Mouth(t []database.Transaction) float64 {
	if len(t) == 0 {
		return 0
	}
	limitTime := t[len(t)-1].Times - mouthTime
	var endPrice int
	for _, v := range t {
		if limitTime < v.Times {
			newPrice, _ := strconv.Atoi(v.TokenPrice)
			endPrice = newPrice
		}
	}
	price1, _ := strconv.Atoi(t[0].TokenPrice)
	a := float64(endPrice) / float64(price1)

	if a >= 1 {
		a = a*100 - 100
	} else {
		a = -(100 - a*100)
	}
	return a
}
