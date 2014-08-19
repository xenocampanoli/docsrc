# docsrc - Makes a document from your golang source tree.

docsrc is a program that tries to help you make a document from a go source
tree, utilizing supplemental documents that you may choose to omit.  All Go
source files in your tree, however, are always included.

## Installation

There is a simple install script, build.bash, in this directory.

## USAGE

The program, when set up properly, should generate a fairly bookish looking
markdown file with the following command executed in the top level source
directory you wish to feature:

```
docsrc -report1
```

If you have docsrc properly installed somewhere, it should display help more
or less as follows:

```
$ docsrc -help
Usage of docsrc:
  -concurrent=false: Use the concurrent version of the directory tree reader (for large trees).
  -report0=false: Generate report0.txt, simple uncommented organized source listing.
  -report1=false: Generate report1.md, a relatively simple document but using all supporting files.
  -report2=false: Generate report2.html, report1 is generated, and report2 from report1.
  -stdout=false: Write reports to stdout instead of report named files.
---snip---
```
## Builds

As with all golang programs, but with minor provisions elaborated next, you can
generate general executeables using the standard Go sequence:

go install docsrc.go

These provisions are as follows:

1.  Use this on a Unix-like system with go installed.  For now I have not tested
anything on Windows/M$, and I do not wish to spend the time.
2.  Create all the standard GO environment variables, GOPATH, GOBIN, GOROOT, etc
as per the standard instructions found in [http://golang.org/doc](http://golang.org/doc).
3.  Install markdown shell program (TODO:  Caution to self say: , or blackfriday go package,
or both, but that is also a step I need to finish).

TODO:  I also have a build.bash file that will try to do these setups for you, but it is as yet
feeble and unfinished.
This program works by simply running it in the uppermost node of the source
branch you wish to see a document for.  Existing reports will work as follows:

### report0

This report just puts together the source code in the hierarchy selected
with source directories and go file names as titles.  No additional documentation
will be displayed.  The intent here is to be able to create a quick file which
can be looked through quickly in an arbitrary file reader for optimal quick code
reading and searching.

### report1

This report provides up to five sections for each source directory, and up to
three for each Go source file.  In each directory node:

	1.  If Directory_description.md exists, it will be listed at the top of
	the section for that directory.
	2.  If the README.md exists, it will be listed next.
	3.  Then reporting is done specific to the sourccode for all Go source
	files in this directory, including description and notes specific to these
	(see below).
	4.  Then reporting on the node subdirectories, effectively a nested
	operation, is done.
	5.  Finally, contents of Directory_notes.md is listed, if this file exists.

Regarding the source code files, given a name 'sourcecode.go' of a given file,
these four aspects are provisioned:

	1.  The full file path is used as the heading of the section.
	2.  If a "sourcecode_description.md" file exists, it is listed next.
	3.	The source file content is then listed with line number to the left of
	each line.
	4.  Finally, contents of "sourcecode_notes.md" is listed, if this file exists.

Note no "sourcecode_README.md" file is provisioned for.  It is expected that
any such details will be provided in the sources themselves.

### report2

This report is a combination, and will first run report1, then run the markdown
shell command (or a golang version if I find one) to generate the HTML files.
If they exist in the running directory node, the following files will be tacked
onto the beginning and end of the resulting HTML file:

	1.  Prefatory.html
	2.  Ending.html

If these do not exist, proper HTML beginning and ending tags will be inserted
instead, with the author's guesses as to helpful style (which is to say you
probably do want to make your own).

TODO:  Part of the intention is to use markdown files for all the initial reports.
However given time I do hope to also create a report3 that generates more nicely
detailed HTML, and a report4 that squirts out reasonable Docbook XML.  However, for
now this is my humble version of Terry Gilliam's "Time Bandits", and I really prefer
to get to my own version of "Brazil" next before doing further refining here.

### Source

The source files for this application are docsrc.go at the top node, and various
others in the treeseers and reports subdirectories which by name represent the
corresponding packages imported to docsrc.go.  It is hoped that the use of this
application will then correspond to one of the best ways to study the application,
as well, perhaps, the Golang programming environment.

