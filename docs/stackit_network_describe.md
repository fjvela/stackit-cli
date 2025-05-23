## stackit network describe

Shows details of a network

### Synopsis

Shows details of a network.

```
stackit network describe NETWORK_ID [flags]
```

### Examples

```
  Show details of a network with ID "xxx"
  $ stackit network describe xxx

  Show details of a network with ID "xxx" in JSON format
  $ stackit network describe xxx --output-format json
```

### Options

```
  -h, --help   Help for "stackit network describe"
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

* [stackit network](./stackit_network.md)	 - Provides functionality for networks

