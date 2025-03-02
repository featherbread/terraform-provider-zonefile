package main

//go:generate go tool tfplugindocs generate -provider-name zonefile

//go:generate rm -rf ./release/sources
//go:generate go tool go-licenses save . --save_path=./release/sources --ignore=github.com/ahamlinman/terraform-provider-zonefile
