# gh-get-ssh-keys

GitHub CLI extension to get a given user's authentication and signing SSH keys.

## Usage

```shell
gh get-ssh-keys $username
```

This will return a JSON blob containing two objects:
* `auth`, a list of authentication keys in the form returned by [_List public keys for a user_](https://docs.github.com/en/rest/users/keys?apiVersion=2026-03-10#list-public-keys-for-a-user),
* `signing`, a list of signing keys in the form returned by [_List SSH signing keys for a user_](https://docs.github.com/en/rest/users/ssh-signing-keys?apiVersion=2026-03-10#list-ssh-signing-keys-for-a-user).

For example,

```shell
gh get-ssh-keys fionn | jq -r ".[][].key" | sort -u
```
would get all keys for user `fionn`.
