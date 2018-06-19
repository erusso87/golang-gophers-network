Gophers Network
===============

Gophers network example application made with Golang

#### Endpoints

- [http://localhost:8080](http://localhost:8080) GiQL (GUI)
- [http://localhost:8080/graphql](http://localhost:8080) GiQL (GUI) 

#### TODOs list
- [X] Migrate to GraphQL
- [X] Improve the entities definition
- [ ] Implement more actions
- [ ] Refactor Postgres Docker
- [X] Dev mode (hot reload?)
- [ ] Add nginx as proxy for prod
- [X] Add dependency manager
- [ ] Add unit testing
- [ ] Add e2e testing
- [ ] Extract the GiQL (GUI) to another service 
- [ ] Clean & improve docker-compose config files
- [ ] Solve collisions between Docker images (related with the above one!)
- [ ] Un-replicate config params @ docker-compose.yaml
- [ ] Study how to filter query by attributes (for example by name)
- [ ] Study how to make nested mutations (for example passing image content?)
- [ ] Study how to resolve nested entities (for example the images inside gophers)
- [ ] Do front-end application with VueJS (& Docker, obviously)