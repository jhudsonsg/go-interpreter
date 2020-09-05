## 📕 Gramática

### 🔑 Palavras reservadas
- inicio 
- fim
- entao
- se
- enquanto
- definir
- inteiro
- real
- cadeia
- logico
- verdadeiro
- falso
- e
- ou

### 📐 Escrita
- programa                 : 'inicio' listaComandos 'fim'
- listaComandos            : comando | comando listaComandos
- comando                  : anicializacao | atribuicao | declaracao | controladorDeFluxo | lacoDeRepeticao
- anicializacao            : 'definir' variável declaracaoDeTipo = tipo
- atribuicao               : variavel = tipo
- declaracao               : 'definir' variavel declaracaoDeTipo
- estruturaCondicional     : 'se' expressaoLogica 'entao' listaComandos 'fim'listaComandos 'fim'
- estruturaDeRepeticao     : 'enquanto' expressaoLogica 'entao' listaComandos 'fim'
- expressaoLogica          : expressaoAritmetica restoExpressaoLogica
- restoExpressaoLogica     : operadorRelacional expressaoAritmetica | operadorRelacional expressaoAritmetica operadorLogico expressaoLogica
- operadorRelacional       : == | >= | <= | <> | > | <
- expressaoAritmetica      : fator restoExpressaoAritmetica
- restoExpressaoAritmetica : operadorAritimetico fator | operadorAritimetico fator restoExpressaoAritmetica
- operadorAritimetico      :  '+' | '-' | '*' | '/'
- declaracaoDeTipo         : 'inteiro' | 'real' | 'cadeia' | 'logico'
- fator                    : tipo | '('expressaoAritmetica')' | variável
- tipo                     : 132 | 1.1 | "string" | true

### 💎 Tokens
Token                               | Exemplo
----------------------------------- | ------------------
TIPO_NUMERO_INTEIRO                 | 1
TIPO_NUMERO_REAL                    | 1.0
TIPO_CADEIA                         | 'cadeia'
TIPO_LOGICO                         | verdadeiro | falso
OPERADOR_LOGICO_E                   | e
OPERADOR_LOGICO_OU                  | ou
OPERADOR_ATRIBUICAO                 | =
OPERADOR_LOGICO_IGUALDADE           | ==
OPERADOR_LOGICO_MAIOR_QUE           | >=
OPERADOR_MAIOR                      | >
OPERADOR_LOGICO_MENOR_QUE           | <=
OPERADOR_LOGICO_MENOR               | <
OPERADOR_LOGICO_DIFERENTE           | <>
OPERADOR_ARITMETICO_SOMA            | +
OPERADOR_ARITMETICO_SUBTRACAO       | -
OPERADOR_ARITMETICO_MULTIPLICACAO   | *
OPERADOR_ARITMETICO_DIVISAO         | /
PALAVRA_RESERVADA_INICIO            | inicio
PALAVRA_RESERVADA_FIM               | fim
PALAVRA_RESERVADA_ENTAO             | entao
PALAVRA_RESERVADA_SE                | se
PALAVRA_RESERVADA_ENQUANTO          | enquanto
PALAVRA_RESERVADA_DEFINIR           | definir
PALAVRA_RESERVADA_INTEIRO           | inteiro
PALAVRA_RESERVADA_REAL              | real
PALAVRA_RESERVADA_CADEIA            | cadeia
PALAVRA_RESERVADA_LOGICO            | logico
ABRE_CHAVE                          | {
FECHA_CHAVE                         | }
ABRE_PARENTESE                      | (
FECHA_PARENTESE                     | )
VARIAVEL                            | variavel