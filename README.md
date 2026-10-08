# gh-get-ssh-keys

GitHub CLI extension to get a given user's authentication and signing SSH keys.

## Installation

```shell
gh extension install fionn/gh-get-ssh-keys
```

## Usage

```
gh get-ssh-keys [-json] [username]
```

If `username` is omitted, we default to the current user.

This will print unique SSH keys for the user in standard authorised keys format.

If passed `-json`, this will print a list of SSH key objects with fields `id`, `key` `created_at` and `type`. `type` will be either `"authentication"` or `"signing"`, If the key is an authentication key, we also include `last_used`. If the key is a signing key, we also include `title`.


## Examples

### Authorised Keys Format

```console
$ gh get-ssh-keys fionn
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSydtI14Kok5n/hdqTvuGkZWQhB5BcqIN6kKqxr0I2d xyza
sk-ssh-ed25519@openssh.com AAAAGnNrLXNzaC1lZDI1NTE5QG9wZW5zc2guY29tAAAAIOZ84JMwAYOcbxXLg2gmbREeMAjet9iDHrUITVeOkVTnAAAABHNzaDo= 17244581
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILbkp0LwqqV/w6wAGV9bwiR6FpHC/5DtiBAKFLZxvaSp lotus
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMhHsaAGLMPDuuzy/RO4CI3QQ2rnb1/3Q+8ytg+7KDBk
```

### JSON Format

#### Get First Authentication Key

```console
$ gh get-ssh-keys -json fionn | jq "[.[] | select(.type == \"authentication\")][0]"
{
  "id": 17783706,
  "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSydtI14Kok5n/hdqTvuGkZWQhB5BcqIN6kKqxr0I2d",
  "created_at": "2016-06-23T15:50:06Z",
  "last_used": "2025-11-23T07:50:18Z",
  "type": "authentication"
}
```

#### Get First Signing Key

```console
$ gh get-ssh-keys -json fionn | jq "[.[] | select(.type == \"signing\")][0]"
{
  "id": 2581,
  "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBSydtI14Kok5n/hdqTvuGkZWQhB5BcqIN6kKqxr0I2d",
  "title": "xyza",
  "created_at": "2022-08-26T23:52:13.946+08:00",
  "type": "signing"
}
```
