# Discord bot for minecraft docker server

### Recommand docker image - https://github.com/karlrees/docker_bedrockserver

<hr>

Discord Bot Functions
 - container status
 - container start
 - *container stop
 - *container restart
 - *lockmode on/off
 - *add whitelist
 - *remove whitelist
 - *list whitelist
 - list player
 - about

"*" Mods only commands.Other public commands can be on/off by lockmode.

conf.json

```
{
     "mod" : "<discord id>",
     "token" : "<discord bot token>",
     "container_id" : "<container id>",
     "bot_id" : "<discord bot id>"
}
```