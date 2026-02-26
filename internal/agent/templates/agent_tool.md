Launch a new agent that has access to the following tools: GlobTool, GrepTool, LS, View. Use this tool only after direct search attempts with GrepTool/GlobTool/LS/View are inconclusive.

<usage>
- Always try direct tools (GrepTool/GlobTool/LS/View) first
- Use the Agent tool only when direct searches remain ambiguous, broad, or inconclusive after at least one focused attempt
- Use the Agent tool when results are truncated and a refined direct query still does not identify the right target
- If you want to read a specific file path, use the View or GlobTool tool instead of the Agent tool, to find the match more quickly
- If you are searching for a specific class definition like "class Foo", use the GlobTool tool instead, to find the match more quickly
</usage>

<usage_notes>
1. Launch multiple agents concurrently whenever possible, to maximize performance; to do that, use a single message with multiple tool uses
2. When the agent is done, it will return a single message back to you. The result returned by the agent is not visible to the user. To show the user the result, you should send a text message back to the user with a concise summary of the result.
3. Each agent invocation is stateless. You will not be able to send additional messages to the agent, nor will the agent be able to communicate with you outside of its final report. Therefore, your prompt should contain a highly detailed task description for the agent to perform autonomously and you should specify exactly what information the agent should return back to you in its final and only message to you.
4. The agent's outputs should generally be trusted
5. IMPORTANT: The agent can not use Bash, Replace, Edit, so can not modify files. If you want to use these tools, use them directly instead of going through the agent.
</usage_notes>
