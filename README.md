# Transparency
## Sistema para consulta de instituições de ensino

**Pacotes de terceiros**
* Nome: pq | Código para instalação: `go get github.com/lib/pq` | [Repositório GitHub](https://github.com/lib/pq "Repositório GitHub") 

**Configurações
* Linha 38 do arquivo main.go, setar os parâmetros para conexão ao banco de dados.
* Executar as querys do arquivo database/database.sql, não esqueça de selecionar o database correto que é `transparency`
* Quanto for executar a query de importação dos dados, linha 40 do arquivo database/database.sql, favor alterar o diretório do arquivo csv para o seu diretório atual.

** Build do app
* Entrar no diretório bin
* Executar o comando `go build main.go`, será gerado o main.exe.
* Execute o main.exe via cmd `main`
* Acesse a url localhost:8000 pelo seu navegador favorito, e pronto, app em execução ;) 