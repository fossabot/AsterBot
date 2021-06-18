# Discord bot for minecraft docker server

### Recommand docker image - https://github.com/karlrees/docker_bedrockserver

<hr>


[![.github/workflows/go.yml](https://github.com/peterzam/AsterianBot/actions/workflows/go.yml/badge.svg)](https://github.com/peterzam/AsterianBot/actions/workflows/go.yml)


Discord Bot Functions
```
 "aster status"                        - container status
 "aster start"                         - container start
 "aster stop"                          - container stop*
 "aster restart"                       - container restart*
 "aster lockmode on/off"               - lockmode on/off*
 "aster add whitelist [<username>]"    - add whitelist*
 "aster remove whitelist [<username>]" - remove whitelist*
 "aster list whitelist"                - list whitelist*
 "aster list"                          - list player
 "aster about"                         - about
```
\* Mods only commands.Other public commands can be on/off by lockmode.

conf.json

```json
{
     "mod" : "<discord id>",
     "token" : "<discord bot token>",
     "container_id" : "<container id>",
     "bot_id" : "<discord bot id>"
}
```
