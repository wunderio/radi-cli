# wundertools-go

Wundertools docker prototype rewritten in go

## To Install

### GO

wundertools-go is a GO app, and therefore requires that you install
and configure GO on your system before you use it.  This usually involves
installing GO, and settings a $GOPATH to point to where you want go
to compile and download source.  Additionally, you may want to add the 
GO ./bin path to your system $PATH, in order to get easy access to any
GO binaries.

### wundertools-go

With GO working in your system, you can get the wundertools-go binary
simply by running this command.

    $/> go get github.com/james-nesbitt/wundertools-go

If you added the GO ./bin path to your PATH then you are ready to go,
otherwise you will need to run the full path to the bin, alias it, or
sym-link it to one of the various system /bin paths.

## Status

While still very early in development, the tool allows the basic operations
needed to get a system up and running, plus it includes the ported init
system from coach, which allows templating of project configurations.

## Using wundertools-go

### Managing project containers

Wundertools can be used as a docker-compose wrapper, using the compose
command.  Not all compose commands are implemented, but most are.

````
$/> wundertools-go compose up
$/> wundertools-go compose stop --quick
$/> wundertools-go compose start 
$/> wundertools-go compose down
````

### Getting information about the project

Wundertools has an info command, which will output status of services
and containers.

    $/> wundertools-go info

### Adding commands to wundertools-go

Wundertools-go has an extensible command implementation, which can be used
to define additional containerized, or shell commands that can be added 
to the system.  This involved create a .wundertools/commands.yml file
defining the commands.

This is typically used to run disposable containers that connect to the
running services, to perform operations.

### Starting a new project

All you need to do to mark a project as a wundertools-go target is to 
create a .wundertools folder in your project.

Wundertools-go also expects that you have the following files:

#### .wundertools/settings.yml

This yml contains base configurations for the project

#### docker-compose.yml

This docker-compose file is used to manage containers for services in the project

### using init to start a project

wundertools-go has an init system, which can add wundertools-go resources to a 
folder, either creating a new project, or adding to an existing project. The init 
system is template based, usually using a single yml file to define files, and 
actions taken to create a system.

If you have a template file, local or remote, you can use the init system like
this:

    $/> wundertools-go init <path-to-template-yml-file>

#### init-generate

Wundertools also contains a command which tries to generate .wundertools/init.yml as
a new init template, based on the current project state.  The generate tries to
template any text files, track any git repositories that are needed.  It is not yet
a complex tool, but can form the basis of initial template generation/

    $/> wundertools-go init-generate
