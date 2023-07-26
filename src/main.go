package main

import (
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
    und_test := BinaryInfo{
        binary_name : "und",
        path: "../tmp/und/",
        url : "https://github.com/unification-com/mainchain/releases/download/v1.7.0/und_v1.7.0_linux_x86_64.tar.gz",
    }
    // cosmovisor_test := BinaryInfo{
    //     binary_name : "cosmovisor",
    //     path: "../tmp/cosmovisor/",
    //     url: "https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.2.0/cosmovisor-v1.2.0-linux-amd64.tar.gz",
    // }
    genesis_test := JSONInfo{
        file_name: "genesis.json",
        path: "../tmp/genesis/",
        url: "https://raw.githubusercontent.com/unification-com/mainnet/master/latest/genesis.json",
    }
    fetchGenesis(genesis_test)
    fetchBinary(und_test)
    // fetchBinary(cosmovisor_test)
    cleanUp("../tmp/")

    // getBinary("und", "../tmp/und/", "https://github.com/unification-com/mainchain/releases/download/v1.7.0/und_v1.7.0_linux_x86_64.tar.gz")
    // getBinary("cosmovisor", "../tmp/cosmovisor/", "https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.2.0/cosmovisor-v1.2.0-linux-amd64.tar.gz")
}
