# xp task dealer

## O que é

O projeto propõe um distribuidor de tarefas para facilitar a escolha de User Stories na metodologia de Programação Extrema (XP). 

Utilizando Grandes Modelos de Linguagem (LLM), o projeto sugere tarefas alinhadas com as habilidades e histórico dos desenvolvedores.
Isso promove a diversidade de habilidades e evita acomodações, seguindo os princípios ágeis de valorizar indivíduos e encorajar experimentações para melhorar os processos de desenvolvimento de software.

### Como funciona

1. As tarefas são adicionadas com uma descrição dos seus requisitos
2. Os desenvolvedores são adicionados com uma descrição das suas habilidades, histórico de tarefas, e qualquer informação que possa ser útil para descrever a pessoa
3. É possível acessar a página inicial da aplicação para escolher entre tarefas ou desenvolvedores, e solicitar uma sugestão
4. Após analisar a sugestão, é possível aceitar a sugestão para que aquele desenvolvedor não seja mais sugerido em outras tarefas, assim como a tarefa não aparecer para outros desenvolvedores
5. Também é possível recusar uma sugestão para que a aplicação faça uma nova sugestão evitando a anterior

## Como iniciar o serviço

1. [Instale o Docker](https://docs.docker.com/engine/install/)
2. Gere uma chave secreta na OpenAI e coloque em um arquivo `.env` no diretório raíz do projeto com a chave `OPENAI_KEY`
3. Rode o comando `docker-compose up`
4. Visite o link https://localhost:8080/

## Organização do projeto

```
├── api_collections
├── core
│   ├── in_memory
│   ├── models
│   ├── openai_dealer
│   └── sqlite_store
├── server
│   └── dto
└── static
```

`root`: arquivos de definição do projeto, como os arquivos para executar através do Docker, o arquivo de exemplo da `.env` necessária e a versão do Go para rodar o projeto

`api_collections`: contém as rotas para requisitar a API do projeto através da [IDE Bruno](https://github.com/usebruno/bruno/)

`core`: contém as interfaces e suas implementações para rodar o projeto

`core/in_memory`: implementação da interface `Storer` para facilitar o desenvolvimento dos testes

`core/models`: definição das estruturas utilizadas na aplicação

`core/openai_dealer`: implementação da interface `Dealer` utilizando a LLM da OpenAI

`core/sqlite_store`: implementação da interface `Storer` utilizando o SQLite para persistir os desenvolvedores e tarefas disponíveis

`core/service.go`: estrutura que prove as funções para acessar os casos de uso do projeto

`server`: implementação da apresentação do projeto utilizando o protocolo HTTP para acessar os casos de uso, e uma página HTML para utilizar o projeto pelo navegador

`server/dto`: funções para mapear as estruturas entre a camada de apresentação e a camada de serviço

`static`: arquivos para apresentar a página web