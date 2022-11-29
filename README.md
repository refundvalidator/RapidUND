# RapidUND
#### Bash script to rapidly deploy an UND Node with StateSync/Cosmovisor as Daemon

Currently, the script is set to use UND v1.6.3 and COSMOVISOR v1.2.0

No upgrade is currently in place for the next UND version within the COSMOVISOR Files




## *!!!THIS SCRIPT WILL DESTROY THE FOLLOWING FILES IF PRESENT!!!*

-$HOME/.und_mainchain

-$HOME/temp

-$HOME/UNDBackup

-/usr/local/bin/und

-/usr/local/bin/cosmovisor

-/etc/systemd/system/und.service



## *PREREQUISITES:*

-jq

-wget

-curl


## *USAGE:*

Simply run "sh /path/to/undsetup.sh"

If the node is stuck endlessly on "Dialing peer address" then run the script once again.

If you have SELinux installed it may block the und.service file for the daemon, only workaround I know of is to disable SELinux in /etc/selinux/config

If you are importing a current Validator, stop both this node, and your original node, then replace the node_key.json and priv_validator_key.json in $HOME/.und_mainchain/config with your original keys, then start this new node again. If both nodes are running at the same time with the same priv_validator_key.json you will be jailed.

Script will copy node_key.json and priv_validator_key.json into $HOME/UNDBackup

Node will likely fail if running a node on the current network

## *DETAILS:*

The node deployed will be using default configuration given by the Unification Docs at https://docs.unification.io/

Pruning remains as default, this node will not act as an archive node

This script is meant to be used by experienced operators, used to quickly deploy a node in a time of need or when migrating to another machine, this is not recommended if you have not yet set up a node for yourself.



