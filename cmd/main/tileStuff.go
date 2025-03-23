package main

import (
	//"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
)

type Tileset struct {
	suits [][]string
	tiles map[string]int
}

func buildTileset() Tileset {
	suits := [][]string{
		{"0m", "1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m"},
		{"0p", "1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p"},
		{"0s", "1s", "2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s"},
		{"1z", "2z", "3z", "4z", "5z", "6z", "7z"},
	}
	tiles := map[string]int{}

	// build tile pool for drawing from
	for i := range suits {
		for _, t := range suits[i] {
			if strings.HasPrefix(t, "0") {
				tiles[t] = 1
			} else if strings.HasPrefix(t, "5") && t != "5z" {
				tiles[t] = 3
			} else {
				tiles[t] = 4
			}
		}
	}

	return Tileset{
		suits: suits,
		tiles: tiles,
	}
}

func (ts Tileset) drawTriplet() []string {
	redfiveroll := rand.IntN(4)
	suitrollnum := rand.IntN(4)
	suitroll := ts.suits[suitrollnum]

	tilerollnum := rand.IntN(len(suitroll))
	tileroll := suitroll[tilerollnum]
	redfive := suitroll[0]
	redfivecount := ts.tiles[redfive]

	tilecount := ts.tiles[tileroll]

	for tilecount < 3 {
		if (strings.HasPrefix(tileroll, "5") && tileroll != "5z") && tilecount == 2 && redfivecount > 0 {
			redfiveroll = 0
			break
		}
		tilerollnum = rand.IntN(len(suitroll))
		tileroll = suitroll[tilerollnum]
		tilecount = ts.tiles[tileroll]
	}

	triplet := []string{tileroll, tileroll, tileroll}

	if (strings.HasPrefix(tileroll, "5") && tileroll != "5z") && redfivecount > 0 && redfiveroll == 0 {
		ts.tiles[tileroll] -= 2
		ts.tiles[redfive]--
		triplet = []string{redfive, tileroll, tileroll}
	} else {
		ts.tiles[tileroll] -= 3
	}

	return triplet
}

func (ts Tileset) drawSequence() []string {
	var first string
	var second string
	var third string
	redfiveroll := rand.IntN(4)
	firstcount := 0
	secondcount := 0
	thirdcount := 0

	for firstcount == 0 || secondcount == 0 || thirdcount == 0 {
		suitrollnum := rand.IntN(3)
		suitroll := ts.suits[suitrollnum]
		tilerollnum := rand.IntN(7) + 2

		first = suitroll[tilerollnum-1]
		second = suitroll[tilerollnum]
		third = suitroll[tilerollnum+1]

		redfive := suitroll[0]
		redfivecount := ts.tiles[redfive]

		if tilerollnum == 5 && ((ts.tiles[second] < 1 && redfivecount > 0) || (redfivecount > 0 && redfiveroll == 0)) {
			second = suitroll[0]
		}
		if tilerollnum-1 == 5 && ((ts.tiles[first] < 1 && redfivecount > 0) || (redfivecount > 0 && redfiveroll == 0)) {
			first = suitroll[0]
		}
		if tilerollnum+1 == 5 && ((ts.tiles[third] < 1 && redfivecount > 0) || (redfivecount > 0 && redfiveroll == 0)) {
			third = suitroll[0]
		}

		firstcount = ts.tiles[first]
		secondcount = ts.tiles[second]
		thirdcount = ts.tiles[third]
	}
	ts.tiles[first]--
	ts.tiles[second]--
	ts.tiles[third]--

	sequence := []string{first, second, third}
	return sequence
}

func (ts Tileset) drawPair(isTanyao bool, isChiitoi bool, hand []string) []string {
	redfiveroll := rand.IntN(4)
	var redfive string
	var redfivecount int
	var tileroll string
	var suitrollnum int
	tilecount := 0

	for tilecount < 2 {
		if isTanyao {
			suitrollnum = rand.IntN(3)
		} else {
			suitrollnum = rand.IntN(4)
		}
		suitroll := ts.suits[suitrollnum]

		tilerollnum := rand.IntN(len(suitroll))
		if isTanyao && (tilerollnum < 2 || tilerollnum > 8) {
			tilerollnum = rand.IntN(7) + 2
		}

		tileroll = suitroll[tilerollnum]
		tilecount = ts.tiles[tileroll]

		// if we already have a pair of this tile,
		// and the hand is a chiitoi,
		// try rolling again (goes to next iteration of loop)
		if slices.Contains(hand, tileroll) && isChiitoi {
			tilecount = 0
			continue
		}

		redfive = suitroll[0]
		redfivecount = ts.tiles[redfive]

		if (strings.HasPrefix(tileroll, "5") && tileroll != "5z") && tilecount == 1 && redfivecount > 0 {
			redfiveroll = 0
			break
		}
	}
	pair := []string{tileroll, tileroll}

	if (strings.HasPrefix(tileroll, "5") && tileroll != "5z") && redfivecount > 0 && redfiveroll == 0 {
		ts.tiles[tileroll]--
		ts.tiles[redfive]--
		pair = []string{redfive, tileroll}
	} else {
		ts.tiles[tileroll] -= 2
	}

	return pair
}

func sortHand(hand []string) []string {
	correctOrder := []string{"1", "2", "3", "4", "5", "0", "6", "7", "8", "9"}
	sort.SliceStable(hand, func(i, j int) bool {
		t1 := hand[i]
		t2 := hand[j]
		t1first := string(t1[0])
		t2first := string(t2[0])

		return slices.Index(correctOrder, t1first) < slices.Index(correctOrder, t2first)
	})
	// second sort by letter (suit)
	sort.SliceStable(hand, func(i, j int) bool {
		t1 := hand[i]
		t2 := hand[j]
		t1second := t1[1]
		t2second := t2[1]

		return (t1second < t2second)
	})
	return hand
}

func (ts Tileset) generateTestHand() []string {
	hand := []string{}
	numseqs := rand.IntN(5)
	numtrips := 4 - numseqs
	var sequence []string
	var triplet []string

	for range numseqs {
		sequence = ts.drawSequence()
		hand = slices.Concat(hand, sequence)
	}
	for range numtrips {
		triplet = ts.drawTriplet()
		hand = slices.Concat(hand, triplet)
	}
	pair := ts.drawPair(false, false, hand)
	hand = slices.Concat(hand, pair)

	hand = sortHand(hand)
	return hand
}

func (ts Tileset) generateHaipai() []string {
	haipai := []string{}

	for range 14 {
		suitrollnum := rand.IntN(4)
		suitroll := ts.suits[suitrollnum]
		tilerollnum := rand.IntN(len(suitroll))
		tileroll := suitroll[tilerollnum]
		tile := ts.tiles[tileroll]

		for tile < 1 {
			tilerollnum = rand.IntN(len(suitroll))
			tileroll = suitroll[tilerollnum]
			tile = ts.tiles[tileroll]
		}
		ts.tiles[tileroll]--

		haipai = append(haipai, tileroll)
	}
	haipai = sortHand(haipai)

	return haipai
}

// type Winninghand struct {
// 	fullHand    [][]string
// 	openTiles   []string
// 	closedTiles []string
// 	yakuWon     []string
// 	winningTile []string
// 	waitType    string
// 	pairType    string
// 	triplets    int
// 	sequences   int
// 	han         int
// 	fu          int
// 	points      int
// }

// func initWinninghand() Winninghand {
// 	return Winninghand{
// 		fullHand:    [][]string{},
// 		openTiles:   []string{},
// 		closedTiles: []string{},
// 		yakuWon:     []string{},
// 		winningTile: []string{},
// 		waitType:    "any",
// 		pairType:    "any",
// 		triplets:    0,
// 		sequences:   0,
// 		han:         0,
// 		fu:          0,
// 		points:      0,
// 	}
// }

// func (h *Winninghand) populateHand() {
// 	h.fullHand = append(h.fullHand, h.openTiles, h.closedTiles, h.winningTile)
// }
