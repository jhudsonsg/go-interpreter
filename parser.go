package main

import (
	"log"

	"github.com/jhudsonsg/interpreter/token"
)

var bufferTokens []token.Token
var tamanhoBufferTokens int

func carregarBufferTokens() {
	for token := proximoToken(); token.Type != ""; token = proximoToken() {
		bufferTokens = append(bufferTokens, token)
	}

	atualizarTamanhoBuffer()
}

func atualizarTamanhoBuffer() {
	tamanhoBufferTokens = len(bufferTokens)
}

func shiftBuffer() token.Token {
	token := bufferTokens[0:1][0]
	bufferTokens = bufferTokens[1:]

	return token
}

func lerToken() token.Token {
	return shiftBuffer()
}

func lookahead(k int) token.Token {

	if k-1 > tamanhoBufferTokens-1 {
		return bufferTokens[tamanhoBufferTokens-1]
	}

	return bufferTokens[k-1]
}

func combine(tipoToken string) {
	if tipoToken == lookahead(1).Type {
		lerToken()
		atualizarTamanhoBuffer()
	} else {
		log.Fatal("Erro do tipo sintático, esperando token: ", tipoToken, ", mas foi encontrado um token: ", lookahead(1).Type, lookahead(1))
	}
}

// programa: 'inicio' listaComandos 'fim'
func programa() {
	combine(token.PALAVRA_RESERVADA_INICIO)
	listaComandos()
	combine(token.PALAVRA_RESERVADA_FIM)
}

// listaComandos: comando listaComandosExtras
func listaComandos() {
	comando()
	listaComandosExtras()
}

// listaComandosExtras: anicializacao | atribuicao | declaracao | estruturaCondicional | estruturaDeRepeticao
func listaComandosExtras() {
	if lookahead(1).Type == token.PALAVRA_RESERVADA_DEFINIR ||
		lookahead(1).Type == token.VARIAVEL {
		listaComandos()
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_SE {
		listaComandos()
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_ENQUANTO {
		listaComandos()
	}
}

// comando: anicializacao | atribuicao | declaracao | estruturaCondicional | estruturaDeRepeticao
func comando() {
	if lookahead(5).Type == token.OPERADOR_ATRIBUICAO {
		anicializacao()
	} else if lookahead(2).Type == token.OPERADOR_ATRIBUICAO {
		atribuicao()
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_DEFINIR {
		declaracao()
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_SE {
		estruturaCondicional()
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_ENQUANTO {
		estruturaDeRepeticao()
	} else {
		log.Fatal("Erro do tipo sintático. Comando inválido.")
	}
}

// anicializacao: 'definir' variavel declaracaoDeTipo = tipo
func anicializacao() {
	combine(token.PALAVRA_RESERVADA_DEFINIR)
	combine(token.VARIAVEL)
	combine(token.OPERADOR_DE_TIPAGEM)
	declaracaoDeTipo()
	combine(token.OPERADOR_ATRIBUICAO)
	expressaoAritmetica()
}

// atribuicao: variavel = tipo
func atribuicao() {
	combine(token.VARIAVEL)
	combine(token.OPERADOR_ATRIBUICAO)
	expressaoAritmetica()
}

// declaracao: 'definir' variavel declaracaoDeTipo
func declaracao() {
	combine(token.PALAVRA_RESERVADA_DEFINIR)
	combine(token.VARIAVEL)
	combine(token.OPERADOR_DE_TIPAGEM)
	declaracaoDeTipo()
}

// estruturaCondicional: 'se' expressaoLogica 'entao' listaComandos 'fim'
func estruturaCondicional() {
	combine(token.PALAVRA_RESERVADA_SE)
	expressaoLogica()
	combine(token.PALAVRA_RESERVADA_ENTAO)
	listaComandos()
	combine(token.PALAVRA_RESERVADA_FIM)
}

// estruturaDeRepeticao: 'enquanto' expressaoLogica 'entao' listaComandos 'fim'
func estruturaDeRepeticao() {
	combine(token.PALAVRA_RESERVADA_ENQUANTO)
	expressaoLogica()
	combine(token.PALAVRA_RESERVADA_ENTAO)
	listaComandos()
	combine(token.PALAVRA_RESERVADA_FIM)
}

// expressaoLogica: expressaoAritmetica restoExpressaoLogica
func expressaoLogica() {
	expressaoAritmetica()
	restoExpressaoLogica()
}

// restoExpressaoLogica: operadorRelacional expressaoAritmetica
//					   | operadorRelacional expressaoAritmetica operadorLogico expressaoLogica
func restoExpressaoLogica() {
	if lookahead(3).Type == token.OPERADOR_LOGICO_E ||
		lookahead(3).Type == token.OPERADOR_LOGICO_OU {
		operadorRelacional()
		expressaoAritmetica()
		operadorLogico()
		expressaoLogica()
	} else if lookahead(1).Type == token.OPERADOR_RELACIONAL_IGUALDADE ||
		lookahead(1).Type == token.OPERADOR_RELACIONAL_MAIOR_QUE ||
		lookahead(1).Type == token.OPERADOR_RELACIONAL_MAIOR ||
		lookahead(1).Type == token.OPERADOR_RELACIONAL_MENOR_QUE ||
		lookahead(1).Type == token.OPERADOR_RELACIONAL_MENOR ||
		lookahead(1).Type == token.OPERADOR_RELACIONAL_DIFERENTE {
		operadorRelacional()
		expressaoAritmetica()
	}
}

func operadorLogico() {
	if lookahead(1).Type == token.OPERADOR_LOGICO_E {
		combine(token.OPERADOR_LOGICO_E)
	} else if lookahead(1).Type == token.OPERADOR_LOGICO_OU {
		combine(token.OPERADOR_LOGICO_OU)
	} else {
		log.Fatal("Erro do tipo sintático. operador inválido na linha: ", lookahead(1).LinhaDoToken)
	}
}

// operadorRelacional: == | >= | <= | <> | > | <
func operadorRelacional() {
	if lookahead(1).Type == token.OPERADOR_RELACIONAL_IGUALDADE {
		combine(token.OPERADOR_RELACIONAL_IGUALDADE)
	} else if lookahead(1).Type == token.OPERADOR_RELACIONAL_MAIOR_QUE {
		combine(token.OPERADOR_RELACIONAL_MAIOR_QUE)
	} else if lookahead(1).Type == token.OPERADOR_RELACIONAL_MAIOR {
		combine(token.OPERADOR_RELACIONAL_MAIOR)
	} else if lookahead(1).Type == token.OPERADOR_RELACIONAL_MENOR_QUE {
		combine(token.OPERADOR_RELACIONAL_MENOR_QUE)
	} else if lookahead(1).Type == token.OPERADOR_RELACIONAL_MENOR {
		combine(token.OPERADOR_RELACIONAL_MENOR)
	} else if lookahead(1).Type == token.OPERADOR_RELACIONAL_DIFERENTE {
		combine(token.OPERADOR_RELACIONAL_DIFERENTE)
	} else {
		log.Fatal("Erro do tipo sintático. operador inválido na linha: ", lookahead(1).LinhaDoToken)
	}
}

// expressaoAritmetica: fator restoExpressaoAritmetica
func expressaoAritmetica() {
	fator()
	restoExpressaoAritmetica()
}

// restoExpressaoAritmetica: operadorAritimetico fator
// 						   | operadorAritimetico fator restoExpressaoAritmetica
func restoExpressaoAritmetica() {
	if lookahead(1).Type == token.OPERADOR_ARITMETICO_SOMA ||
		lookahead(1).Type == token.OPERADOR_ARITMETICO_SUBTRACAO ||
		lookahead(1).Type == token.OPERADOR_ARITMETICO_MULTIPLICACAO ||
		lookahead(1).Type == token.OPERADOR_ARITMETICO_DIVISAO {
		operadorAritimetico()
		fator()
		restoExpressaoAritmetica()
	}
}

// operador:  '+' | '-' | '*' | '/'
func operadorAritimetico() {
	if lookahead(1).Type == token.OPERADOR_ARITMETICO_SOMA {
		combine(token.OPERADOR_ARITMETICO_SOMA)
	} else if lookahead(1).Type == token.OPERADOR_ARITMETICO_SUBTRACAO {
		combine(token.OPERADOR_ARITMETICO_SUBTRACAO)
	} else if lookahead(1).Type == token.OPERADOR_ARITMETICO_MULTIPLICACAO {
		combine(token.OPERADOR_ARITMETICO_MULTIPLICACAO)
	} else if lookahead(1).Type == token.OPERADOR_ARITMETICO_DIVISAO {
		combine(token.OPERADOR_ARITMETICO_DIVISAO)
	} else {
		log.Fatal("Erro do tipo sintático. operador inválido na linha: ", lookahead(1).LinhaDoToken)
	}
}

// declaracaoDeTipo: 'inteiro' | 'real' | 'cadeia' | 'logico'
func declaracaoDeTipo() {
	if lookahead(1).Type == token.PALAVRA_RESERVADA_INTEIRO {
		combine(token.PALAVRA_RESERVADA_INTEIRO)
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_REAL {
		combine(token.PALAVRA_RESERVADA_REAL)
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_CADEIA {
		combine(token.PALAVRA_RESERVADA_CADEIA)
	} else if lookahead(1).Type == token.PALAVRA_RESERVADA_LOGICO {
		combine(token.PALAVRA_RESERVADA_LOGICO)
	} else {
		log.Fatal("Erro do tipo sintático. Tipo declarado inválido na linha: ", lookahead(1).LinhaDoToken)
	}
}

// fator: tipo | '('expressaoAritmetica')' | variável
func fator() {
	if lookahead(1).Type == token.ABRE_PARENTESE {
		combine(token.ABRE_PARENTESE)
		expressaoAritmetica()
		combine(token.FECHA_PARENTESE)
	} else if lookahead(1).Type == token.TIPO_NUMERO_INTEIRO ||
		lookahead(1).Type == token.TIPO_NUMERO_REAL ||
		lookahead(1).Type == token.TIPO_CADEIA ||
		lookahead(1).Type == token.TIPO_LOGICO {
		tipo()
	} else if lookahead(1).Type == token.VARIAVEL {
		combine(token.VARIAVEL)
	} else {
		log.Fatal("Erro do tipo sintático, operação aritimetica inválida. linha: ", lookahead(1).LinhaDoToken)
	}
}

// tipo: 132 | 1.1 | "string" | true
func tipo() {
	if lookahead(1).Type == token.TIPO_NUMERO_INTEIRO {
		combine(token.TIPO_NUMERO_INTEIRO)
	} else if lookahead(1).Type == token.TIPO_NUMERO_REAL {
		combine(token.TIPO_NUMERO_REAL)
	} else if lookahead(1).Type == token.TIPO_CADEIA {
		combine(token.TIPO_CADEIA)
	} else if lookahead(1).Type == token.TIPO_LOGICO {
		combine(token.TIPO_LOGICO)
	} else {
		log.Fatal("Erro do tipo sintático. Estavamos esperando tipo válido. linha: ", lookahead(1).LinhaDoToken)
	}
}
