# Proposta

Gostaríamos de criar um jogo de Star Wars com algumas informações da franquia. E para possibilitar a
equipe de frontend criar essa aplicação, precisamos desenvolver uma API REST que contenha os dados dos
planetas da franquia.

# Escopo
Para cada planeta, os seguintes dados devem ser obtidos do banco de dados da aplicação, inseridos a partir
de requisições disparadas para a API pública do Star Wars:
- Nome, clima e terreno;
- Para cada planeta também devemos ter os filmes com o nome, diretor e data de lançamento;
  
Todas as informações necessárias podem ser obtidas pela API pública do [Star Wars](https://swapi.dev/).
- Funcionalidades desejadas:
- Carregar um planeta da API através do Id
- Listar os planetas
- Buscar planeta por nome
- Buscar por ID,
- Remover planeta

# Requisitos
- utilize git ou hg para fazer o controle de versão da solução do teste e hospede-a no Github ou Bitbucket;
- armazene os dados no banco de dados que você julgar apropriado;
- a API deve seguir os conceitos REST;
- o Content-Type das respostas da API deve ser application/json ;
- o código da solução deve conter testes e algum mecanismo documentado para - gerar a informação de
cobertura dos testes;
- a aplicação deve gravar logs estruturados em arquivo texto;

# Desenho da Solução