### Comandos

Comando para verificar a cobertura:

```sh
go test -cover ./...
```

Comando para gerar o relatório:

```sh
# o texto "coverage.out" é o nome do arquivo que será gerado, pode ser outro também
go test -cover -coverprofile=coverage.out  ./...
```

Comando para ver o relatório em uma página html

```sh
# o texto "cov.out" é o nome do arquivo que gerou anteriormente
 go tool cover -html=coverage.out
```
