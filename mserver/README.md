# Pacote mserver

O objetivo do pacote mserver é proporcionar abstrações, por exemplo na configuração das respostas de seus serviços. Logo em seguida, teremos uma demonstração de sua implementação.

## Vamos nos divertir!

Vou passar aqui como fazer para testar o pacote, mas a implementação fica a critério da pessoa que estiver desenvolvendo. No main vamos incluir antes alguns recursos que vou explicar separadamente.

### Adicionando constants

```
const (
	TESTE1 mserver.KeyResponse = "TESTE1"
	TESTE2 mserver.KeyResponse = "TESTE2"
)
```

As constants adicionadas em seu main como indicado acima, se trata de chaves que seram utilizadas para indicar a implementação em sua fábrica de Response. Esse assunto será abordado mais a frente na leitura.

### Adicionando estruturas locais

```
type MResponse struct {
	Cod int    `json:"cod"`
	Msg string `json"msg"`
}

type MainServer struct {
}

// Irá implementar AssemblyResponse() e adicionar a instância da Fábrica de Response
type Retorno1 struct {
}

// Irá implementar AssemblyResponse() e adicionar a instância da Fábrica de Response
type Retorno2 struct {
}

```

Essas são estruturas que possibilitará a disposição de um ou mais objetos Response para cada serviço. E implementar interfaces, possibilitando que você siga o contrato definido no pacote.

### Adicionando variáveis 

```
// Disponibiliza a instância de uma estrutura local. Com a finalidade de escrever a implementação das interfaces
// BadMethod(r *http.Request, verb mserver.MVerb) (bool, interface{}) : Implementa a verificação do Method processado naquela solicitaçao
// Response(keyResponse mserver.KeyResponse) interface{} : Implementa o retorno para aquela solicitação, obtendo a interface da fábrica de response
var iMainServer *MainServer

// Disponibiliza uma fábrica de response 
var facResponse *mserver.FacResponse
```

### Vamos ver as implementações!

```
//
func (mainServer *MainServer) BadMethod(r *http.Request, verb mserver.MVerb) (bool, interface{}) {

	if r.Method != string(verb) {
		var mResponse MResponse = MResponse{Cod: 403, Msg: "Solicitação não é GET"}
		return true, mResponse
	}

	return false, nil
}

func (mainServer *MainServer) Response(keyResponse mserver.KeyResponse) interface{} {
	implRet := facResponse.GetImpl(keyResponse)
	return implRet.AssembyResponse()
}

```

A implementação de BadMethod, gostei desse nome, me perdoe. Recebe uma instância de request e um type MVerb que faz parte do pacote. Quando ocorre uma chamada  pRequest.ValidMethod(mserver.GET), que estará em uma explicação logo a frente. Ocorre a execução da implementação. 