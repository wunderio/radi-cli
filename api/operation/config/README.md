CONFIG
======

Config operations exist to give the system access to large 
keyed bytestreams, that can be used to configure other aspects
of the system.

For example, if there are config operations that can retrieve
application settings from a local file, or from a key value storage.
The settings operations can then be used to interpret the config
bytestream as parseable settings information (like as yml.)

Config exists primarily as internal operations, but it is 
really in place because it makes sense to create an abstract
method to provide configuration of all other operation types.

- Settings: Settings be provided as a yml or json array
- Commands: Commands can be interpreted as yml
- Documentation: Docs/Help could by yml
etc.


# Wrapper

The ConfigWrapper exists as a re-usable base, in order to
make the Config operations easier to use, and also to allow a
single config connection to be reused across different operations
making it easier to handle changes to config across the application.

# Connector

A connector is a connection/socket concept that can be used to
produce config value for multiple operations from a single sockey.  
The connector can be used in multiple operations, and can then
handle maintaining data across changes.
