# MAF Agent Samples

This repository contains a small set of sample Microsoft Agent Framework applications that use Azure AI Foundry and Azure AI Projects.

## Projects

- `MAF-Agent-01` — a hosted agent sample that registers a Foundry responses endpoint with `AgentHost`.
- `MAF-Agent-02` — a console app that creates an AI agent from an Azure AI Foundry project and runs a sample prompt.
- `MAF-Agent-01.slnx` — solution file for the two projects.

## Prerequisites

- .NET 10 SDK
- Azure CLI installed and signed in (`az login`)
- An Azure AI Foundry project with a valid `FOUNDRY_PROJECT_ENDPOINT`
- A deployment in that project, typically exposed through `AZURE_AI_MODEL_DEPLOYMENT_NAME` (or `AZURE_OPENAI_DEPLOYMENT_NAME` in the second sample)

## Setup

From the repository root, set the required environment variables before running the apps:

PowerShell:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<your-project-endpoint>"
$env:AZURE_AI_MODEL_DEPLOYMENT_NAME = "gpt-5-mini"
```

Or for the second sample:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<your-project-endpoint>"
$env:AZURE_OPENAI_DEPLOYMENT_NAME = "gpt-5-mini"
```

## Build

```powershell
dotnet build .\MAF-Agent-01.slnx
```

## Run

Run the hosted agent sample:

```powershell
dotnet run --project .\MAF-Agent-01\MAF-Agent-01.csproj
```

Run the console sample:

```powershell
dotnet run --project .\MAF-Agent-02\MAF-Agent-02.csproj
```

## Notes

These samples are intended for learning and experimentation with the Microsoft Agents SDK and Azure AI Foundry integration.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
