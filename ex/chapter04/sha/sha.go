package main

import (
    "os"
    "fmt"
    "log"
    "flag"
    "io/ioutil"
    "crypto/sha256"
    "crypto/sha512"
)

func sha(bytes []byte, alg string) (x []byte, err error) {

    switch (alg) {
        case "256":
            y := sha256.Sum256(bytes)
            x = y[:]
        case "384":
            y := sha512.Sum384(bytes)
            x = y[:]
        case "512":
            y := sha512.Sum512(bytes)
            x = y[:]
        default:
            return x, fmt.Errorf("unsupported sha algorithm %s\n", alg)
    }

    return x, nil

}

func main () {

    alg := flag.String("alg", "256", "")
    flag.Parse()

    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        log.Printf("can't read stdin: %v\n", err)
        return
    }

    x, err := sha(bytes, *alg)
    if err != nil {
        log.Printf("can't compute sha sum: %v\n", err)
        return
    }

    fmt.Printf("sha%s: %x\n", *alg, x)
}
