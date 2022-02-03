package pkg

import (
	"math/rand"
	"time"
)

//Shuffle Document
/*
 *----------------------------------------
 * @brief    Shuffle
 *           stringSlice ShuffleFunc
 *
 * @Pram[in] []string
 *
 * @return
 *
 * @rules
 *----------------------------------------
 */
func Shuffle(data []string) {
	n := len(data)
	rand.Seed(time.Now().Unix())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
