package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

/* Each rune has its own mutex. */
type LockMap map[rune]*sync.Mutex

var runeLocks LockMap = LockMap{}

/* Each rune's lock must be initialized so the map itself has a single lock. */
var runeLocksMapMutex sync.Mutex

/* Lock each update to a FreqMap key. So there is a parallel key-to-lock map. On
the first time seeing a key add a lock to the key then lock to update.
On subsequent times seeing the key only use the lock linked with the key.
Unlock FreqMap lock after update completes. */
func Frequency(wg *sync.WaitGroup, s string, rf FreqMap) {
	defer wg.Done()

	for _, r := range s {
		runeLocksMapMutex.Lock()
		if _, keyFound := runeLocks[r]; !keyFound {
			runeLocks[r] = &sync.Mutex{}
		}

		runeLocks[r].Lock()
		rf[r]++
		runeLocks[r].Unlock()
		runeLocksMapMutex.Unlock()
	}
}

func ConcurrentFrequency(s string) FreqMap {
	/* Start two frequency one at the beginning and one at the half-way mark. */
	halfIndex := int(len([]rune(s)) / 2)
	firstHalf, secondHalf := s[:halfIndex], s[halfIndex:]
	halves := []string{firstHalf, secondHalf}

	var runeFreq FreqMap = FreqMap{}

	/* Wait for both goroutines to finish using a WaitGroup. */
	var frequencyGroup sync.WaitGroup

	for _, half := range halves {
		frequencyGroup.Add(1)
		go Frequency(&frequencyGroup, half, runeFreq)
	}

	frequencyGroup.Wait()

	return runeFreq
}
