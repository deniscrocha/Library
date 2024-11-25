# Database
## Index
**[ENGLISH](#en)**
   - [How to Setup the Database](#how-to-setup-the-database)
   - [Database Structure](#database-structure)
   
**[PORTUGUÊS](#pt-br)**
   - [Como Configurar o Banco de Dados](#como-configurar-o-banco-de-dados)
   - [Estrutura do Banco de Dados](#estrutura-do-banco-de-dados)

## EN

### How to Setup the Database

You can set up the database in two ways: using Docker or configuring your own MySQL or other data source.

1. **Docker**

To use Docker, first, configure the `docker-compose.yml` file as desired, or you can use the default configurations. Then, execute the following commands in the terminal inside the `db` folder.

- `docker-compose create`: This command will generate Docker configurations on your machine.
- `docker-compose start`: This command will start the Docker image with MySQL on your machine.
- `docker-compose exec -p mysql-db mysql -u root -p`: This command will enter MySQL and prompt for the administrator password configured in the `docker-compose.yml` file.

Note: If you change the image name, you need to use the new name instead of "mysql-db".

Inside MySQL, you need to use the database:
- `USE library;`: This command selects the database created in the Docker image.
- `SOURCE script.sql`: This command will generate the tables in the database.

After completing these steps, the database will be set up.

2. **Other Sources**

You should configure the database according to the settings that will be used in the backend. Once the configuration is correct, run the `script.sql` file in the new database, using the chosen database.

### Database Structure

Here, I explain each table in the database.

- **Genders**:  
  This table stores book genres.  
  - **Id**: Unique identifier for the category.  
  - **Name**: The name of the category, such as "drama".  
  - **Description**: A brief description of books in the genre.

- **BooksModels**:  
  This table stores book models.  
  - **Id**: Unique identifier for the book.  
  - **Name**: The book's name.  
  - **Release_Date**: The date the book was released.  
  - **Author**: The author of the book.  
  - **Language**: The language of the book.  
  - **Quantity**: The number of copies of the book available in the bookstore.  
  - **Series**: The series to which the book belongs, if applicable.

- **BooksGenders**:  
  This is a relationship table between books and genres. It stores the genres of the books.  
  - **Id**: Unique identifier for the relationship.  
  - **Book_Model**: The identifier (id) of the book.  
  - **Gender**: The identifier (id) of the book's genre.

- **Books**:  
  This table stores the copies of books that the bookstore owns.  
  - **Id**: Unique identifier for the book copy.  
  - **Book_Model**: The identifier for the book model.  
  - **Status**: The status of the book (whether it is loaned, available, or damaged).

- **Users**:  
  This is the user table. It stores whether the user is a consumer or an employee of the bookstore, as well as login information (email and password), and whether the account is active or inactive.  
  - **Id**: Unique identifier for the user.  
  - **User_Type**: The type of user (consumer, employee, or system administrator).  
  - **Name**: The user's name.  
  - **Email**: The user's email.  
  - **Password**: The user's password.  
  - **Enabled**: Indicates whether the user is active or inactive in the system.

- **UsersBooks**:  
  This is the relationship table between book copies and users. It stores whether a user has borrowed a book and also tracks past book returns.  
  - **Id**: Identifier for the relationship.  
  - **User_Id**: Identifier of the user.  
  - **Book**: Identifier for the book copy.  
  - **Pickup_Date**: The date the user borrowed the book.  
  - **Expected_Date**: The date the user should return the book.  
  - **Delivery_Date**: The date the user returned the book.  
  - **Price**: The fine amount paid, if applicable.  
  - **Status**: The status of the book loan (whether it is borrowed, returned, or overdue).

---

## PT-BR

### Como Configurar o Banco de Dados

Você pode configurar o banco de dados de duas maneiras: utilizando o Docker ou configurando seu próprio MySQL ou outra fonte de dados.

1. **Docker**

Para utilizar o Docker, primeiro você deve configurar o arquivo `docker-compose.yml` como desejar ou pode usar as configurações padrão. Após isso, basta executar os comandos no terminal dentro da pasta `db`.

- `docker-compose create`: Este comando vai gerar as configurações do Docker em sua máquina.
- `docker-compose start`: Este comando vai inicializar a imagem do Docker com MySQL em sua máquina.
- `docker-compose exec -p mysql-db mysql -u root -p`: Este comando vai abrir o MySQL e pedir a senha de administrador configurada no arquivo `docker-compose.yml`.

Obs: Se você mudar o nome da imagem, será necessário usar o novo nome no lugar de "mysql-db".

Dentro do MySQL, você precisa usar o banco de dados:
- `USE library;`: Com este comando, você seleciona o banco de dados criado na imagem do Docker.
- `SOURCE script.sql`: Este comando vai gerar as tabelas no banco de dados.

Após seguir todos esses passos, o banco de dados estará configurado.

2. **Outras Fontes**

Você deve configurar o banco de dados com as configurações que serão usadas no backend. Após as configurações estarem corretas, execute o arquivo `script.sql` no novo banco de dados, na database escolhida.

### Estrutura do Banco de Dados

Aqui, explico cada tabela do banco de dados.

- **Genders**:  
  Esta tabela armazena os tipos de gêneros de livros.  
  - **Id**: Identificador único da categoria.  
  - **Name**: Nome da categoria, como por exemplo "drama".  
  - **Description**: Uma breve descrição dos livros do gênero.

- **BooksModels**:  
  Esta tabela armazena os modelos de livros.  
  - **Id**: Identificador único do livro.  
  - **Name**: Nome do livro.  
  - **Release_Date**: Data de lançamento do livro.  
  - **Author**: Autor do livro.  
  - **Language**: Idioma do livro.  
  - **Quantity**: Quantidade de cópias do livro disponíveis na livraria.  
  - **Series**: Saga à qual o livro pertence, caso exista.

- **BooksGenders**:  
  Tabela de relação entre livros e gêneros. Ela armazena os gêneros dos livros.  
  - **Id**: Identificador único da relação.  
  - **Book_Model**: Identificador (id) do livro.  
  - **Gender**: Identificador (id) da categoria do livro.

- **Books**:  
  Esta tabela armazena as cópias dos livros que a livraria possui.  
  - **Id**: Identificador único da cópia do livro.  
  - **Book_Model**: Identificador do modelo do livro.  
  - **Status**: Estado do livro, se está emprestado, disponível ou danificado.

- **Users**:  
  Tabela de usuários. Ela armazena se o usuário é consumidor ou funcionário da livraria, além de informações sobre login e senha (e-mail e senha), e se a conta está ativa ou inativa.  
  - **Id**: Identificador único do usuário.  
  - **User_Type**: Tipo do usuário (consumidor, funcionário ou administrador).  
  - **Name**: Nome do usuário.  
  - **Email**: E-mail do usuário.  
  - **Password**: Senha do usuário.  
  - **Enabled**: Estado do usuário, se está ativo ou inativo no sistema.

- **UsersBooks**:  
  Tabela de relacionamento entre as cópias dos livros e os usuários. Ela armazena os livros que o usuário pegou emprestado e os livros já entregues no passado.  
  - **Id**: Identificador da relação.  
  - **User_Id**: Identificador do usuário.  
  - **Book**: Identificador da cópia do livro.  
  - **Pickup_Date**: Data em que o usuário pegou o livro.  
  - **Expected_Date**: Data limite para a entrega do livro.  
  - **Delivery_Date**: Data de entrega do livro.  
  - **Price**: Multa paga caso haja atraso na devolução.  
  - **Status**: Estado do empréstimo do livro (se está emprestado, entregue ou se a entrega está atrasada).
