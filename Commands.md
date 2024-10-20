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
List out commands: "-list [all | global | clan {Clan Name}]"

### Add
Add a command to the bot: "-add [global | clan {Clan Name}] [slash | user | message] {description}

Add subcommand group to a command "-add-subcommand-group [global | clan {Clan Name}] {Command Name} {Group Name} {Description}
    ↳ The location of the target command comes first to reduce calls to API

Add a subcommand to a command "-add-subcommand [global | clan {Clan Name} | global-subgroup | clan-subgroup {Clan Name}] {Command Name} 
                                (global-subgroup & clan-subgroup){Subcommand Group} {Name} {Description}

Need to implement command options -> Come with ArgumentEngine?

### Delete
Delete a command: "-del [global | clan {Clan Name}] {Command Name}"