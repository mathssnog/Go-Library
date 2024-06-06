# Biblioteca
- struct livro
- struct usuario
- criar api server
- criar endpoint listar usuarios
- criar endpoint listar livros
- criar endpoint pegar user pelo ID
- criar endpoint de cadastro de usuário
- criar endpoint de cadastro de usuário
- criar endpoint de atualizar usuário por id
- criar endpoint de atualizar livro por id
- criar endpoint deletar usuario
- criar endpoint deletar livro

Aplicação para gerenciamento de bibliotecas. Considerações:

Sistema de gerenciamento de uma biblioteca que permite a organização e controle de livros, membros, empréstimos, devoluções e reservas. O sistema será implementado em Python, utilizando o framework Flask e PostgreSQL como banco de dados. O sistema deve:

1. **Cadastrar Livros**:
    Para o cadastro de livros vamos ter algumas variáveis como título do livro, autor(es), número de cópias disponíveis, editora e ano de publicação.

2. **Cadastrar Membros**:
    Para o cadastro de membros vamos ter algumas variáveis como nome, ID e informações pessoais.
   
3. **Empréstimos**:
    Para o empréstimos de livros vamos ter variáveis como registro do membro que solicitou o empréstimo e sua data.

4. **Devoluções**:
    Para a devolução, será registrado o membro que realizou a devolução e sua data. Vale lembrar que pode-se aplicar multas à entregas atrasadas.

5. **Reservas**:
    Esta seção seria para caso acontecesse uma grande demanda de um mesmo livro, onde será necessário uma fila com base na ordem de solicitação do livro.

6. **Pesquisa e Filtros**:
    Plataforma capaz de pesquisar os livros por título, autor(es) para verificar a disponibilização do mesmo.

7. **Autenticação e Controle de Acesso**:
    Criar acessos administrativos para ações administrativas.

8. **Notificações**:
    Notificação para os membros sobre datas de devolução e disponibilidade de reservas caso solicitado.

9. **Multas e Penalidades**:
    Seção destinada para o cálculo da multa cobrada por atraso de devolução.

10. **Gerenciamento de Estoque**:
    Seção destinada a atualizar o número de cópias de um livro quando este sofrer empréstimo e devoluções.

11. **Interface de Usuário (frontend**).
    Uma interface de usuário amigável para administradores da biblioteca e membros.

**PostgreSQL**

O banco de dados utilizado neste projeto será o PostgreSQL:

- Instalação do PostgreSQL em [postgresql.org](https://www.postgresql.org/).
    - Verificação do status do postgres, após a instação
        - sudo systemctl status postgresql

    - Podemos conectar ao postgreSQL utilizamos o seguinte comando:
        - sudo -u postgres psql
    
    - Após a conexão com o postgreSQL, podemos realizar o seguinte comando:
        - <comando para criar o usuário e db>
        - CREATE DATABASE livros; 
            - Verifica-se em \l (list of db)
        - CREATE USER livros_admin WITH PASSWORD '123';
        - ALTER ROLE livros_admin SET client_encoding TO 'utf8';
        - ALTER ROLE livros_admin SET timezone TO 'UTC';
        - GRANT ALL PRIVILEGES ON DATABASE livros TO livros_admin;

        Obs.: Alteração no arquivo: /etc/postgresql/14/main/pg_hba.conf
        Para autenticação por senha para todas as conexões locais   -> mudança do tipo peer para md5.

**Linguagem**

- A aplicação será desenvolvido em Golang.

