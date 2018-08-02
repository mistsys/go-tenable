# Go-Tenable
Wrapper/library for interacting with Tenable's API. In early stages, but the idea is to automatically create JIRA tickets, and track issues.

# Configuration
API keys can be passed on the command line, or via a file `~/.config/tenable/config.yml`. The format is
```
# ~/.config/tenable/config.yml
accesskey: ${TENABLE_ACCESS_KEY}
secretkey: ${TENABLE_SECRET_KEY}
```

# Usage
```
$ tenable
A CLI for the Tenable API

Usage:
  tenable [command]

Available Commands:
  editor      Use the Tenable editor API
  folders     Use the Tenable folders API
  help        Help about any command
  scanners    Use the Tenable scanners API
  scans       Use the Tenable scans API
  server      Use the Tenable server API
  workbenches Use the Tenable workbenches API

Flags:
  -k, --accesskey string     Tenable Access Key (required)
  -f, --configFile string    Config file to read from
      --debug                Run in debug mode (dump raw request bodies)
  -h, --help                 help for tenable
  -o, --output-file -        Output file. Passing - writes to stdout (default "-")
      --params string        Query parameters given as a string of "key=value,key=value,..."
  -s, --secretkey string     Tenable Secret Key (required)
  -v, --verbose              Verbose output

Use "tenable [command] --help" for more information about a command.
```

Export results of a scan to JIRA tickets with `tenable scans export $scan_id --format jira -o tickets.csv`.

## Contributing
Add the `pre-push-hook` to your `.git/hooks/pre-push` file with `cp pre-push-hook .git/hooks/pre-push`


## Other similar clients
https://github.com/attwad/nessie
