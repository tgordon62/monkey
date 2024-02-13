package main

import (
	"fmt"
	"io"
	"monkey-interpreter/evaluator"
	"monkey-interpreter/lexer"
	"monkey-interpreter/object"
	"monkey-interpreter/parser"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 1 {
		readEvalPrintLoop(user)
	} else {
		fileArgument(user)
	}

}

func readEvalPrintLoop(user *user.User) {
	fmt.Printf("Hello %s! This is the Monkey language console!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func fileArgument(user *user.User) {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, stat.Size())

	_, err = f.Read(buf)
	if err != nil {
		panic(err)
	}

	str := string(buf)
	l := lexer.New(str)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stderr, p.Errors())
		os.Exit(1)
	}

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	evaluator.Eval(expanded, env)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
