# MAF Agent Samples

This repository contains a small set of sample Microsoft Agent Framework applications that use Azure AI Foundry and Azure AI Projects.

## Projects

- `MAF-Agent-CS-01` — a C# console app that creates an AI agent from a Microsoft Foundry project and runs a sample prompt.
- `MAF-Agent-CS-02` — a C# hosted agent sample that registers a Foundry responses endpoint with `AgentHost`.
- `MAF-Agent-GO-03` — a Go console app that creates and runs a Microsoft Agent Framework agent backed by a Foundry project.
- `MAF-Agent-GO-04` — a containerized Go hosted agent that exposes the AG-UI contract through Foundry's Invocations protocol.
- `MAF-Agents-Samples.slnx` — solution file for the two C# projects.

## Prerequisites

- .NET 10 SDK
- Go 1.26 SDK (for the Go samples)
- Azure CLI installed and signed in (`az login`)
- An Azure AI Foundry project with a valid `FOUNDRY_PROJECT_ENDPOINT`
- A deployment in that project, typically exposed through `AZURE_AI_MODEL_DEPLOYMENT_NAME` (or `AZURE_OPENAI_DEPLOYMENT_NAME` in the C# hosted sample)

## Setup

From the repository root, set the required environment variables before running the apps:

PowerShell:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<your-project-endpoint>"
$env:AZURE_AI_MODEL_DEPLOYMENT_NAME = "gpt-5-mini"
```

For the C# hosted sample:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<your-project-endpoint>"
$env:AZURE_OPENAI_DEPLOYMENT_NAME = "gpt-5-mini"
```

## Build

```powershell
dotnet build .\MAF-Agents-Samples.slnx
```

## Run

Run the C# console sample:

```powershell
dotnet run --project .\MAF-Agent-CS-01\MAF-Agent-CS-01.csproj
```

Run the C# hosted agent sample:

```powershell
dotnet run --project .\MAF-Agent-CS-02\MAF-Agent-CS-02.csproj
```

Run the Go console sample:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<your-project-endpoint>"
$env:FOUNDRY_MODEL = "gpt-5-mini"
go run .\MAF-Agent-GO-03
```

Run the Go hosted agent sample:

```powershell
Set-Location .\MAF-Agent-GO-04
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<your-project-endpoint>"
$env:AZURE_AI_MODEL_DEPLOYMENT_NAME = "gpt-5.4-mini"
go run .
```

See `MAF-Agent-GO-04/README.md` for local invocation and Foundry deployment instructions.

## Notes

These samples are intended for learning and experimentation with the Microsoft Agents SDK and Azure AI Foundry integration.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
