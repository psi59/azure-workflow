# Azure Portal Alfred Workflow

An Alfred workflow for quick access to Azure Portal services. Search and open Azure services directly from Alfred with fuzzy search support.

## Features

- **Quick Access**: Open Azure Portal services instantly from Alfred
- **Fuzzy Search**: Find services by name, alias, or Korean keywords
- **Service Icons**: Visual identification with official Azure icons
- **Customizable**: Add or modify services via YAML configuration
- **Universal Binary**: Supports both Apple Silicon and Intel Macs

## Installation

1. Download the latest `azure-workflow.alfredworkflow` from [Releases](https://github.com/psi59/azure-workflow/releases)
2. Double-click the downloaded file to install in Alfred
3. Start using with the `azure` keyword

## Usage

1. Open Alfred (default: `Cmd + Space`)
2. Type `azure` followed by your search query
3. Select a service and press `Enter` to open it in your browser

### Examples

```
azure vm          # Open Virtual Machines
azure storage     # Open Storage Accounts
azure k8s         # Open AKS (Kubernetes)
azure openai      # Open Azure OpenAI Service
```

## Supported Services

| Service | Aliases |
|---------|---------|
| Virtual Machines | vm, virtual machine |
| App Service | app, web app, webapp |
| Storage Accounts | storage, blob |
| SQL Database | sql, database, db, mssql |
| Functions | function, func, serverless |
| Container Instances | container, aci |
| AKS | aks, kubernetes, k8s |
| CosmosDB | cosmos, cosmosdb, nosql, documentdb |
| Key Vault | keyvault, vault, secret, key |
| Monitor | monitor, monitoring, log |
| Azure OpenAI Service | openai, gpt, chatgpt, ai, llm |
| AI Foundry | ai foundry, ml, machine learning, aml, mlops |
| Machine Learning | ml, machine learning, aml |

## Adding Custom Services

Edit `services.yaml` to add or modify services:

```yaml
services:
  - name: Your Service Name
    aliases: [alias1, alias2, shortcut]
    url: https://portal.azure.com/#blade/...
    icon: icons/your-icon.svg
```

## Building from Source

### Prerequisites

- Go 1.21 or later
- Make

### Build Commands

```bash
# Run tests
make test

# Build binary
make build

# Create Alfred workflow package
make package

# Install directly to Alfred
make install
```

## Dependencies

- [github.com/sahilm/fuzzy](https://github.com/sahilm/fuzzy) - Fuzzy string matching
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) - YAML parser

## License

MIT License

## Author

Sangil Park ([@psi59](https://github.com/psi59))

---

> **Note**: This project was 100% generated using LLM (Large Language Model) with [Claude Code](https://claude.ai/). The entire codebase, including Go source code, tests, configuration files, and this README, was written through AI-assisted development following Test-Driven Development (TDD) methodology.
