package core

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	N = 6 // number of participant
	K = 5 // sample size 1 < K < N
	A = 3 // quorum size 1 < a < K
	B = 2 // decision threshold >= 1
)

func Decide(block Block, peers []uint16) bool {
	fmt.Println("Process decide block")
	consecutiveSuccesses := 0
	preference := true
	for {
		// choose peer to preference
		preferPeers := randomCandidate(K)
		fmt.Println(preferPeers)

		type MaxSamePreference struct {
			Amount     int
			Preference bool
		}

		maxSamePreference := MaxSamePreference{}
		amountTrue := 0
		amountFalse := 0

		data, _ := json.Marshal(block)
		for _, position := range preferPeers {
			address := ":" + strconv.Itoa(int(peers[position]))
			if address == Node.Addr() {
				continue
			}
			response, err := Node.Request(context.TODO(), address, data)
			fmt.Println("response from "+address+" : ", string(response))
			if err != nil {
				fmt.Println(err)
				continue
			}
			if string(response) == "1" {
				amountTrue++
			} else {
				amountFalse++
			}
		}
		if amountTrue > amountFalse {
			maxSamePreference.Amount = amountTrue
			maxSamePreference.Preference = true
		} else {
			maxSamePreference.Amount = amountFalse
			maxSamePreference.Preference = false
		}
		fmt.Println("max same preference", maxSamePreference)
		if maxSamePreference.Amount >= A {
			newPreference := maxSamePreference.Preference
			if newPreference == preference {
				consecutiveSuccesses++
			} else {
				consecutiveSuccesses = 1
				preference = newPreference
			}
		} else {
			consecutiveSuccesses = 0
		}
		fmt.Println("consecutive success :", consecutiveSuccesses)
		if consecutiveSuccesses >= B {
			return preference
		}
	}
}

func randomCandidate(k int) []int {
	result := []int{}
	for i := 0; i < k; i++ {
	TRY:
		candidate := rand.Intn(N)
		if checkInArray(candidate, result) {
			goto TRY
		} else {
			result = append(result, candidate)
		}
	}
	return result
}

func checkInArray(x int, arr []int) bool {
	for _, i := range arr {
		if i == x {
			return true
		}
	}
	return false
}
