package main

import (
    // "os"
    "fmt"
)

type JSONInfo struct {
    file_name string
    path string
    url string
}
type BinaryInfo struct {
    binary_name string
    path string
    url string
}

func main(){
    fmt.Println("Hello World!")
    test := BinaryInfo{
        binary_name : "und",
        path: "../tmp/und/",
        // path: "./",
        url : "https://github.com/unification-com/mainchain/releases/download/v1.7.0/und_v1.7.0_linux_x86_64.tar.gz",
    }
    fetchBinary(test)
    // getBinary("und", "../tmp/und/", "https://github.com/unification-com/mainchain/releases/download/v1.7.0/und_v1.7.0_linux_x86_64.tar.gz")
    // getBinary("cosmovisor", "../tmp/cosmovisor/", "https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.2.0/cosmovisor-v1.2.0-linux-amd64.tar.gz")
}
