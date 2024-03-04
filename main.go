package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var allClocksAmount, interval, allRingAmountForWeakUp int

	_, err := fmt.Scanf("%d %d %d\n", &allClocksAmount, &interval, &allRingAmountForWeakUp)
	if err != nil {
		fmt.Println(err)
	}

	ringTimes := make([]int, allClocksAmount)

	for i := 0; i < allClocksAmount; i++ {
		_, err = fmt.Scan(&ringTimes[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	qsort(ringTimes)

	allRingTimesByPeriodsStorage := make([][]int, allClocksAmount)

	var wg sync.WaitGroup

	for i, _ := range ringTimes {
		ringCount := 0
		wg.Add(1)
		go func(ii int) {
			lastRings := make(map[int]int)
			var ringTimesByPeriods []int
			defer wg.Done()
			for ringCount < allClocksAmount {
				ringTime := ringTimes[ii]
				_, ok := lastRings[ringTime]
				if !ok {
					lastRings[ringTime] = 1
					ringTimesByPeriods = append(ringTimesByPeriods, ringTime)
					ringCount++
				}

				ringTimes[ii] = ringTime + interval
			}

			allRingTimesByPeriodsStorage[ii] = ringTimesByPeriods
		}(i)
	}

	wg.Wait()

	var allRingTimesByPeriods []int

	for _, elems := range allRingTimesByPeriodsStorage {
		allRingTimesByPeriods = append(allRingTimesByPeriods, elems...)
	}

	allRingTimesByPeriods = distinct(allRingTimesByPeriods)
	qsort(allRingTimesByPeriods)
	fmt.Println(allRingTimesByPeriods[allRingAmountForWeakUp-1])
}

func distinct(a []int) []int {
	m := make(map[int]int)
	var sl []int

	for _, el := range a {
		_, ok := m[el]
		if !ok {
			m[el] = 1
			sl = append(sl, el)
		}
	}

	return sl
}
func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	// Pick a pivot
	pivotIndex := rand.Int() % len(a)
	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]
	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}
