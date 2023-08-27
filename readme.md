#Project by 
@atastemi
@sfaizull

<h1>FORUM</h1>

run project without docker type this on terminal 

``
go run .  
``
  OR  
``
go run -port=<PORT_NUMBER>
``



to run the project build docker
``
docker build -t  my-app .
``           
``
docker run -p 8080:8080 my-app
``