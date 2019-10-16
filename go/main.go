package main

import (
	"fmt"
	"io"
	"os"
)

type Token struct {
	Op  byte
	Pos int
}

func compile(f *os.File) ([]Token, error) {
	var (
		err    error
		token  Token
		token1 Token
	)

	pos := -1
	pos1 := -1
	b := make([]byte, 1)
	tokens := make([]Token, 0)
	stack := make([]int, 0)
	for {
		_, err = f.Read(b)
		if err != nil {
			if err != io.EOF {
				return nil, err
			} else {
				if len(stack) != 0 {
					pos = stack[len(stack)-1]
					err = fmt.Errorf("no corresponding ] for [ at pos %d", pos)
					return nil, err
				} else {
					return tokens, nil
				}
			}
		}
		pos++

		switch b[0] {
		case '>':
			token = Token{Op: '>', Pos: -1}
			tokens = append(tokens, token)
		case '<':
			token = Token{Op: '<', Pos: -1}
			tokens = append(tokens, token)
		case '+':
			token = Token{Op: '+', Pos: -1}
			tokens = append(tokens, token)
		case '-':
			token = Token{Op: '-', Pos: -1}
			tokens = append(tokens, token)
		case '.':
			token = Token{Op: '.', Pos: -1}
			tokens = append(tokens, token)
		case ',':
			token = Token{Op: ',', Pos: -1}
			tokens = append(tokens, token)
		case '[':
			token = Token{Op: '[', Pos: -1}
			tokens = append(tokens, token)
			stack = append(stack, pos)
		case ']':
			if len(stack) == 0 {
				err = fmt.Errorf("no corresponding [ for ] at pos %d", pos)
				return nil, err
			}
			pos1 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			token1 = Token{Op: '[', Pos: pos}
			token = Token{Op: ']', Pos: pos1}
			tokens[pos1] = token1
			tokens = append(tokens, token)
		default:
			err = fmt.Errorf("invalid operator %c", b[0])
			return nil, err
		}
	}
}

// 1. restrict memory to 3000
// 2. ignore byte overflow check
func vm(tokens []Token) {
	var token Token

	memory := make([]int, 3000)
	index := 0
	pos := 0
	for {
		if pos == len(tokens) {
			break
		}
		token = tokens[pos]

		switch token.Op {
		case '>':
			index++
			pos++
		case '<':
			index--
			pos++
		case '+':
			memory[index]++
			pos++
		case '-':
			memory[index]--
			pos++
		case '.':
			fmt.Printf("%c", memory[index])
			pos++
		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)
			memory[index] = int(b[0])
			pos++
		case '[':
			if memory[index] == 0 {
				pos = token.Pos
			}
			pos++
		case ']':
			pos = token.Pos
		}
	}
}

func main() {
	var (
		f      *os.File
		tokens []Token
		err    error
	)

	f, err = os.Open("../test/integer.bf")
	if err != nil {
		panic(err)
	}
	tokens, err = compile(f)
	if err != nil {
		panic(err)
	}
	vm(tokens)
	fmt.Println()

	f, err = os.Open("../test/cycle.bf")
	if err != nil {
		panic(err)
	}
	tokens, err = compile(f)
	if err != nil {
		panic(err)
	}
	vm(tokens)
	fmt.Println()

	f, err = os.Open("../test/helloworld.bf")
	if err != nil {
		panic(err)
	}
	tokens, err = compile(f)
	if err != nil {
		panic(err)
	}
	vm(tokens)
}
