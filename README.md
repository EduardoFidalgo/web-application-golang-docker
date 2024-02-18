Execute this commands to start the project

Building the project with all dependencies starting application + mysql database:
- sudo docker-compose up -- build

Migrate default tables to your mysql (change the parameters as you need):
- sudo migrate -path migrations -database mysql://root:secret@tcp(localhost:3306)/webservice-database up

Enjoy! =)