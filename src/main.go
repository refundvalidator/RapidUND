package main

import (
    // "os"
    "fmt"
)

func main(){
    fmt.Println("Hello World!")
    getBinary("und", "./", "https://github.com/unification-com/mainchain/releases/download/v1.7.0/und_v1.7.0_linux_x86_64.tar.gz")
    getBinary("cosmovisor", "./", "https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.2.0/cosmovisor-v1.2.0-linux-amd64.tar.gz")
}
