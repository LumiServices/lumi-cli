# lumi-cli

A CLI for managing **lumi/lumina** services

## Setup

`lumi-cli` requires authentication credentials to operate.
You can create them interactively with:

```bash
lumi-cli config init
```

Or manually create the configuration file `~/.lumi-config/config.toml`

### Example `config.toml`

```toml
[default]
endpoint = "http://localhost:8000" # Replace with your Lumi instance URL
output = "xml"                     # Options: xml | json
verify_ssl = false
access_key = "lumiserver"
secret_access_key = "lumiserver"
```

## Usage

```bash
# List buckets
lumi-cli ls

# Upload a file
lumi-cli upload <object> <bucket>
```
