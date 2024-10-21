# Go-Interpreter
Writing an interpreter with Go - reading and implementing the book from Thoresten Ball.

# Introduction
This is for people  who love to look under the hood. **Interpreters are magical**: the one fundamental attribute they all 
share is that they  take source code and evaluate it without producing some visible, intermediate result that can
later be executed.

There are interpreters that parse the source code,  build an abstract syntax tree (AST) out of it and then evaluate this tree. 
This type of interpreter is sometimes called “tree-walking” interpreter, because it “walks” the AST and interprets it.
The interpreter we’re building will tokenize and parse Monkey source code in a REPL, building up an internal representation of
the code called abstract syntax tree and then evaluate this tree. It will have a few major parts:
- the lexer
- the parser
- the Abstract Syntax Tree (AST)
- the internal object system
- the evaluator

# Lexing
## Lexical Analysis
Represent our source code in other forms that are easier to work with.
We’re going to change the representation of our source code two times before we evaluate it: 
> Source Code -> Tokens -> Abstract Syntax Tree.

The first transformation, from source code to tokens, is called “lexical analysis”, or “lexing” for short. It’s done by 
a lexer (tokenizer or scanner). **Tokens** themselves are small, easily categorizable data structures that are then fed to the parser,
which does the second transformation and turns the tokens into an “Abstract Syntax Tree”.

## REPL
**REPL** stands for “Read Eval Print Loop”. Sometimes
the REPL is called “console”, sometimes “interactive mode”. The concept is the same: the
REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the
interpreter and starts again.
# Parsing
> A parser is a software component that takes input data (frequently text) and builds
a data structure – often some kind of parse tree, abstract syntax tree or other
hierarchical structure – giving a structural representation of the input, checking for
correct syntax in the process. […] The parser is often preceded by a separate lexical
analyser, which creates tokens from the sequence of input characters;

In most interpreters and compilers the data structure used for the internal representation of
the source code is called a “syntax tree” or an “abstract syntax tree” (AST). The
“abstract” is based on the fact that certain details visible in the source code are omitted in the
AST.

Parsers take source code as input (either as text or tokens) and produce
a data structure which represents this source code. While building up the data structure, they
unavoidably analyse the input, checking that it conforms to the expected structure. Thus the
process of parsing is also called _**syntactic analysis**_.

There are a lot of parser generators, differing in the format of the input they accept and the
language of the output they produce. The majority of them use a context-free grammar (CFG)
as their input. A CFG is a set of rules that describe how to form correct (valid according to
the syntax) sentences in a language. The most common notational formats of CFGs are the
Backus-Naur Form (BNF) or the Extended Backus-Naur Form (EBNF).

There are two main strategies when parsing a programming language: top-down parsing or  bottom-up parsing. “Recursive 
descent parsing”, “Early parsing” or “predictive parsing” are all variations of top-down  parsing.

There are a set of “assertion functions” nearly all parsers share. Their primary purpose is to enforce  the correctness
of the order of tokens by checking the type of the next token.

## Parsing let statements
Expressions produce values, statements don’t.

> let \<identifier> = \<expression>;

## Parsing return statements

> return \<expression>;

## Parsing expressions
The parser has to know about operator precedences
where the precedence of * is higher than +.

### Pratt Parsing
“Top Down Operator Precedence” by Vaughan Pratt is an approach to parsing expressions. It was invented as an alternative
to parsers based on context-free grammars and the Backus-Naur-Form. Pratt associates parsing functions (which he calls 
“semantic code”) with single token types. A crucial part of this idea is that each token type can have two parsing functions 
associated with it, depending on the token’s position - infix or prefix.

> operator precedence == order of operations 
> 
> (_which priority do different operators have_)

