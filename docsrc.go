// docsrc - Program to generate source tree documentation.
package main

import (
	"docsrc/reports"
	"docsrc/treeseers"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// For now structures section (later reorganize)

// CLConfiguration - Command Line Configurations for use by flag package.
// Values are parsed into this struct for passing to main for use in
// top level conditionals.
type CLConfiguration struct {
	report0    bool
	report1    bool
	report2    bool
	concurrent bool
	stdout     bool
}

// getCLAC() - Get the command line switch values.
func getCLAC() *CLConfiguration {
	var clc CLConfiguration
	flag.BoolVar(&clc.report0, "report0", false, "Generate report0.txt, simple uncommented organized source listing.")
	flag.BoolVar(&clc.report1, "report1", false, "Generate report1.md, a relatively simple document but using all supporting files.")
	flag.BoolVar(&clc.report2, "report2", false, "Generate report2.html, report1 is generated, and report2 from report1.")
	flag.BoolVar(&clc.concurrent, "concurrent", false, "Use the concurrent version of the directory tree reader (for large trees).")
	flag.BoolVar(&clc.stdout, "stdout", false, "Write reports to stdout instead of report named files.")
	flag.Parse()
	// Apparently no err returned on zero value Flagset, which I think means this form.
	// I really did look around, and the examples are like this.
	return &clc
}

// noReportSpecified() -  boolean function to indicate if a report is
// indicated on the command line so a default alternative may be chosen.
func (pCLC *CLConfiguration) noReportSpecified() bool {
	if pCLC.report0 {
		return false
	}
	if pCLC.report1 {
		return false
	}
	if pCLC.report2 {
		return false
	}
	return true
}

// printUsage() - display USAGE statement.
func printUsage() {
	fmt.Printf("USAGE:  ./%s [argswtchs]", os.Args[0])
	flag.PrintDefaults()
}

// Generates output to either named file, or stdout, according
// to the user's command line specification.
func outputReport(stdOut bool, cStr string, rName string) {
	if stdOut {
		fmt.Print(cStr)
	} else {
		fpo, err := os.Create(rName)
		if err != nil {
			panic(err)
		}
		// close fpo on exit and check for its returned error
		defer func() {
			if err := fpo.Close(); err != nil {
				panic(err)
			}
		}()
		_, err = fmt.Fprint(fpo, cStr)
		if err != nil {
			panic(err)
		}
	}
}

// Generate report2.html from report1.md using shell markdown command.
func genHTMLfromMarkdown() (content string) {
	var out1, out2 bytes.Buffer
	cmd1 := exec.Command("which","markdown")
	cmd1.Stdout = &out1
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	cmd2fs := strings.TrimSuffix(out1.String(),"\n")
	// The shortest path is /usr/bin/markdown, which is 17...
	if len(cmd2fs) > 16 {
		fmt.Println(cmd2fs)
		cmd2 := exec.Command(cmd2fs,"report1.md")
		cmd2.Stdout = &out2
		err = cmd2.Run()
		if err != nil {
			panic(err)
		}
		content = out2.String()
		return
	} else {
//345678901234567890123456789012345678901234567890123456789012345678901234567890
		nomdmsg := `
Cannot generate HTML output report2.html because markdown utility was not found
on your computer.  report1.md is at least available.  You may generate your HTML
file by hand after installing markdown by using the command:

		markdown report1.md>out.html

`
		fmt.Print(nomdmsg)
	}
	return
}

// Generate the output report for the directory tree specified.
// Note that for now there is no way to specify a directory except
// to run the program in that directory.  This simplifies a lot of
// things, and reduces likelihood of messy mistakes.
func main() {
	pclc := getCLAC()
	pwd := os.Getenv("PWD")
	var pdo *treeseers.DirObject
	if pclc.concurrent {
		pdo = treeseers.ConLoadDirObject(pwd)
	} else {
		pdo = treeseers.SeqLoadDirObject(pwd)
	}
	if pclc.report0 || pclc.noReportSpecified() {
		content := reports.Report0(pdo, 1)
		outputReport(pclc.stdout, content, "report0.txt")
	}
	if pclc.report1 || pclc.report2 {
		content := reports.Report1(pdo, 1)
		outputReport(pclc.stdout, content, "report1.md")
		if pclc.report2 {
			content = genHTMLfromMarkdown()
			outputReport(pclc.stdout, content, "report2.html")
		}
	}
}

//345678901234567890123456789012345678901234567890123456789012345678901234567890
// End of docsrc.go
