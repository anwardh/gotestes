# desafio-cierre-testing
Proyecto Base para cumplir el Desafio de Testing

### Cenários de teste

#### Handler

* Deve retornar status code 400 se o query param seller_id não for passado

* Deve retornar status code 404 se o seller não existir

* Deve retornar status code 204 se o seller não tiver produtos

* Deve retornar status code 500, caso o método GetAllBySeller da repository retorne um erro

* Deve retornar status code 200, com os produtos como resposta

#### Service

* Deve retornar os sellers ao chamar a repository 
* Deve retornar o erro caso ocorra um erro na chamada da repository
