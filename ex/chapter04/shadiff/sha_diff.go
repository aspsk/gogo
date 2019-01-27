package main

import (
    "fmt"
    "math/rand"
    "math/bits"
    "crypto/sha256"
    "strconv"
)

func diffBits(x1, x2 [sha256.Size]byte) (n int) {

    for i, _ := range x1 {
        n += bits.OnesCount8(x1[i] ^ x2[i])
    }

    return n
}

func main () {

    const N = 10 * 1024 * 1024
    var sum float64

    for i := 0; i < N; i++ {
        x1 := sha256.Sum256([]byte(strconv.Itoa(rand.Int())))
        x2 := sha256.Sum256([]byte(strconv.Itoa(rand.Int())))
        sum += float64(diffBits(x1, x2))
    }

    fmt.Println(sum / N)

}
