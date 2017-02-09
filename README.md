radi-cli
--------------

this is a cli implementation of the radi api, which gives some command line
tools that can be used to run docker based applications from a project
definition.

The project started off as a port of wundertools to a docker based system, and
then ended up being an abstraction for a more generic project configuration.

This CLI is not the only possible way to use the radi-api as a command line
tool, but rather it is a first run at a typical radi-handler coordiantion for
managing projects.

Some evidence of that:
1. by default the upcloud and rancher handlers are included, even though you
may not use them.

While the radi-api takes some learning, it should be easy to fork this tool,
and rewire the base handlers used to meet any corporate standard, including
handlers from any source.

## Install the tool

### Requirements

The tools is a go implementation, but it can be build without any go runtime,
by building inside a container.
This requires a working local docker implementation (as local volume binds are
used to return the binary.)

If you want to fully build the binary yourself, there is a make system in place
that requires a go runtime

### Building inside a docker container

#### 1. retrieve this source code via git.

```
$/> git clone https://github.com/wunderkraut/radi-cli

```

#### 2. run the build script

in the repository root (`cd radi-cli`) run the build script

```
$/> sh ./build.sh
```

This produces a binary `radi`, which you can put wherever you want to put such
binaries.  If there is consensus on install paths, we can make an installer

### Manually building

When manually building the source code, you will need a working `go`
environment and the source must be checked out inside the GOPATH.

* The project uses a [vendor](https://golang.org/cmd/go/#hdr-Vendor_Directories) folder, so **go 1.6 is required as a minimum**.
* *buntu users note that full `golang` is required.  The `gccgo-go` package is not a complete go builder (no core libraries)

#### Options 1 : Manually check out the source

##### 1. use git to clone the repository into the appropriate path

```
$/> git clone https://github.com/wunderkraut/radi-cli "${GOPATH}/src/github.com/wunderkraut/radi-cli"
```

##### 2. run the build script in the radi-cli root

```
$/> cd "${GOPATH}/src/github.com/wunderkraut/radi-cli"
$/> make all
```

Which will pull dependencies and product a binary locally.  It also installs
the binary to a reasonable path.

#### Option 2 : Using go get

Try simply running:

```
$/> go get -u github.com/wunderkraut/radi-cli/radi
```

If this works, you should have a radi binary executable in the `$GOPATH/bin`
folder. ** Unfortunately go-get doesn't use the packaged libraries vendor git 
submodules properly yet, which means that dependency versioning isn't strict.**

You should then manually rebuild the binary using:

```
$/> cd "${GOPATH}/src/github.com/wunderkraut/radi-cli"
$/> make all

```

This will result in a properly built binary.

## Using the CLI

once the cli is installed, it can be used withing the scope of any project that
has a .radi folder, which demarks the root of the project (like git)

### general usage


#### First steps

running `radi help` should list all operations.

running `radi help <operation>` should give more information about each
operation and what parameters it expects.

#### Global flags

if you run radi with --debug, you will get verbose output.

#### Typical operations

There are no base or default operations. All operations are provided by the
handlers, but radi will try to at least include the init/create commands even
if no other operations are defined.

Typically, a local project will include some orchestration operations:

``` radi orchestrate.up ```
``` radi orchestrate.down ```

And a number of custom commands for the project as defined in
.radi/commands.yml

### enabling radi in an existing folder

If you have a project which you would like to use as a radi project, you can
manually add a .radi folder to the root.

An empty .radi folder will produce some errors, but it works.

### use init/create to start a project from using templates

#### init

You can intialize any folder with the command

```
$/> radi local.project.init
```

which will add some .radi files to your project

#### create

Creation adds files to your project using yml templates, which can be easily
shared, as gists, or locally available files, or raw git repo files.  Any http
url can be used.

The templates contain sequential instructions for building a project.

##### generate

The command `radi local.project.generate` attempts to build a template from the
current project, that can be used with the create command.

## Radi usage

### radi-api

The tools are an implementation that uses the radi-api, and a number of
radi-handlers to build a cli with a set of operations that can manage a local
project.

The API relied on to define handler behaviour, but also the API provides an API
builder approach for pulling together handlers and elements from different
sources.

### radi-handlers

From the API perspective, the CLI relies heavily on the `local` handler, which
uses local configuration files to define an active project.

Additionally, currently the handlers for upcloud, and rancher are included as
they are the primary tools that we are using to manager servers and
orchestration outside of local operations.

As with all radi-api implementations, operations are provided primarily by the
included handlers, when they are activated.

### Conversion of handler operations to cli behaviour

Currently the cli tool takes the simplest approach to implementing handler
operations by literally exposing any operation not marked as "internal" to the
cli.  Any operation parameter that can be interpreted is made available to the
CLI.
