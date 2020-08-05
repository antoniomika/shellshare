# shellshare

A CLI tool to share command output with [seeshell](https://github.com/antoniomika/seeshell)

## Deploy

## CLI Flags

```text
The shellshare command

Usage:
  shellshare [flags]

Flags:
      --address string                The address to connect to (default "ssi.sh:1337")
      --command string                The command to execute (default "/bin/bash")
  -c, --config string                 Config file (default "config.yml")
      --data-directory string         Directory that holds data (default "deploy/data/")
      --debug                         Enable debugging information
      --delay duration                The delay to wait after printing the first message from the socket (default 4s)
  -h, --help                          help for shellshare
      --log-to-file                   Enable writing log output to file, specified by log-to-file-path
      --log-to-file-compress          Enable compressing log output files
      --log-to-file-max-age int       The maxium number of days to store log output in a file (default 28)
      --log-to-file-max-backups int   The maxium number of rotated logs files to keep (default 3)
      --log-to-file-max-size int      The maximum size of outputed log files in megabytes (default 500)
      --log-to-file-path string       The file to write log output to (default "/tmp/shellshare.log")
      --log-to-stdout                 Enable writing log output to stdout (default true)
      --remote                        Allow remote access to the command
      --time-format string            The time format to use for general log messages (default "2006/01/02 - 15:04:05")
  -v, --version                       version for shellshare
```
