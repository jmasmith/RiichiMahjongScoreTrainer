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
	//"slices"
	//"sort"
	//"strings"
)

func main() {
	tileset2 := buildTileset()
	fmt.Println("before draw: ", tileset2.tiles)
	myHand2 := tileset2.generateHaipai()
	fmt.Println("after draw: ", tileset2.tiles)
	fmt.Println("My hand: ", myHand2)

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
