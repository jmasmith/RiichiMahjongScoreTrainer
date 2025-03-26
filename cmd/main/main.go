package main

/* RIICHI MAHJONG STUFF
SS doukou = SSK, SS doujun = SSJ
Yaku compatibility:
	riichi       -> any; closed (double riichi, no chankan)
	ippatsu      -> riichi; closed
	menzen tsumo -> NO chankan, NO houtei; closed
	tanyao       -> NO: itsuu, yakuhai, chanta, junchan, shosangen, honroutou; open/closed
	pinfu        -> NO: yakuhai (triplet or pair), SSK, toitoi, 3-ankou, 3-kantsu, shosangen, honroutou, chiitoi, rinshan; closed
	iipeikou     -> NO: SSK, toitoi, 3-ankou, 3-kantsu, 2-iipeikou, honroutou, chiitoi; closed
	ittsuu       -> NO: tanyao, SSJ/SSK, toitoi, 3-ankou, 3-kantsu, chanta, junchan, 2-iipeikou, shosangen, honroutou, chiitoi; open/closed
	yakuhai      -> NO: tanyao, pinfu, junchan, 2-iipeikou, chinitsu, chiitoi; open/closed
	SSJ          -> NO: ittsuu, SSK, toitoi, 3-ankou/kantsu, 2-iipeikou, shosangen, honroutou, honitsu, chinitsu, chiitoi; open/closed
	SSK          -> NO: pinfu, iipeikou, ittsuu, SSJ, 2-iipeikou, shosangen, honroutou, chinitsu, chiitoi; open/closed
	toitoi       -> NO: menzen tsumo, pinfu, iipeikou, ittsuu, SSJ, chanta, junchan, 2-iipeikou, chiitoi, chankan; open/closed
	3-ankou      -> NO: pinfu, iipeikou, ittsuu, SSJ, 2-iipeikou, chiitoi; open/closed (aside from the 3 triplets)
	3-kantsu     -> NO: pinfu, iipeikou, ittsuu, SSJ, 2-iipeikou, chiitoi; open/closed
	chanta       -> NO: tanyao, ittsuu, toitoi, junchan, honroutou, chinitsu, chiitoi; open/closed
	junchan      -> NO: tanyao, ittsuu, yakuhai, toitoi, chanta, shosangen, honroutou, honitsu, chiitoi; open/closed
	2-iipeikou   -> NO: iipeikou, ittsuu, yakuhai, SSJ, SSK, toitoi, 3-ankou/kantsu, shosangen, honroutou, chiitoi, rinshan, chankan; closed
	shosangen    -> NO: tanyao, pinfu, ittsuu, SSJ, SSK, junchan, 2-iipeikou, chanta, chiitoi; open/closed
	honroutou    -> NO: tanyao, pinfu, iipeikou, ittsuu, SSJ, chanta, junchan, 2-iipeikou, chinitsu, chankan; open/closed
	honitsu      -> NO: tanyao, SSJ, SSK, junchan, chinitsu; open/closed
	chinitsu     -> NO: yakuhai, SSJ, SSK, chanta, shosangen, honroutou, honitsu; open/closed
	chiitoi      -> NO: pinfu, iipeikou, ittsuu, yakuhai, SSJ/SSK, toitoi, 3-ankou/kantsu, chanta, junchan, 2-iipeikou, shosangen, rinshan, chankan; closed
	rinshan      -> NO: iipeikou, pinfu, 2-iipeikou, chiitoi, haitei, houtei, chankan; open/closed
	haitei       -> NO: rinshan, houtei, chankan; open/closed
	houtei       -> NO: ippatsu, menzen tsumo, rinshan, haitei, chankan; open/closed
	chankan      -> NO: double riichi, menzen tsumo, toitoi, 2-iipeikou, honroutou, chiitoi, rinshan, haitei, houtei; open/closed

	HONOR TILE SHORTHAND
		1z, 2z, 3z, 4z -> East, South, West, North
		5z, 6z, 7z -> White, Green, Red
*/

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	//"sort"
	//"strings"
)

type Hand struct {
	sequences    [][]string
	triplets     [][]string
	pairs        [][]string
	fullHand     []string
	closedGroups [][]string
	openGroups   [][]string
	isRiichi     bool
	agari        string
	isClosed     bool
	hanValue     int
	fuValue      int
	isIppatsu    bool
	isTanyao     bool
	isChiitoi    bool
	isPinfu      bool
	hasIipeikou  bool
	hasYakuhai   bool
	isToitoi     bool
	hasItsuu     bool
	hasSSJ       bool
	hasSSK       bool
	hasShosangen bool
	isHonroutou  bool
	hasSanankou  bool
	isHonitsu    bool
	flushSuit    int
	isChanta     bool
	isJunchan    bool
	isRyanpeikou bool
}

func newHand() Hand {
	return Hand{}
}

func (h *Hand) testGenerateHand(ts Tileset) {
	h.isTanyao = true
	h.isPinfu = true
	h.isClosed = true
	h.isChiitoi = false
	h.sequences = [][]string{}
	h.pairs = [][]string{}
	h.fullHand = []string{}
	h.closedGroups = [][]string{}
	var tilegroup []string

	//draw 4 sequences, make sure there are no dupes
OUTER:
	for len(h.sequences) < 4 {
		tilegroup = ts.drawSequence(*h)

		//check for dupes
		for _, group := range h.sequences {
			if reflect.DeepEqual(tilegroup, group) {
				ts.returnTiles(tilegroup)
				continue OUTER
			}
		}

		// append to sequences 2d slice
		h.sequences = append(h.sequences, tilegroup)
		h.closedGroups = append(h.closedGroups, tilegroup)
		h.fullHand = slices.Concat(h.fullHand, tilegroup)
	}

	tilegroup = ts.drawPair(*h)
	h.pairs = append(h.pairs, tilegroup)
	h.closedGroups = append(h.closedGroups, tilegroup)
	h.fullHand = slices.Concat(h.fullHand, tilegroup)

	h.fullHand = sortHand(h.fullHand)
}

func main() {
	tileset := buildTileset()
	newhand := newHand()

	newhand.testGenerateHand(tileset)

	fmt.Println("Tanyao pinfu, hopefully: ", newhand.fullHand)
	fmt.Println("after draw: ", tileset.tiles)

	// for range 7 {
	// 	pair = tileset.drawPair(false, true, hand)

	// 	hand = slices.Concat(hand, pair)
	// }
	// hand = sortHand(hand)

	// fmt.Println("Yanyao chiitoi: ", hand)

	roundRoll := rand.IntN(2)
	if roundRoll == 0 {
		fmt.Println("East round")
	} else {
		fmt.Println("South round")
	}

	seatRoll := rand.IntN(4)
	switch seatRoll {
	case 0:
		fmt.Println("East seat")
	case 1:
		fmt.Println("South seat")
	case 2:
		fmt.Println("West seat")
	default:
		fmt.Println("North seat")
	}

	roll1 := rand.IntN(11)
	switch {
	case roll1 >= 0 && roll1 < 5:
		fmt.Println("Normal hand")
		// roll for "normal hand" yaku
	case roll1 >= 5 && roll1 < 10:
		fmt.Println("7-pair-like hand")
		// roll for 7 pairs OR ryanpeikou
	default:
		fmt.Println("Yakuman")
		// roll for which yakuman, find yakuman compatibility
	}
}
