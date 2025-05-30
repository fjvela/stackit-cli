## stackit server public-ip attach

Attaches a public IP to a server

### Synopsis

Attaches a public IP to a server.

```
stackit server public-ip attach PUBLIC_IP_ID [flags]
```

### Examples

```
  Attach a public IP with ID "xxx" to a server with ID "yyy"
  $ stackit server public-ip attach xxx --server-id yyy
```

### Options

```
  -h, --help               Help for "stackit server public-ip attach"
      --server-id string   Server ID
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

* [stackit server public-ip](./stackit_server_public-ip.md)	 - Allows attaching/detaching public IPs to servers

