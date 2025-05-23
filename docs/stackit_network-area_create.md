## stackit network-area create

Creates a STACKIT Network Area (SNA)

### Synopsis

Creates a STACKIT Network Area (SNA) in an organization.

```
stackit network-area create [flags]
```

### Examples

```
  Create a network area with name "network-area-1" in organization with ID "xxx" with network ranges and a transfer network
  $ stackit network-area create --name network-area-1 --organization-id xxx --network-ranges "1.1.1.0/24,192.123.1.0/24" --transfer-network "192.160.0.0/24"

  Create a network area with name "network-area-2" in organization with ID "xxx" with network ranges, transfer network and DNS name server
  $ stackit network-area create --name network-area-2 --organization-id xxx --network-ranges "1.1.1.0/24,192.123.1.0/24" --transfer-network "192.160.0.0/24" --dns-name-servers "1.1.1.1"

  Create a network area with name "network-area-3" in organization with ID "xxx" with network ranges, transfer network and additional options
  $ stackit network-area create --name network-area-3 --organization-id xxx --network-ranges "1.1.1.0/24,192.123.1.0/24" --transfer-network "192.160.0.0/24" --default-prefix-length 25 --max-prefix-length 29 --min-prefix-length 24

  Create a network area with name "network-area-1" in organization with ID "xxx" with network ranges and a transfer network and labels "key=value,key1=value1"
  $ stackit network-area create --name network-area-1 --organization-id xxx --network-ranges "1.1.1.0/24,192.123.1.0/24" --transfer-network "192.160.0.0/24" --labels key=value,key1=value1
```

### Options

```
      --default-prefix-length int   The default prefix length for networks in the network area
      --dns-name-servers strings    List of DNS name server IPs
  -h, --help                        Help for "stackit network-area create"
      --labels stringToString       Labels are key-value string pairs which can be attached to a network-area. E.g. '--labels key1=value1,key2=value2,...' (default [])
      --max-prefix-length int       The maximum prefix length for networks in the network area
      --min-prefix-length int       The minimum prefix length for networks in the network area
  -n, --name string                 Network area name
      --network-ranges strings      List of network ranges (default [])
      --organization-id string      Organization ID
      --transfer-network string     Transfer network in CIDR notation
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

* [stackit network-area](./stackit_network-area.md)	 - Provides functionality for STACKIT Network Area (SNA)

