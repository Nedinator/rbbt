# 🐸 Rbbt
*A simple URL shortener implemented in Go using Fiber and MongoDB and a frontend with HTML templates.*

## To-do list

- [x] Routes post request at `/new-url`
    - [x] Check if a short URL already exists before creating a new one.
- [x] All setup (for now)
- [x] Actually generate a random but short link to use instead of supplying one
- [x] Create a handler file to manage routes
- [ ] HTML Template for the frontend
    - [x] Home Page
        - [ ] Placeholder for now, gonna add some screenshots and a brief description of the project once TailwindCSS is implemented
    - [x] Stats Page
    - [x] Submit URL Page
    - [x] Search Stuff
        - [x] Needs some serious styling
    - [x] 404 Page
    - [x] General Styling - Chose to go with a minimal design instead and used Water.css.
        - [ ] There's still some formatting I need to fix.
    - [ ] Blog Page
    - [ ] About Page



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
