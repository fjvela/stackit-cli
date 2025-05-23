## stackit key-pair create

Creates a key pair

### Synopsis

Creates a key pair.

```
stackit key-pair create [flags]
```

### Examples

```
  Create a new key pair with public-key "ssh-rsa xxx"
  $ stackit key-pair create --public-key `ssh-rsa xxx`

  Create a new key pair with public-key from file "/Users/username/.ssh/id_rsa.pub"
  $ stackit key-pair create --public-key `@/Users/username/.ssh/id_rsa.pub`

  Create a new key pair with name "KEY_PAIR_NAME" and public-key "ssh-rsa yyy"
  $ stackit key-pair create --name KEY_PAIR_NAME --public-key `ssh-rsa yyy`

  Create a new key pair with public-key "ssh-rsa xxx" and labels "key=value,key1=value1"
  $ stackit key-pair create --public-key `ssh-rsa xxx` --labels key=value,key1=value1
```

### Options

```
  -h, --help                    Help for "stackit key-pair create"
      --labels stringToString   Labels are key-value string pairs which can be attached to a key pair. E.g. '--labels key1=value1,key2=value2,...' (default [])
      --name string             Key pair name
      --public-key string       Public key to be imported (format: ssh-rsa|ssh-ed25519)
```

### Options inherited from parent commands

```
  -y, --assume-yes             If set, skips all confirmation prompts
      --async                  If set, runs the command asynchronously
  -o, --output-format string   Output format, one of ["json" "pretty" "none" "yaml"]
  -p, --project-id string      Project ID
      --region string          Target region for region-specific requests
      --verbosity string       Verbosity of the CLI, one of ["debug" "info" "warning" "error"] (default "info")
```

### SEE ALSO

* [stackit key-pair](./stackit_key-pair.md)	 - Provides functionality for SSH key pairs

