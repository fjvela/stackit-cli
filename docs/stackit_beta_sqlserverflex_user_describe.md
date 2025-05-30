## stackit beta sqlserverflex user describe

Shows details of a SQLServer Flex user

### Synopsis

Shows details of a SQLServer Flex user.
The user password is only visible upon creation. You can reset it by running:
  $ stackit beta sqlserverflex user reset-password USER_ID --instance-id INSTANCE_ID

```
stackit beta sqlserverflex user describe USER_ID [flags]
```

### Examples

```
  Get details of a SQLServer Flex user with ID "xxx" of instance with ID "yyy"
  $ stackit beta sqlserverflex user describe xxx --instance-id yyy

  Get details of a SQLServer Flex user with ID "xxx" of instance with ID "yyy" in JSON format
  $ stackit beta sqlserverflex user describe xxx --instance-id yyy --output-format json
```

### Options

```
  -h, --help                 Help for "stackit beta sqlserverflex user describe"
      --instance-id string   ID of the instance
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

* [stackit beta sqlserverflex user](./stackit_beta_sqlserverflex_user.md)	 - Provides functionality for SQLServer Flex users

