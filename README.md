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
APIs:
- get /events: get all events
- get /events/<id>: get event
- post /events: create event - auth required
- put /events/<id>: update event - auth required - only owner can update
- delete /events/<id>: delete event - auth required - only owner can delete

- post /signup: create user

- post /login: login user - JWT token
- post /events/<id>/register: register user to event - auth required
- delete /events/<id>/register: unregister user from event - auth required