## stackit network-area route delete

Deletes a static route in a STACKIT Network Area (SNA)

### Synopsis

Deletes a static route in a STACKIT Network Area (SNA).

```
stackit network-area route delete ROUTE_ID [flags]
```

### Examples

```
  Delete a static route with ID "xxx" in a STACKIT Network Area with ID "yyy" in organization with ID "zzz"
  $ stackit network-area route delete xxx --organization-id zzz --network-area-id yyy
```

### Options

```
  -h, --help                     Help for "stackit network-area route delete"
      --network-area-id string   STACKIT Network Area ID
      --organization-id string   Organization ID
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

* [stackit network-area route](./stackit_network-area_route.md)	 - Provides functionality for static routes in STACKIT Network Areas

