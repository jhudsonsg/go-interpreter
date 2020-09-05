// Pacote responsável por gerar tokens de um determinado arquivo.
package main

import (
	"log"

	"github.com/jhudsonsg/interpreter/reconhecer"
	"github.com/jhudsonsg/interpreter/token"
)

// proximoToken - lê varios caracteres para gerar um token reconhecido pela linguagem.
func proximoToken() token.Token {
	for caractere := lerProximoCaractere(); caractere != ""; caractere = lerProximoCaractere() {

		if reconhecerComentarios(caractere) {
			continue
		}

		if reconhecerEspacos(caractere) {
			continue
		}

		if token := reconhecerOperador(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerOperadorLogicos(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerOperadorAritimetica(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerLimitadores(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerNumeros(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerCadeias(caractere); token.Type != "" {
			return token
		}

		if token := reconhecerCaracteres(caractere); token.Type != "" {
			return token
		}

		log.Fatal("Erro do tipo léxico, Token: ", caractere, ", não conhecido. Linha: ", linhaAtual)
	}

	return token.Token{}
}

// reconhecerComentarios - pula linha com comentários.
func reconhecerComentarios(caractere string) bool {
	if caractere == "/" {
		caractere = lerProximoCaractere()

		if caractere == "/" {
			for caractere := lerProximoCaractere(); caractere != "\n"; caractere = lerProximoCaractere() {
			}

			avancarLinha()
			return true
		}
	}

	return false
}

// reconhecerEspacos - pula linhas com espaços em brancos ou quebra de linha
func reconhecerEspacos(caractere string) bool {
	if caractere == " " || caractere == "\r" {
		return true
	}

	if caractere == "\n" {
		avancarLinha()
		return true
	}

	return false
}

// reconhecerOperador - reconhece o caractere : paara indicar o tipo de uma variável
func reconhecerOperador(caractere string) token.Token {
	if caractere == ":" {
		return token.Token{Type: token.OPERADOR_DE_TIPAGEM, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}

// reconhecerOperador - reconhece tokens de operadores lógicos.
func reconhecerOperadorLogicos(caractere string) token.Token {
	if caractere == "=" {
		lexema := caractere
		caractere = lerProximoCaractere()

		if caractere == "=" {
			lexema += caractere
			return token.Token{Type: token.OPERADOR_RELACIONAL_IGUALDADE, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		RetrocederNoBuffer()
		return token.Token{Type: token.OPERADOR_ATRIBUICAO, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	if caractere == ">" {
		lexema := caractere
		caractere = lerProximoCaractere()

		if caractere == "=" {
			lexema += caractere
			return token.Token{Type: token.OPERADOR_RELACIONAL_MAIOR_QUE, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		RetrocederNoBuffer()
		return token.Token{Type: token.OPERADOR_RELACIONAL_MAIOR, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	if caractere == "<" {
		lexema := caractere
		caractere = lerProximoCaractere()

		if caractere == "=" {
			lexema += caractere
			return token.Token{Type: token.OPERADOR_RELACIONAL_MENOR_QUE, Lexema: lexema, LinhaDoToken: linhaAtual}
		} else if caractere == ">" {
			lexema += caractere
			return token.Token{Type: token.OPERADOR_RELACIONAL_DIFERENTE, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		RetrocederNoBuffer()
		return token.Token{Type: token.OPERADOR_RELACIONAL_MENOR, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}

// reconhecerOperadorAritimetica - reconhece tokens de operadores aritimeticos.
func reconhecerOperadorAritimetica(caractere string) token.Token {
	if caractere == "+" {
		return token.Token{Type: token.OPERADOR_ARITMETICO_SOMA, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "-" {
		return token.Token{Type: token.OPERADOR_ARITMETICO_SUBTRACAO, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "*" {
		return token.Token{Type: token.OPERADOR_ARITMETICO_MULTIPLICACAO, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "/" {
		return token.Token{Type: token.OPERADOR_ARITMETICO_DIVISAO, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}

// reconhecerLimitadores - reconhece tokens de limitadores.
func reconhecerLimitadores(caractere string) token.Token {
	if caractere == "{" {
		return token.Token{Type: token.ABRE_CHAVE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "}" {
		return token.Token{Type: token.FECHA_CHAVE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == "(" {
		return token.Token{Type: token.ABRE_PARENTESE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	if caractere == ")" {
		return token.Token{Type: token.FECHA_PARENTESE, Lexema: caractere, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}

// reconhecerCaracteres - reconhe tokens de caracteres, distinguindo-os palavras chaves de variáveis.
func reconhecerCaracteres(caractere string) token.Token {
	if reconhecer.HeCaractere(caractere) {
		lexema := caractere

		for caractere := lerProximoCaractere(); reconhecer.HeCaractere(caractere) || reconhecer.HeNumero(caractere); caractere = lerProximoCaractere() {
			lexema += caractere
		}

		RetrocederNoBuffer()

		if hePalavraReservada, tipoToken := reconhecer.HePalavraReservada(lexema); hePalavraReservada {
			return token.Token{Type: tipoToken, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		return token.Token{Type: token.VARIAVEL, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}

// reconhecerNumeros - reconhe tokens de sequ^qncia de números, distinguindo-os de reais e inteiros.
func reconhecerNumeros(caractere string) token.Token {
	if reconhecer.HeNumero(caractere) {
		lexema := caractere

		for caractere := lerProximoCaractere(); reconhecer.HeNumero(caractere) || caractere == "."; caractere = lerProximoCaractere() {
			lexema += caractere
		}

		RetrocederNoBuffer()

		if reconhecer.HeReal(lexema) {
			return token.Token{Type: token.TIPO_NUMERO_REAL, Lexema: lexema, LinhaDoToken: linhaAtual}
		}

		return token.Token{Type: token.TIPO_NUMERO_INTEIRO, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}

// reconhecerCadeias - reconhece tokens de cadeira com inicio de uma "'".
func reconhecerCadeias(caractere string) token.Token {
	if caractere == "'" {
		lexema := caractere

		for caractere := lerProximoCaractere(); caractere != "'"; caractere = lerProximoCaractere() {
			lexema += caractere
		}
		lexema += "'"

		return token.Token{Type: token.TIPO_CADEIA, Lexema: lexema, LinhaDoToken: linhaAtual}
	}

	return token.Token{}
}
