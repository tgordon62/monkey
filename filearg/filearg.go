package filearg

import (
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
)

func Start(filename string) {
	f, err := os.Open(filename)
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
		parser.PrintParserErrors(os.Stderr, p.Errors())
		os.Exit(1)
	}

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	evaluator.Eval(expanded, env)
}
