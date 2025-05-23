## stackit security-group delete

Deletes a security group

### Synopsis

Deletes a security group by its internal ID.

```
stackit security-group delete GROUP_ID [flags]
```

### Examples

```
  Delete a named group with ID "xxx"
  $ stackit security-group delete xxx
```

### Options

```
  -h, --help   Help for "stackit security-group delete"
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

* [stackit security-group](./stackit_security-group.md)	 - Manage security groups

