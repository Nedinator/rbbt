# 🐸 Rbbt
*A simple URL shortener implemented in Go using Fiber and MongoDB and a frontend with HTML templates.*

## To-do list

- [x] Routes post request at `/new-url`
    - [x] Check if a short URL already exists before creating a new one.
- [x] Routes get request at `/:url`
- [x] Actually generate a random but short link to use instead of supplying one
- [x] Create a handler file to manage routes
- [ ] HTML Template for the frontend
    - [ ] Home Page
    - [x] Stats Page
    - [x] Submit URL Page
    - [x] Search Stuff
        - [ ] Needs some serious styling
    - [ ] 404 Page
    - [ ] General Styling

## Setup

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Nedinator/rbbt.git
   cd rbbt
   ```

2. **Install Dependencies**
    Make sure you have Go installed. Then, run:
    ```bash
    go mod tidy
    ```

3. **Configure MongoDB:**
    Update the mongoURI variable in mongo.go with your MongoDB connection string.


4. **Run the Application:**

    ```bash
    go run .
    ```
5. **Sending test GET/POST requests**
*I recommend the POSTMAN extension on VSCode to do this*

**POST**
1. Set method to POST
2. URL - ` http://127.0.0.1:3000/api/v1/new-url`
3. `{"long_url": "https://example.com"}`

**GET**

1. Set method to GET
2. Set URL to `http://127.0.0.1:3000/api/v1/:id` with ':id' being the short_url of a corresponding document in your MongoDB
