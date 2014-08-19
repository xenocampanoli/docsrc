# treeseer package

## Function:

If it isn't clear, the primary calling program is docsrc.go, in the directory above
this one.

This package should simply read files of interest (*.go, *.md only for now) for use
in a quick example documenting scheme for a golang source tree.  I would like tests
to be reasonably supportive, but not thickly complete, but I am uneasy about the
test needs in this new language, so please I am open to suggestions, especially
from seasoned TDD/BDD afficionados in golang.

The data gets put into a tree of DirObject and FileObject structs, which are defined
in common.go.  This tree is then handed to an arbitrary reporter in the reports
package for use in generating output.

## Intents and Purposes:

I am calling this tree seer as a deprecative nom-de-guerre for software that is not
intended to be a whirlwind of lexing/parsing efficiency, but is only to get a simple
job done, to serve as a learning exercise for me, and to perhaps an example for
others interested in Go.

Tree Seer loads go source files and my own idea of documentation (which is optional) into a
struct tree for use by presentation reports.  My intention is to be able to helpfully
document an arbitrary go programming project tree with simple markdown, html or other
familiar format files for efficient understanding, AND instructional purposes.  Since this
is my first golang project I will humbly admit I am not sure this is ultimately helpful to
anyone but me as an exercise, but that it will merely do what I intended.  My primary
intentions with the activity were to:  1) learn go, 2) make something that others could
learn from that is clear, and well explained, and 3) possibly make something useful with an
unlikely bit of luck, but which is so simple that at least it could be expanded upon, if
realistic, to something more real anyway, if it were at least in some ballpark of long term
usefulness.

Originally I was to make a byreaddir.go and a byfpwalk.go, but filepath walk does a lexical
scan rather than a structural one of the directory tree, so it breaks the original idea I had
in mind for the structs/objects trees that I wanted to drive reports off of.  Given you end up
with an ordered list, (but out of order from the tree) it does make sense for some reporting purposes, but I want a structured
data representation, and then have the reports decide what to do with this, so I want only
availability created at load time, and not formatting or re-organization.  Given there is not
an issue with the amount of items likely to be seen for documenting, this just makes sense.

Tree seer therefore is divided between byreaddir.go, and common.go, the latter still which could
conceiveably be useful in the future if I add other load methods.  However, it is admittedly
counter to practice as typically golang philosophy is DRY and otherwise minimalist.  Hence,
should the possibility go away of future multiple loaders, or perhaps even not, the files and
tests should be merged into one.
