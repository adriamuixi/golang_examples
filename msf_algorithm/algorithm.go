package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Seller struct {
	SellerAccountId *int     `json:"seller_account_id"`
	Locales         []string `json:"locales"`
}

type MockData struct {
	Seller          []Seller
	SellerAccountId int
	Locale          string
}

func main() {
	mocksData := createMocks(15)
	for i, v := range mocksData {
		fmt.Printf("SCENARIO %d: \n\n%+v\n\n", i+1, v)
		start1 := time.Now()
		sol1 := selectBestSellerOption1(v.Seller, v.SellerAccountId, v.Locale)
		elapsed1 := time.Since(start1)
		log.Printf("----- Time spend: alg 1 with scenario %d took %s ,solution sellerID %d \n", i+1, elapsed1, sol1)
		start2 := time.Now()
		sol2 := selectBestSellerOption2(v.Seller, v.SellerAccountId, v.Locale)
		elapsed2 := time.Since(start2)
		log.Printf("----- Time spend: alg 2 with scenario %d took %s ,solution sellerID %d \n", i+1, elapsed2, sol2)
		fmt.Println("\n###########################\n\n")
	}
}

func createMocks(number int) []MockData {
	var result []MockData

	var locales = []string{
		"en_GB",
		"fr_FR",
		"ge_GE",
		"es_ES",
		"it_IT",
	}

	var sellers [][]Seller

	for i := 0; i < number; i++ {
		randomNumberSellers := rand.Intn(35-1) + 1
		var sellerSlice []Seller
		for z := 0; z < randomNumberSellers; z++ {
			var localesRandom []string
			for j := 0; j < rand.Intn(len(locales)-1)+1; j++ {
				loc := locales[rand.Intn(len(locales))]
				localesRandom = append(localesRandom, loc)
			}
			randomSellerId := rand.Intn(1000-1) + 1
			seller := Seller{
				&randomSellerId,
				localesRandom,
			}
			sellerSlice = append(sellerSlice, seller)
		}
		sellers = append(sellers, sellerSlice)
	}

	for i := 0; i < number; i++ {
		randomIndex := rand.Intn(len(locales)-1) + 1
		m := MockData{
			sellers[i],
			i,
			locales[randomIndex],
		}
		result = append(result, m)
	}
	return result
}

func selectBestSellerOption1(sellers []Seller, suggestedSellerAccountId int, suggestedLocale string) int {
	hasFoundSuggestedSeller := false                  // also means Solution One Is Dead
	hasFoundSuggestedLocale := false                  // also means Solution Two Has Been Found
	memorizedAccountId := *sellers[0].SellerAccountId // Sol 4 memorized
	for _, seller := range sellers {
		if !hasFoundSuggestedSeller {
			if *seller.SellerAccountId == suggestedSellerAccountId {
				if hasSuggestedLocale(seller, suggestedLocale) {
					return *seller.SellerAccountId // Sol 1 good seller good locale
				}
				if !hasFoundSuggestedLocale {
					memorizedAccountId = *seller.SellerAccountId // Sol 3 memorized as we're not sure sol 2 exists
				}
				hasFoundSuggestedSeller = true
			} else {
				if !hasFoundSuggestedLocale && hasSuggestedLocale(seller, suggestedLocale) {
					memorizedAccountId = *seller.SellerAccountId // Sol 2 memorized as we're not sure sol 1 exists
					hasFoundSuggestedLocale = true
				}
			}
		} else { // Sol 1 is dead
			if hasFoundSuggestedLocale {
				return memorizedAccountId // Sol 2 previously found
			}
			if hasSuggestedLocale(seller, suggestedLocale) {
				return *seller.SellerAccountId // Sol 2
			}
		}
	}
	return memorizedAccountId // Sol 4, 3 or 2 that is memorized
}

func selectBestSellerOption2(sellers []Seller, suggestedSellerAccountId int, suggestedLocale string) int {
	// 1 : Same seller, same language => ask other pages
	// 2 : Same language
	// 3 : Same seller
	// 4 : First found
	hasFoundSuggestedSellerBis := false
	for _, seller := range sellers {
		if *seller.SellerAccountId == suggestedSellerAccountId {
			if hasSuggestedLocale(seller, suggestedLocale) {
				return *seller.SellerAccountId // Sol 1 good seller good locale
			}
			hasFoundSuggestedSellerBis = true
		}
	}
	for _, seller := range sellers {
		if hasSuggestedLocale(seller, suggestedLocale) {
			return *seller.SellerAccountId // Sol 2 good locale
		}
	}
	if hasFoundSuggestedSellerBis {
		return suggestedSellerAccountId // Sol 3 good seller
	}
	return *sellers[0].SellerAccountId // Sol 4 first choice
}

func hasSuggestedLocale(seller Seller, suggestedLocale string) bool {
	for _, locale := range seller.Locales {
		if locale == suggestedLocale {
			return true
		}
	}
	return false
}
