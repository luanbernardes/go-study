# go-study
go study udemy (Maximilian)

1) profit_calculator
- First steps with Go simple interactive console
- to run : `cd profit_calculator && go run .`

2) bank
- This is a simple bank application with a few features
- persist data in a txt file
- to run : `cd bank && go run .`

3) final project:
 ![Screenshot from 2024-09-30 14-14-08](https://github.com/user-attachments/assets/851a83dc-95fa-43b6-ae13-4820512f1f1e)

### Anotations
#### Documentation
Testing documentations for APIs:
- with postman: https://documenter.getpostman.com/view/1369025/2sAXxMhZyi
  - Simple to export and share for every one, every engineer can see and use too 
  - It is public
  - Don't need hosting
- with swagger: WIP

### APIs that will have:
- get /events: get all events
- get /events/<id>: get event

- post /events: create event - auth required
- put /events/<id>: update event - auth required - only owner can update
- delete /events/<id>: delete event - auth required - only owner can delete

- post /signup: create user

- post /login: login user - JWT token
- post /events/<id>/register: register user to event - auth required
- delete /events/<id>/register: unregister user from event - auth required
