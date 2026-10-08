package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/jsonpretty"
	"github.com/cli/go-gh/v2/pkg/term"
)

type sshKey struct {
	ID        int        `json:"id"`
	Key       string     `json:"key"`
	Title     string     `json:"title,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	LastUsed  *time.Time `json:"last_used,omitempty"`
	Type      string     `json:"type"`
}

func parseArgs() (string, bool, error) {
	useJSON := flag.Bool("json", false, "Print output in JSON format")
	flag.Parse()

	username := ""
	arguments := flag.Args()
	if len(arguments) > 1 {
		return "", false, fmt.Errorf("too many arguments, expected at most one")
	}
	if len(arguments) == 1 {
		username = arguments[0]
	}

	return username, *useJSON, nil
}

func stringRepresentation(signingKeys []sshKey, authKeys []sshKey) (sshKeys []string) {
	for _, signingKey := range signingKeys {
		comment := ""
		if signingKey.Title != "" {
			comment = fmt.Sprintf(" %s", signingKey.Title)
		}
		sshKeys = append(sshKeys, fmt.Sprintf("%s%s", signingKey.Key, comment))
	}
out:
	for _, authKey := range authKeys {
		for _, signingKey := range signingKeys {
			if authKey.Key == signingKey.Key {
				continue out
			}
		}
		sshKeys = append(sshKeys, authKey.Key)
	}
	return sshKeys
}

func main() {
	username, useJSON, err := parseArgs()
	if err != nil {
		log.Fatalf("Failed to parse arguments: %s", err)
	}

	client, err := api.DefaultRESTClient()
	if err != nil {
		log.Fatalf("Failed to create REST client: %s", err)
	}

	if username == "" {
		login := struct{ Login string }{}
		err = client.Get("user", &login)
		if err != nil {
			fmt.Println(err)
			return
		}
		username = login.Login
	}

	var authKeys []sshKey
	err = client.Get(fmt.Sprintf("users/%s/keys", username), &authKeys)
	if err != nil {
		log.Fatalf("Failed to get authentication keys for user %s: %s", username, err)
	}

	var signingKeys []sshKey
	err = client.Get(fmt.Sprintf("users/%s/ssh_signing_keys", username), &signingKeys)
	if err != nil {
		log.Fatalf("Failed to get signing keys for user %s: %s", username, err)
	}

	if useJSON {
		for i := range authKeys {
			authKeys[i].Type = "authentication"
		}
		for i := range signingKeys {
			signingKeys[i].Type = "signing"
		}

		buf, err := json.Marshal(slices.Concat(authKeys, signingKeys))
		if err != nil {
			log.Fatalf("Failed to marshal SSH keys: %s", err)
		}
		if err = jsonpretty.Format(os.Stdout, bytes.NewReader(buf), "  ", term.IsTerminal(os.Stdout)); err != nil {
			log.Fatalf("Failed to format JSON: %s", err)
		}
	} else {
		for _, k := range stringRepresentation(signingKeys, authKeys) {
			fmt.Println(k)
		}
	}
}
