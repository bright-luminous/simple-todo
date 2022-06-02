# Simple Todo

Your todo app will never be this simple. Server is listening on port 3000.

## Quickstart

Download all dependencies

```sh
go mod download
```

To start the server, simply run

```sh
make
```

## Available APIs

1. Echo request

   ANY_METHOD `/`

   This route accepts any request and it back to the client.

   Example result

   ```json
   {
      "data": {
         "method": "POST",
         "header": {
               "Accept": [
                  "*/*"
               ],
               "Accept-Encoding": [
                  "gzip, deflate, br"
               ],
               "Connection": [
                  "keep-alive"
               ],
               "Content-Length": [
                  "0"
               ],
               "Postman-Token": [
                  "22527ab1-601d-4837-8611-497cd566e59a"
               ],
               "User-Agent": [
                  "PostmanRuntime/7.29.0"
               ]
         },
         "body": {},
         "query": {}
      }
   }
   ```

1. Get all todos

   GET `/todos`

   This route shows all todos in the database.

   Example result

   ```json
   {
      "data": [
         {
               "id": "05360b86-5911-4935-bc50-324697632c62",
               "task": "Learn how to code Go",
               "description": "Go is easy!",
               "createdAt": "2022-06-02T10:03:58.581539847+07:00",
               "done": false
         },
         {
               "id": "b2d598b1-4eee-4701-8a36-65315dee1722",
               "task": "Buy coffee",
               "description": "I'm so sleepy!",
               "createdAt": "2022-06-02T10:03:58.581541651+07:00",
               "done": false
         },
         {
               "id": "2d3c9e68-0d18-42db-9e86-13ae4cf70e55",
               "task": "Clean the room",
               "description": "Is this really a human room?",
               "createdAt": "2022-06-02T10:03:58.581543334+07:00",
               "done": false
         }
      ]
   }
   ```

1. Get a todo by ID

   GET `/todos/{id}`

   This route returns a todo by the given ID.

1. Create a new todo

   POST `/todos`

   This route creates a new todo in the database.

   Request body

   ```json
   {
      "task": "string",
      "description": "string"
   }
   ```

1. Update an existing todo

   PUT `/todos/{id}`

   This route updates an existing todo in the database.

   Request body

   ```json
   {
      "task": "string",
      "description": "string",
      "done": false
   }
   ```
