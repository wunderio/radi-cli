BYTESOURCE 
==========

Bytes is a BaseHandler which allows Operations to be provided
which are based on configuration from Bytes streams.
Typically Byte streams will come from files, so there is a file
based handler available, but any io.Reader can be used as a bytes
stream source, which should allow any Reader provided to be
used. If you have a bytes array, there is an option for a []byte
wrapped in a streamer.
