# byreaddir.go

This is a treeseer package activity to load the directory tree with a single
threaded recursive set of calls associated with the directory nodes.  This
will likely be the default call as most of the time a source tree is small
enough that concurrency is not justiifed.  The concurrent case could be programmed
exclusively with more economy, but having the two in parallel allows for what
I see as helpful comparison of the code difference between the two package files
byreaddir.go and goreaddir.go.
