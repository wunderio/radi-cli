WUNDERTOOLS API
---------------

# Concepts

## API

## Operations

## Handlers


# Handler Implementations

## Base Handlers

These are handlers which provide base functionality on 
which other handlers can be based.

### File/Bytes handler

The bytes handler, which can retrieve bytes from files, uses 
files (or any bytes reader) as a source of configuration, which
can be used for pretty much all operations handling

## Implementation Handlers

### RSA

The RSA handler provides security Authentication by requiring
an RSA encryption/signing task to be passed.  Typically this 
involves passing a public key, along with some ciphertext/signature
of some passed Cleartext.

### Local

The Local handler assumes that it has local file access to a
project folder, which should contain a number of configurations
that can be used to provide configuration, orchestration and
monitoring of locally running containers, as well as local 
documentation.
This handler relies in a large part on the Base File/Bytes handler

## Vendor Handlers

### 