# School-REST-API
A Go REST API connecting to a PostgreSQL database with CRUD operations. As well as having secured routes with JWT authentications. The use case for this project
is to manage student related data and upload it to a database and returning JSON to the client so it can be displayed for the user.

# Purpose
The main reason for this project was for me to give my self a reason to use the http/net package.
As well other community developed packages dealing with routing and networking.
It also gave me a purpose to make a **Go** based REST API and in all honestly I really enjoyed the development experience.

It was really straightforward dealing with the routing and making the handlers for each routes.
Using [httprouter](https://github.com/julienschmidt/httprouter) for the routing was really enjoyable especially the ease of use for getting the routes parameters and using the values of it, to interact with the handler functions and supply them inside my SQL functions to activate database queries.

The project over is quite simple, it has seven routes each with simple functionalities two of them are protected for they deal with creating, editing and deleting from the database. 
The basic rundown for the routes is that they either pull classroom data (In JSON) from the database or allow creating, editing and deleting from the database.

I also created an account table in the database, where two routes are connected one for creating an admin account with just a email and password.
The password is then encrypted using the bcrypt package before being stored in the database. 
Second route is pulling in the email and hashed password, so that the password can be decrypted to see if the client password is correct when compared to the stored password.
If the user is authenticated they receive a JWT which will then allow them to access the routes dealing with deleting and editing/adding the class table in the database.

Overall this was a simple project but I really did enjoy it and did learn quite a bit especially making tests for SQL based functions, I still need to learn how to make tests for 
the routes though.

# Learned
- Making a REST API in GO.
- Routing and Handlers for the API.
- Securing Routes with JWT.
- Mocking SQL drivers to test using [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock).
- Creating SQL functions to interact with the database when an API is called.
- Making a API with CRUD operations.
- Creating slices of data from the database to convert it to json.
- Practice in creating and using custom data structures.
- Basic usage of Bcrypt package.
- Unpacking payloads from the client and upload it to the database.


# Struggled With
- Testing SQL based functions.
- Testing Routes (Still need to research the subject in depth).
- Making middleware to check JWT.
- Getting the dockerized API working with the dockerized psql db.

# Credits
- Helped me as well teaching me the basic concept of making tests for databases. [article](https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e)
- Sad to say it took me 6 hours to figure out that the query was not matching for the testing was cause of regex. [stack overflow](https://stackoverflow.com/questions/59652031/sqlmock-is-not-matching-query-but-query-is-identical-and-log-output-shows-the-s) and [the tool](https://regex-escape.com/regex-escaper-online.php)

# Issues
- Docker file not working api is refusing connections. But the PostgreSQL container does seem like it is working, just not the go api (need to do some research).
- Route testing has not been done.
