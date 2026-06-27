
# Role based Authentication in Golang using JWT Authentication

A simple role based Authentication system in Golang using JWT Authentication. DB used is Mongodb.

There are 3 levels of users. Unregistered, registered and Admin.

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`MONGO_URI`=<URL>

`MONGO_DB_NAME`=<NAME>

`JWT_SECRET`=<Secret Text>

`DB collection name` = "users"
