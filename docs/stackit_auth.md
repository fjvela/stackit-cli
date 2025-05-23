## stackit auth

Authenticates the STACKIT CLI

### Synopsis

Authenticates in the STACKIT CLI.

```
stackit auth [flags]
```

### Options

```
  -h, --help   Help for "stackit auth"
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

* [stackit](./stackit.md)	 - Manage STACKIT resources using the command line
* [stackit auth activate-service-account](./stackit_auth_activate-service-account.md)	 - Authenticates using a service account
* [stackit auth get-access-token](./stackit_auth_get-access-token.md)	 - Prints a short-lived access token.
* [stackit auth login](./stackit_auth_login.md)	 - Logs in to the STACKIT CLI
* [stackit auth logout](./stackit_auth_logout.md)	 - Logs the user account out of the STACKIT CLI

