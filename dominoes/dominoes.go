package dominoes

type Domino []int

func removeDomino(i int, doms []Domino) []Domino {
	domsCopy := make([]Domino, len(doms))
	copy(domsCopy, doms)

	domsCopy[i] = domsCopy[len(domsCopy)-1]
	domsCopy = domsCopy[:len(domsCopy)-1]
	return domsCopy
}

func twoRemainingFit(toFindLeft, toFindRight int, remainingDomA, remainingDomB Domino) bool {
	fitsLeft := toFindLeft == remainingDomA[0]
	middleMatch := remainingDomA[1] == remainingDomB[0]
	fitsRight := remainingDomB[1] == toFindRight
	return fitsLeft && middleMatch && fitsRight
}

func reverseDomino(d Domino) Domino {
	return Domino{d[1], d[0]}
}

func permuteDominos(domA, domB Domino) [][]Domino {
	return [][]Domino{{domA, domB}, {domA, reverseDomino(domB)}, {reverseDomino(domA), domB},
		{reverseDomino(domA), reverseDomino(domB)}}
}

func build(doms []Domino, chain []Domino, buildSpot int) ([]Domino, bool) {
	if buildSpot == 0 {
		valid := false
		// [ | ] ... [ | ]
		for i, domA := range doms {
			domsRemovedA := removeDomino(i, doms)
			for j, domB := range domsRemovedA {
				for _, domPair := range permuteDominos(domA, domB) {
					first, second := domPair[0], domPair[1]
					if first[0] == second[1] {
						chain[0], chain[len(chain)-1] = first, second
						domsRemovedAB := removeDomino(j, domsRemovedA)
						retChain, ok := build(domsRemovedAB, chain, buildSpot+1)
						valid = valid || ok
						if valid {
							return retChain, true
						}
					}
				}
			}
		}
	} else if buildSpot != int((len(chain)-1)/2) { // not done building
		// ... [ |l][A| ] ... [ |B][r| ]...
		valid := false
		prevDomA, prevDomB := chain[buildSpot-1], chain[len(chain)-buildSpot]
		toFindLeft, toFindRight := prevDomA[1], prevDomB[0]
		for i, domA := range doms {
			domsRemovedA := removeDomino(i, doms)
			for j, domB := range domsRemovedA {
				for _, domPair := range permuteDominos(domA, domB) {
					first, second := domPair[0], domPair[1]
					if first[0] == toFindLeft && second[1] == toFindRight {
						chain[buildSpot], chain[len(chain)-1-buildSpot] = first, second
						domsRemovedAB := removeDomino(j, domsRemovedA)
						retChain, ok := build(domsRemovedAB, chain, buildSpot+1)
						valid = valid || ok
						if valid {
							return retChain, true
						}
					}
				}
			}
		}
	} else { // final spot(s)
		// ... [ |l] X [r| ] ... or
		// ... [ |l] X X [r| ] ...
		prevDomA, prevDomB := chain[buildSpot-1], chain[len(chain)-buildSpot]
		toFindLeft, toFindRight := prevDomA[1], prevDomB[0]
		if len(doms) == 1 {
			remainingDom := doms[0]
			if toFindLeft == remainingDom[0] && remainingDom[1] == toFindRight {
				chain[buildSpot] = remainingDom
				return chain, true
			}
		} else {
			for i, remainingDomA := range doms {
				domsRemovedA := removeDomino(i, doms)
				for _, remainingDomB := range domsRemovedA {
					for _, domPair := range permuteDominos(remainingDomA, remainingDomB) {
						first, second := domPair[0], domPair[1]
						if twoRemainingFit(toFindLeft, toFindRight, first, second) {
							chain[buildSpot], chain[len(chain)-1-buildSpot] = first, second
							return chain, true

						}
					}
				}
			}
		}
	}

	return nil, false // could not find matching tail ends or could not find valid intermediates
}

// MakeChain creates a valid chain of dominos where adjacent values match allowing for reversal of
// dominos if and only if the input contains a valid chain.
func MakeChain(input []Domino) ([]Domino, bool) {
	// if empty is always true
	// if only one then both sides must be the same
	if len(input) == 0 {
		return input, true
	}

	if len(input) == 1 {
		singleton := input[0]
		if singleton[0] == singleton[1] {
			return input, true
		}
	}

	chain := make([]Domino, len(input))

	return build(input, chain, 0)
}
