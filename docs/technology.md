# Technology Selection Rationale 

`chi`: lightweight and extensible HTTP router in Golang
 - `chi` provides a robust middleware package and is less opinionated and lighter than alternatives like `Gin`

`postgres`: provides strict atomic transaction guarantees required by the sheer scale of drop culture load 

`redis`: tentative solution for routing around database row contention. 
