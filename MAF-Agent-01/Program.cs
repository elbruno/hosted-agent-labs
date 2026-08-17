using Azure.AI.AgentServer.Core;
using Azure.AI.Projects;
using Azure.Identity;
using Microsoft.Agents.AI;
using Microsoft.Agents.AI.Foundry.Hosting;

var projectEndpoint = new Uri(Environment.GetEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT")
    ?? throw new InvalidOperationException("FOUNDRY_PROJECT_ENDPOINT is not set."));
var deployment = Environment.GetEnvironmentVariable("AZURE_AI_MODEL_DEPLOYMENT_NAME") ?? "gpt-5-mini";

AIAgent agent = new AIProjectClient(projectEndpoint, new AzureCliCredential())
    .AsAIAgent(
        model: deployment,
        instructions: "You are a friendly assistant. Keep your answers brief.",
        name: "HelloAgent");

var builder = AgentHost.CreateBuilder(args);                                      // 👈 new line 1
builder.Services.AddFoundryResponses(agent);                                      // 👈 new line 2
builder.RegisterProtocol("responses", endpoints => endpoints.MapFoundryResponses()); // 👈 new line 3

var app = builder.Build();
app.Run();