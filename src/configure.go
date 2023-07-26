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
type StateSync struct {
    enabled bool
    height int
    hash string
}
type InstallPlan struct {
    install_preset string
    chain string
    und_url string
    cosmovisor_url string
    und_version string
    cosmovisor_version string
    moniker string
    Pruning
    StateSync
}
// func (i InstallPlan) setPruning( p Pruning){
    // i.Pruning = p
// }
