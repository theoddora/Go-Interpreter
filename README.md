# Go-Interpreter
Writing an interpreter with Go - reading and implementing the book from Thoresten Ball.
# Introduction
This is for people  who love to look under the hood. **Interpreters are magical**: the one fundamental attribute they all share is that they
take source code and evaluate it without producing some visible, intermediate result that can
later be executed!
There are interpreters that parse the source code,  build an abstract syntax tree (AST) out of it and then evaluate this tree. This type of interpreter
is sometimes called “tree-walking” interpreter, because it “walks” the AST and interprets it.
The interpreter we’re building will tokenize and parse Monkey source code in a REPL, building up an internal representation of
the code called abstract syntax tree and then evaluate this tree. It will have a few major parts:
- the lexer
- the parser
- the Abstract Syntax Tree (AST)
- the internal object system
- the evaluator