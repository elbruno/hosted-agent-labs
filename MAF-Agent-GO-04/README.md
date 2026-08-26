# MAF-Agent-GO-04

This sample hosts a Microsoft Agent Framework Go agent as a containerized Microsoft Foundry Hosted Agent. It exposes the Agent Framework Go AG-UI handler at `/invocations`, matching Foundry's Invocations protocol, and provides `/readiness` for platform health checks.

Foundry direct-code deployment currently supports Python and .NET runtimes. This Go sample therefore uses the supported custom-container deployment path defined by `Dockerfile` and `azure.yaml`.

## Prerequisites

- Go 1.26 or later
- Docker or another OCI-compatible container builder, only for local image builds
- Azure CLI and Azure Developer CLI (`azd`)
- The `azure.ai.agents` and `azure.ai.projects` azd extensions
- An Azure account that can create or use a Microsoft Foundry project

## Run locally

Set the project endpoint and model deployment:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:AZURE_AI_MODEL_DEPLOYMENT_NAME = "gpt-5.4-mini"
go run .
```

The server listens on port `8088` by default. Override it with the `PORT` environment variable.

Check readiness:

```powershell
Invoke-RestMethod http://localhost:8088/readiness
```

Invoke the AG-UI agent:

```powershell
$body = @{
    threadId = "thread-1"
    runId = "run-1"
    state = @{}
    messages = @(
        @{
            id = "message-1"
            role = "user"
            content = "Hello! Tell me a fun fact about Go."
        }
    )
    tools = @()
    context = @()
    forwardedProps = @{}
} | ConvertTo-Json -Depth 5

Invoke-WebRequest `
    -Uri http://localhost:8088/invocations `
    -Method Post `
    -ContentType "application/json" `
    -Headers @{ Accept = "text/event-stream" } `
    -Body $body
```

## Build the container

Foundry Hosted Agents require a Linux AMD64 image:

```powershell
docker build --platform linux/amd64 -t maf-agent-go-04 .
```

## Deploy to Microsoft Foundry

The manifest enables an Azure remote container build, so deployment does not require a local Docker daemon. Run these commands from this folder:

```powershell
azd provision
azd deploy
```

The `azure.yaml` file provisions `gpt-5.4-mini` by default. To target an existing Foundry project or model deployment, initialize or configure the azd environment with that project's resource ID and deployment name before deploying.

After deployment, invoke the agent with an AG-UI payload:

```powershell
azd ai agent invoke '{"threadId":"thread-1","runId":"run-1","state":{},"messages":[{"id":"message-1","role":"user","content":"Hello!"}],"tools":[],"context":[],"forwardedProps":{}}'
```
