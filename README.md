<!-- SPDX-License-Identifier: CC-BY-4.0 -->
<!-- Copyright 2022 Atlan Pte. Ltd. -->

<p align="center">
  <img src="atlan-go-logo.png" alt="Atlan Go SDK Logo" width="800" />
</p>

# Atlan Go SDK
This repository houses the code for a Go SDK to interact with [Atlan](https://atlan.com).


## Installing for Development

### Initial Setup
To get started developing or using the SDK:

1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```

2. Ensure you have Go 1.19 or later installed. You can verify your Go version with:
   ```bash
   go version
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Code Formatting
Ensure your code adheres to the repository's formatting guidelines before committing. You can use the following command to format the code:
```bash
go fmt ./...
```

### Environment Setup
To run integration tests or interact with the Atlan API, you'll need to configure your environment:

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Update the `.env` file with your Atlan API key and base URL.

3. Load the environment variables:
    - For macOS/Linux:
      ```bash
      export $(cat .env | xargs)
      ```
    - For Windows (PowerShell):
      ```powershell
      Get-Content .env | ForEach-Object {
          if ($_ -match '^(.*?)=(.*)$') {
              $env:($matches[1]) = $matches[2]
          }
      }
      ```

### Testing the SDK

You can run all tests in local development with the following command:
```bash
go test -v ./...
```

---

License: [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/),
Copyright 2022 Atlan Pte. Ltd.