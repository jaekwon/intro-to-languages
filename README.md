## Introduction

We're going to learn how to create a programming language from scratch.

We're going to create a langauge interpreter, rather than a compiler.  A
compiler typically turns program code into machine bytecode which can be run
directly by the CPU.

```
	Steps of a compiler

	Program code --(tokenizer)-->
	  Stream of tokens --(parser)-->
		Abstract syntax tree representation --(optimizer)-->
		  Abstract syntax tree representation --(compiler)-->
			Machine bytecode --> written to a file (e.g. myprogram.exe)
```

A token is just a short string, e.g. "function", "+", or ";". The tokenizer
chops your program code into these little tokens, and discards things like
newlines, whitespaces, and comments as appropriate.

The abstract syntax tree is an in-memory representation of the program code,
made of structures/objects (rather than bytes as in the program code), and it's
"pure" with no need for human nicities like semi-colons or newlines.

The result of a compiler is a program file.  When you tell the operating system
to run a file, the operating system will load the program file into memory, and
from there the hardware (and operating system) takes over.

On the other hand, an interpreter skips the "write machine bytecode to file"
step and just runs the program directly from within the interpreter program.

```
	Steps of an interpreter

	Program code --(tokenizer)-->
	  Stream of tokens --(parser)-->
		Abstract syntax tree representation --(optimizer)-->
		  Abstract syntax tree representation --(interpreter)-->
			Running instance of the program
```

Not all parsers need tokenizers; depending on the type of parser algorithm it
is possible to parse a program into an AST (abstract syntax tree) without the
tokenizing step.  In other words, some parsers can operate directly on the
program code bytes.

The optimizing step is obviously optional.

The interpreter that we'll design has the following steps:

```
	Steps of the Zen interpreter

	Program code --(parser)-->
	  Abstract syntax tree representation --(interpreter)-->
		Running instance of the program
```

There are many types of parsers.  LL parsers, LR parsers, top-down recursive
descent parsers, etc.  It's an obscure field of computer science that people
generally learn in university and promptly forget about.

One thing to keep in mind when learning about different types of parsers is
that evolution in the field of parsers is slow (because it's not the most
important part of a programming langauge... any parser that works will do, and
if a particular parser is insufficient, you can always tweak the langauge
grammar a little to accomodate, such as introducing semi-colons or requiring
certain prefixes to hint the parser as necessary.

Top-down recursive descent parsers are arguably the most intuitive type of
parser, so we will use them here.

An interpreter can easily borrow features from the host language.  Lets say the
interpreter is written in Golang, which is a garbage-collected language.  When
you create an object/structure in the Zen language (the language that we are
designing), there will be a corresponding structure in the underlying Golang
interpreter program.  In order to garbage-collect the Zen object/structure, all
we need to do is ensure that the underlying Golang structure is garbage
collected when appropriate, by removing all references to that structure (in
the Golang interpreter).

## How this tutorial is structured

This tutorial is composed of self-contained chapters.  Each subsequent chapter
will introduce new concepts, and we'll build out the Zen interpreter slowly
from chapter to chapter.
