package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // identifiers + literals
  IDENT = "IDENT"
  INT   = "INT"

  // Operators
  ASSIGN   = "="
  PLUS     = "+"
  MINUS    = "-"
  BANG     = "!"
  ASTERISK = "*"
  SLASH    = "/"

  LT       = "<"
  GT       = ">"

  // Delemiters
  COMMA     = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET      = "LET"
  TRUE     = "TRUE"
  FALSE    = "FALSE"
  IF       = "IF"
  ELSE     = "ELSE"
  RETURN   = "RETURN"
)

var keyword = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keyword[ident]; ok {
    return tok
  }
  return IDENT
}
