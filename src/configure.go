package main
import (
    // "fmt"
    // toml "github.com/pelletier/go-toml"
)
type Pruning struct {
    pruning_preset string
    pruning string
    pruning_recent int
    pruning_every int
    pruning_interval int
}
type InstallPlan struct {
    install_preset string
    chain string
    und_url string
    cosmovisor_url string
    moniker string
    Pruning
}
// func (i InstallPlan) setPruning( p Pruning){
    // i.Pruning = p
// }
