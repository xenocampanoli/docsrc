#!/bin/bash

#2345678901234567890123456789012345678901234567890123456789012345678901234567890
# Conditionally add GO environment variables to .profile, including GOPATH,
# GOBIN, GOROOT, and the $GOBIN addition to PATH.  Then assert need for
# markdown program if it isn't found, and finally build.

if [[ -z $(which markdown) ]]
then
	echo <<EOEM
This is a markdown oriented program, and the shell markdown command is used
in the app.  BlackFriday is not presently a link option, though it may be
considered later.  You may add markdown to your system if you use a debian
family OS with the command:  sudo apt-get install markdown<enter> (and then
provide your password, presumably).  Markdown is a pretty handy tool, and
this app, for what it is, was conceived as an exercise largely around use
of markdown (note that I am still not taking the step yet of claiming this
tool will be useful).
EOEM
fi

here=$PWD
cd ../..
GOPATH=$PWD
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN 
cd $here

if [[ -z $(grep GOPATH $HOME/.profile) ]]
then
	cat >>$HOME/.profile <<EOPADDENDUM
# Modify the following to your environment before running:
# Xeno like's vi everywhere:
set -o vi
export GOROOT=/usr/lib/go
# Xeno puts works in progress in a shop directory:
export GOPATH=$GOPATH
export GOBIN=$GOBIN
export PATH=$PATH:$GOBIN 
EOPADDENDUM
fi

. $HOME/.profile
mkdir -p $GOPATH

go install docsrc

./docsrc_test.bash docsrc

# As per everywhere else, first place to look if this doesn't
# go well and the problem isn't obvious is:  http://golang.org/doc/

# End of docsrc/build.bash
