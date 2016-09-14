LOCAL
=====

Provide a Handler, and Operations, based on a local project, meaning
based on current local path, and the configuration files contained
therein (and perhaps also in the user home folder somewhere)

## Orchestration

Orchestration is currently handed off to the libcompose handler. This
local handler provides all orchestration operations via small wrapper
that set some of the configs.
