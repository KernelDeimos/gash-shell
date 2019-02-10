# GASh: Go Again Shell
There has certainly already been at least one shell written in go by now
(source: it's 2019), so GASh seemed like an appropriate name for this.

## Important Notes
- REGRESSIONS! This is version 0 of GASh, and therefore changes and regressions
  are expected to happen frequently.

- GASh, by design, is not a POSIX-compliant shell in its default configuration.
  However, making a POSIX-compliant GASh configuration is a long-term
  development goal.

## When will the major version be changed?

### Version 1.0.0
- Configuration for prompt 
- Configuration for which console modules are used
- If a prompt writer takes too long, a fallback prompt will be displayed
- Pipes and redirects are fully supported
- Background tasks are fully supported
- GASh is stable enough for everyday use

### Version 2.0.0
- Scriptable plugin system
- Ability to emulate other shells by replacing modules
- (not requited for 2.0.0 but...) Build-your-own-GASh tutorial exists

### Version 3.0.0
- POSIX-compliant mode

## What is a console module?
`// TODO: this`

## Ideas for Features

Idea contributions are welcome! (provided they're cool enough)

### Features

#### Reminder to take a break
The shell could remind its user to take a break if it detects regular
activity for a long period of time (ex: at least one command every 8 minutes
for a period longer than half an hour). The shell would print "Take a break"
after every command entered until the user types something like
`5-more-minutes` or `taking-a-break-now`. Also if the user types
`taking-a-break-now` and then uses the shell within a short period of time,
the shell should call its user a liar.

### Internal Commands

#### Go Fish
Since the shell is written in Golang, it is only appropriate that an
implementation of the classic game "Go Fish" should exist as an internal
shell command.

#### `gash-df` command
An alternative to the `df` command with progress bars rendered for a visual
representation of disk use. User could alias to `df` in their config if they
find this particularly useful.

### Standard Stuff

- Customizable prompt
  - Probably using the Go template package `text/template`

    For instance
    ```
    {{.User}}@{{.Host}}:{{.WorkingDirectory}}$
    ```
    to display
    ```
    username@hostname:~$
    ```
- Configuration in `.config/gash-shell`
  - Could either be TOML, YAML, or JSON(hacked to allow trailing commas)
