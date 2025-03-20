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

// TODO: method to draw triplets from tileset

// TODO: method to draw sequences from tileset

// TODO: method to draw pair(s) from tileset

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
	correctOrder := []string{"1", "2", "3", "4", "5", "0", "6", "7", "8", "9"}
	sort.SliceStable(haipai, func(i, j int) bool {
		t1 := haipai[i]
		t2 := haipai[j]
		t1first := string(t1[0])
		t2first := string(t2[0])

		return slices.Index(correctOrder, t1first) < slices.Index(correctOrder, t2first)
	})
	// second sort by letter (suit)
	sort.SliceStable(haipai, func(i, j int) bool {
		t1 := haipai[i]
		t2 := haipai[j]
		t1second := t1[1]
		t2second := t2[1]

		return (t1second < t2second)
	})

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
