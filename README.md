# google

a command line client tool for interacting with google APIs

[admin-sdk](https://developers.google.com/admin-sdk/directory/v1/quickstart/go)

## 


## environment config

create a config file in your home directory `~/.config/.google.json`

```json
{
  "default": {
  },
}
```

for support of multiple configurations create additional top level elements in your config file and pass a matching `--env` flag with your commands

if no `--env` flag is passed the cli will load the `default` configuration

## load build harness

```sh
make init
```

## load dependencies

```sh
GO111MODULE=on go mod vendor
```
