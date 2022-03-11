# School-REST-API
A Go REST API connecting to a PostgreSQL database with CRUD operations. As well as having secured routes with JWT authentications. The use case for this project
is to manage student related data and upload it to a database and returning JSON to the client so it can be displayed for the user.

# Purpose
The main reason for this project was for me to give my self a reason to use the http/net package.
As well other community developed packages dealing with routing and networking.
It also gave me a purpose to make a **GO** based REST API and in all honestly I really enjoyed the development experience.
It was really straightforward dealing with the routing and making the handlers for each routes.
Using https://github.com/julienschmidt/httprouter for the routing was really enjoyable especially the ease of use for getting the routes parameters and using the values of it, to interact with the handler functions and supply them inside my SQL functions to activate database queries.

The project over is quite simple, it has seven routes each with simple functionalities two of them are protected for they deal with creating, editing and deleting from the database. 
The basic rundown for the routes is that they either pull classroom data from the database or allow creating, editing and deleting the from database.
I also created an account table in the database, where two routes are connected one for creating an admin account with just a email and password.
The password is then encrypted use bcrypt package before being stored in the database. 
Second route is pulling in the email and hashed password so the password can be decrypted to see if the client password is correct when compared to the stored password.
If the user is authenticated they receive a JWT which will then allow them to access the routes dealing with deleting and editing/adding the class table in the database.
Overall this was a simple project but I really did enjoy it and did learn quite a bit especially making tests for SQL based functions, I still need to learn how to make tests for 
the routes though.
