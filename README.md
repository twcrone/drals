Directory Aliases
=================

Simple command line utility for maintaining an alias file that can
be sourced for directory aliases for jumping around file system easily.

### Install

`go install`

### List directory aliases

`drals`

### Add a directory alias (applies to curent directory)

`drals <alias-name>`

e.g. `drals bob` when present working directory is `/Users/bob/work`
will result in `alias bob='/Users/bob/work'` beind added to `.drals`
file in $HOME directory that can be sourced with ZSH for directory
aliases.