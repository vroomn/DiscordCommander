# Commands the program implements
A "{}" indicates a parameter to be set by the caller
A "[]" indicates it could be any of the provided options
A "()" indicates conditional need to have

### Authorization
Attach proper authorization information: "-at {appID} {authToken}"
Notes:
    This command will overwrite a preexisting cache
    {authToken} requires the "Bot" or "Bearer" pre-tag (https://discord.com/developers/docs/interactions/application-commands#registering-a-command)

### List
List out commands: "-list [all | global | server {Server Name}]"

### Add
Add a command to the bot: "-add [global | server {Server Name}] [slash | user | message] {Command Name} {description}"

Add subcommand group to a command "-add-scg [global | server {Server Name}] {Command Name} {Group Name} {Description}"
    ↳ The location of the target command comes first to reduce calls to API

Add a subcommand to a command "-add-sc [global | server {Server Name} | global-subgroup | server-subgroup {Server Name}] {Command Name} (global-subgroup & server-subgroup){Subcommand Group} {Name} {Description}"

Name associated subcommand groups and subcommands will be combined with the primary add command programatically under one API call

### Delete
Delete a command: "-del [global | server {Server Name}] {Command Name}"