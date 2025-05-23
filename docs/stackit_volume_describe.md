## stackit volume describe

Shows details of a volume

### Synopsis

Shows details of a volume.

```
stackit volume describe VOLUME_ID [flags]
```

### Examples

```
  Show details of a volume with ID "xxx"
  $ stackit volume describe xxx

  Show details of a volume with ID "xxx" in JSON format
  $ stackit volume describe xxx --output-format json
```

### Options

```
  -h, --help   Help for "stackit volume describe"
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

* [stackit volume](./stackit_volume.md)	 - Provides functionality for volumes

