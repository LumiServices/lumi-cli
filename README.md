# lumi-cli
A CLI for managing lumi/lumina services

## Setup
lumi-cli requires authentication credentials to operate. Create them using:
```bash
lumi-cli config init
```
Alternatively, manually create the configuration files:
- `~/.lumi-config/config.toml`
- `~/.lumi-config/credentials.toml`

### Example config.toml
```toml
[default]
endpoint = "http://localhost:8000" #change this to your lumi instance
output = "xml" # xml | json
verify_ssl = false
```

### Example credentials.toml
```toml
[default]
access_key_id = "lumiserver"
secret_access_key = "lumiserver"
```

## Usage
```bash
# List buckets
lumi-cli ls

# Upload a file
lumi-cli upload <object> <bucket>
```