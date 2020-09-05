package core

import "github.com/jhudsonsg/interpreter/token"

// PalavrasChaves - variável que armazena todos os palavras chaves da linguagem.
var PalavrasChaves = map[string]string{
	"inicio":     token.PALAVRA_RESERVADA_INICIO,
	"fim":        token.PALAVRA_RESERVADA_FIM,
	"entao":      token.PALAVRA_RESERVADA_ENTAO,
	"se":         token.PALAVRA_RESERVADA_SE,
	"enquanto":   token.PALAVRA_RESERVADA_ENQUANTO,
	"definir":    token.PALAVRA_RESERVADA_DEFINIR,
	"inteiro":    token.PALAVRA_RESERVADA_INTEIRO,
	"real":       token.PALAVRA_RESERVADA_REAL,
	"cadeia":     token.PALAVRA_RESERVADA_CADEIA,
	"logico":     token.PALAVRA_RESERVADA_LOGICO,
	"verdadeiro": token.TIPO_LOGICO,
	"falso":      token.TIPO_LOGICO,
	"e":          token.OPERADOR_LOGICO_E,
	"ou":         token.OPERADOR_LOGICO_OU,
}
