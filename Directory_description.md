# docsrc - a golang source tree document generator
## To help form documents directly from a Golang Source Tree.

People usually do not write well.  There are loads of reasons for
this, from social desires to hold resources to cognitive difficulties.
Not all tools that allow for more efficiency in writing a doc may
also provide encouragement to do so.  This tool is designed to help
write clear explanations around code without a lot of extra thought
of structure, thus leaving your personal neural networks more
available for the detail of what is actually helpful to explain.

When I first started as a programmer/consultant at the University
of Washington Health Sciences Building in 1985, early they let me
do one hour lectures teaching users how to use some of the
mainframe's facilities.  I had decided after attending the
sessions on these given by others that for most, a very simple
organization might be optimal for teaching and reinforcing the
concepts:

1.  Explain the following list of three additional items, so the
users don't get upset while going through the summary.  Explain as
quickly and clearly as possible that while listening to the summary,
they are to relax and just let it soak in, as during the detail part
all would be reinforced, and that last there would be time devoted
to their questions.  During this explanation, hand out any hardcopy
documents to be used with the session.  We had 50 minute sessions.
2.  Provide a 5 minute summary, concise and clear, but quick, with no
questions allowed, on the program, what it does, and how to use it.  
3.  Walk through the concepts to be taught during the session
exhaustively, and as clearly and concisely as possible, taking any
urgent questions, but at each point of such a question, reinforcing
that they would be allowed question time.  This section
typically should take 30 to 40 minutes of the session.  
4.  Answer questions as courteously, patiently, reponsively and
helpfully as possible, focusing on the need of the person asking,
and then on other group interests that may come to mind.  If the
session does not get taken up by questions, add explanations of
other functions to try to stimulate them, or review areas where
users had questions in the past to see if that brings up some.

I got this format from other kinds of things I'd learned in
school, from reinforcements I knew helped with me, and from other
teachers I had good conversations with about such matters.

This may not seem to fit perfectly with the present mechanism, but
it is intended primarly to accommodate this kind of pattern.  The
five minute introduction should end up in the Directory_description.md
files, followed by the README.md with the more full set of 
explanations for a new user.  Then, the source code itself is included
as a complement, and finally notes are provided at the end to
accommodate questions.  This application also will utilize sourcefile
specific documents:  1) sourcefile_description.md and sourcefile_notes.md,
however no sourcefile_README.md is provided for, as I presume that level
of detail should be documented in the source file itself, which by design
should be included in the reports in their entirety.

It may be that much solid but simple instruction can be weilded with such
a format in ways that reduce the cost of labor on both sides.  It is
certainly not rocket science, and there are likely other formats that can
be similarly helpful.  Anyway, that's my guess.  The report0
format does NOT accommodate this idea, as it is just for the source
code alone, but I hope all the others will.

Regarding usage then, to run this locally:

go run docsrc.go

from the top level directory where docsrc.go is natively found.

But of course, where it should become useful will be when it is
installed for use in arbitrary source trees for making application
documents.
