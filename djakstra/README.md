# Notas Sobre o Algoritmo de Djakstra

Essa é uma implementação mais simples do algoritmo que é apresentada no livro, porém adicionei algumas coisas a mais, eu achei que montar os dados de `cost` para o algoritmo seria muito ineficiente, então criei uma função chamada `NewCosts` que monta os dados de custo iniciais, porém pra isso é necessário conhecer qual é o nó inicial do grafo, uma vez que mapas não necessáriamente mantem a ordem e que não necessariamente precisemos inserir no mapa os grafos na ordem correta, uma função que localiza o inicio do grafo, chamada `GetStart` foi criada.

Após isso o fluxo do algoritmo é esse

1. Localizar, na tabela de custos, o menor custo:

Na primeira iteração a tabela de custos vai ter só os nós que estão ligados ao nó de start, por exemplo:

```
               ┌─┐         
        ┌──2──►│A├──1───┐  
        │      └─┘      │  
    ┌─────┐           ┌─▼─┐
    │Start│           │End│
    └───┬─┘           └─▲─┘
        │      ┌─┐      │  
        └──1──►│B├──3───┘  
               └─┘         
```

Nesse caso essa seria a tabela de custos que seria usada na primeira iteração, isso é, antes do for

| Node  | Custo |
|-------|-------|
| Start | 0     |
| A     | 2     |
| B     | 1     |


Teriamos apenas os custos dos nós partindo do *start*, que são os nós conhecidos. Eles são inseridos na tabela de custos pela função `NewCosts`

A função responsável por buscar o nó de menor custos na tabela de custos é a `GetLower`, éla retorna o nome do nó de menor custo, nesse caso ela deve retornar primeiramente o B, e que ainda não tenha sido processado, ai entra em sena um `map[any]bool` que usamos para marcar os nós já processados 

Após isso o algoritmo entra no primeiro for, que vai repetir até a função `GetLower` retornar um `nil`

2. Checar os vizinhos do menor nó:

Agora que temos o menor nó, neste caso o **B**  devemos buscar os vizinhos de **B**, lembrando que aqui entramos em outro for, para iterar sobre os vizinhos de **B**, que será, segundo o grafo acima, o **End** com custo 3, e então nós calculamos o custo pra chegar ao **End** por esse caminho, que se dá pela soma de 1 Start -> B com 3 B -> End, resultando em um 4

3. Atualizar a tabela de custos

Com esse resultado nós voltamos para a tabela de custos e procuramos para o **End**, que é o nó atual, qual é o custo que estava lá, e caso o custo que estava lá seja maior que o que calculamos ou o custo de **End** não exista na tabela de custos, então inserimos na tabela de custos com o custo resultante da soma realizada acima, no nosso caso o custo para **End** ainda não existia, e após a inserção a tabela de custos deve ficar assim

| Node  | Custo |
|-------|-------|
| Start | 0     |
| A     | 2     |
| B     | 1     |
| End   | 4     |

> Lembrando que o custo para **End** é 4 devido a ser cumulativo de Start -> B + B -> End

Agora, no final do primeiro for, devemos inserir o nó **B** na lista de nós já processados `Processed` e buscar novamente o menor custo na tabela de custos usando a função `GetLower`, lembrando que como o nó **B** já foi processado, ele não será retornado pela função e então será retornado o A, dai repetios o processo.

Nesta segunda rodada, quando estivermos analizando os vizinhos de **A** vamos encontrar o nó final **End**, vamos calcular o custo para **End** porém através do nó **A**, e desta vez, o **End** já existirá na tabela de custos porém desta vez o valor obtido através de **A** será menor do que o valor existente pois start -> A custa 2 e A -> End custa 1, totalizando 3 que é menor do que o 4 que está na tabela

| Node  | Custo |
|-------|-------|
| Start | 0     |
| A     | 2     |
| B     | 1     |
| End   | **3** |

Mais uma vez, marcamos o **A** como vizitado, agora vamos buscar novamente o menor custo na tabela de custos que ainda não foi processado, vamos fazer isso atravez da função `GetLower` que deverá retornar o **End**, porém ele não tem Vizinhos, sendo assim marcamos ele como vizitado e chamamos a `GetLower` novamente, que retornará `nil` pois não tem mais nós que não foram vizitados.

Assim se encerra o algoritmo.
