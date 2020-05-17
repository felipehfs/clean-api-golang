# Testabilidade em Go

## Caso de uso - Exemplo

A API hipotética deverá gerenciar seu acervo pessoal de livros. Ela deve exigir uma autenticação para excluir e atualizar os livros e o restante das operações podem ser públicas.

## Porque Go?

- Go é uma linguagem de fácil aprendizado.
- Go é minimalista.  
- A linguagem Go oferece uma suíte poderosa e simples por padrão para a criação de testes. 
- Go oferece performance sem sacrificar produtividade.


## A criação de Testes em Go

Por critério de organização, em nossos testes colocamos o _test na declaração do pacote, para definir um escopo separado de nossas variáveis globais. 
Podemos configurar e inicializar as nossas depêndecias usando o testmain do Go.

```go
func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    tearDown()
    os.Exit(code)
}
```

Para criar um teste em Go apenas escrevemos Test em maiúsculo seguindo com um nome daquilo que precisa ser testado.
Você pode rodar o teste de maneira recursiva no diretório, usando:
 
```bash
    go test  -v ./... 
```

Existe uma flag _-cover_ que mostra a cobertura de teste no pacote.

## Tabela de Teste

A api fornece suporte a subtestes que é uma excelente maneira de torná seu teste de unidade 
mais robusto. 

```go
func TestSum(t *testing) {
    testCases := []struct{
        Description string
        Number      int 
        Expected    int
    }{
        {"O dobro de 2 deve retornar 4", 2, 4}
    }
    
    for _, testCase := range testCases {
        t.Run(tesCase.Description, func(t *testing) {
            product := twice(testCase.Number)
            if testCase.Expected != product {
                t.Errorf("Expected %v but got %v", testCase.Expected, product)
            }
        }
    }
}
```

## Uso de Examples

O *Example* pode ser utilizado para documentar e também testar a função.

```go
    func ExampleExportCSV() {
        books := []entities.Book{
            {ID: 1, Name: "Pequeno Príncipe", Price: 140.30, ISBN: "RERAIA-EIRURJGM-QQIW"},
        }
        presenters.ExportBookToCSV(books, os.Stdout)
        // output:
        // 1,Pequeno Príncipe,RERAIA-EIRURJGM-QQIW,140.3    
    }
```

## A escolha da arquitetura e como ela influência no teste

Nessa API escolhi utilizar alguns princípios da Arquitetura limpa. Na Arquitetura limpa a divisão fica assim:

- Entities -  O pacote tem as nossas entidades. 
- Repositories - Nossos objetos que vão abstrair o nosso SQL e as particularidades do Banco de dados.
- Usecases -  Possui nossa a lógica de negócio da aplicação.
- Presenters - A interface do cliente com a nossa aplicação, possui os controllers que irão gerenciar a forma que o HTTP irá apresentar para o cliente.

Então, cada camada pode ser testada de forma isolada, e simulada utilizando os princípios de injeção de depedência. Podemos testar toda nossa lógica de negócio 
em usecases e em presenters o status code e a lógica presente no controller.

## Middleware em Go

A linguagem Go é multiparadigma e permite que você utilize *high order function* o que
significa que você pode utilizar uma função como argumento e como retorno de função.

```go

func greet(user User) {
    fmt.Printf("Olá %v", user.Name)
}

func main() {
    pubSub := pubsub.New()
    pubSub.on("new_user", greet)
}

```

Na criação de middleware em Go podemos utilizar o *http.HandlerFunc*. 
```go

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
Ele é um tipo que na tradução literal permite que funções comuns com a sua
assitura implementem a interface *http.Handler*.

Então a escrita do middleware em Go fica assim. 

```go
func hasLogger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host + r.URL.Path)
		handler(w, r)
	}
}
```
Ele tem um funcionamento parecido com um decorator mesmo. Então baseando nisso, nós podemos criar também nosso prório tipo chamado Middleware.

```go
type Middleware func(http.HandlerFunc) http.HandlerFunc

```
E extrapolando um pouco mais podemos juntar todos os nosso middlewares no controller fazendo esta função.

```go
func decorate(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
```

## Referências bibliográficas

- https://golang.org/pkg/testing/ 
- [Golang UK Conference 2015 - Mat Ryer - Building APIs](https://www.youtube.com/watch?v=tIm8UkSf6RA)