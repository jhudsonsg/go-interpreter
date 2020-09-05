# 🤖 go-interpreter 
go-interpreter é um compilador escrito totalmente em português e na linguagem Go, feito com o propósito didático e acadêmicos para facilitar a leitura e o entendimento de como um compilador funciona.

# 🔨 Como usar
Para poder usar a aplicação bastas você executar o seguinte comando(Você pode modificar os input de entrada no arquivo ./file que está na raiz do projeto).

```bash
make // execute dentro da raiz para buildar e executar o código
```

# ✏️ Sobre a linguagem

### 📌 Extras
Foi escrito um analisador léxico próprio para esse frontend, segue o link do projeto, https://github.com/jhudsonsg/go-scanner.

### 💅 Tipos primitivos
Tipo      | Status
--------- | ------
Real      | ✔️
Inteiro   | ✔️
Cadeia    | ✔️
Logico    | ✔️

### 💥 Exemplos de código
Vamos da uma exemplo trivial e inútil para você se familiarizar com a escrita.

```golang
inicio // indica ao compilador onde é o início do programa.
    definir n1: inteiro // Declara uma váriavel do tipo inteiro chamado n1.
    definir n2: inteiro = 10 // Inicialização de uma váriavel do tipo inteiro chamado n2.
    definir str1: cadeia = 'Uma palavra' // Inicialização uma váriavel do tipo cadeia, observe que usamos (') para indicar.

    n1 = 0 // Atribui uma um valor á váriavel n1.
    n2 = 2 // Atribui uma um valor á váriavel n2.

    // Uma estrutura condicional com uma expressão lógica simples.
    se n2 > n1
    entao
        n1 = 1 // faz uma atribuição a n1
    fim

    // Uma estrurua de repetição com uma expressão lógica simples.
    enquanto n1 > n2
    entao
        n1 = n1 +n2
    fim
fim // indica ao compilador onde é o fim do programa.
```

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