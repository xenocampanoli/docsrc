#!/bin/bash
#
#	docsrc_test.bash - Run shell level functional tests on docsrc.go.
#
################################################################################
#2345678901234567890123456789012345678901234567890123456789012345678901234567890

function printUsage
{
	echo "USAGE:  ./test_regression.bash <docsrccmd>"
	echo " docsrccmd may be any valid evocation of this application.  As"
	echo " typically it is most useful after being formally installed with"
	echo " a helpful PATH assignment, in this case the command would be"
	echo " 'docsrc' in the top directory of interest.  If it is compiled "
	echo " but available in a specific location outside an execute path then"
	echo " you of course must specify the correct path.  If you are executing"
	echo " from the app's source tree the command is 'go run docsrc.go'."
}

function pTestR
{
	_testText=$1
	_pass=$2

	if [[ -z $_testText ]]
	then
		echo "Blank test text value passed to pTestR  Invalid Usage."
		exit 99
	fi
	if [[ $_pass != true && $_pass != false ]]
	then
		echo "Invalid shell boolean passed to pTestR for '_pass':  |$_pass|."
		exit 99
	fi
	if $_pass
	then
		echo -n "PASS:  "
	else
		echo -n "FAIL:  "
	fi
	echo "$_testText"
}

################################################################################
#2345678901234567890123456789012345678901234567890123456789012345678901234567890
# Main Procedure

if [[ $1 = '-?' || $1 = '-h' || $1 = '-help' || $1 = '--help' ]]
then
	echo "trace $1"
	printUsage
	exit 0
elif (( $# == 0 ))
then
	DOCSRCCMD="go run docsrc.go"
elif (( $# == 1 ))
then
	DOCSRCCMD=$1
else
	echo "Zero or One argument is required:  $#"
	printUsage
	exit 1
fi


tar zcvf ~/cellar/godocsrctestbackup$(date +%Y%m%d%H%M%S).tgz . >/dev/null

rm -rf tmp ; mkdir tmp

$DOCSRCCMD -help >/tmp/nothin.lst 2>/tmp/helpswitchoutput.lst
if [[ -z $(cat /tmp/nothin.lst) ]]; then r=true; else r=false; fi
pTestR "-help should not write to stdout." $r 
if [[ -n $(cat /tmp/helpswitchoutput.lst | grep Usage) ]]; then r=true; else r=false; fi
pTestR "-help should provide a Usage statement to stderr." $r 

$DOCSRCCMD
cp report0.txt tmp/report0.txt.test1.f1
$DOCSRCCMD -concurrent
cp report0.txt tmp/report0.txt.test1.f2
cmp tmp/report0.txt.test1.f1 tmp/report0.txt.test1.f2
result=$?
if (( result == 0 )); then r=true; else r=false; fi
pTestR "Test report0 with both sequential and concurrent loading." $r


$DOCSRCCMD -report1
cp report1.md tmp/report1.md.test2.f1
$DOCSRCCMD -report1 -concurrent
cp report1.md tmp/report1.md.test2.f2
cmp tmp/report1.md.test2.f2 tmp/report1.md.test2.f2
result=$?
if (( result == 0 )); then r=true; else r=false; fi
pTestR "Test report1 with both sequential and concurrent loading." $r 

if [[ -n $($DOCSRCCMD -stdout) ]]; then r=true; else r=false; fi
pTestR "-stdout should provide report to stdout." $r 

#TODO:
# Performance tests, especially with comparison report with small and large source trees
# (perhaps just clone Gorilla or something).
# End of docsrc_test.bash
