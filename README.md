# questioncicas
Questioncicas is a game where you will test your ability to answer questions on the popular culture and knowledge

# Usage

To get a question use: http://localhost:8070/beAsked 
Select betwen one of the posible answers a,b or c and use:
    http://localhost:8070/validate?answer=

... and enjoy Questioncicas!

# Next steps

- Install Go.
- Create a Hello World! project.
- Usage of libraries.
- Creating new packages.
- Creating an HTTP web server.
- Creating a useful application.
 - Add an endpoint to request a question and multiple endpoints to be able to answer them and get a response.
- Basic tests.
- Parameters in HTTP requests.
 - A single endpoint which accepts parameters to define which is the option that was chosen by the user.
 - Instead of saving which question is being answered within the system, add the question number as a parameter.
- Usage of structs.
 - Instead of saving the questions and answers in a map, create a "question" struct (class) which will hold a UUID, the question, it's possible answers, and the correct answer.
 - Convert the question generator into a struct with methods.
 - Convert the question validator into a struct with methods.
 - You might need to create a "setup" function that creates all structs as objects and passes them where needed (dependency injection)
- Usage of interfaces.
 - Create interfaces for the generator and validator and use them when needed instead of using its implementation.
 - In Go there are no "constructors" but you can simply create a Function like `func NewQuestionGenerator` which returns the interface type.
- Autogeneration of mocks.
 - Usage of gomock https://github.com/uber-go/mock
 - Usage of go:generate comments https://github.com/uber-go/mock
- Dockerizing the application.
 - Create a Dockerfile so the app can run within a docker container.
 - Use docker compose to simplify the set up.
 - Usage of docker volumes so the docker container code is instantlly updated when you modify code in your editor.
 - Use a "file watcher" like https://github.com/githubnemo/CompileDaemon so your docker container auto recompiles the application any time you change something in your editor.
- Adding a database.
 - Add a postgres docker container in docker compose.
 - Update the code in your database layer to use GORM so the DB is used.
 - Usage of seeded data with knex.