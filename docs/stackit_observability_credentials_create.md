## stackit observability credentials create

Creates credentials for an Observability instance.

### Synopsis

Creates credentials (username and password) for an Observability instance.
The credentials will be generated and included in the response. You won't be able to retrieve the password later.

```
stackit observability credentials create [flags]
```

### Examples

```
  Create credentials for Observability instance with ID "xxx"
  $ stackit observability credentials create --instance-id xxx
```

### Options

```
  -h, --help                 Help for "stackit observability credentials create"
      --instance-id string   Instance ID
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

* [stackit observability credentials](./stackit_observability_credentials.md)	 - Provides functionality for Observability credentials

