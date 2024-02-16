# 🐸 Rbbt
*A simple URL shortener implemented in Go using Fiber and MongoDB and a frontend with GO (using Fiber) HTML templates.*

## To-do list

- [x] Routes post request at `/new-url`
    - [x] Check if a short URL already exists before creating a new one.
- [x] All setup (for now)
- [x] Actually generate a random but short link to use instead of supplying one
- [x] Create a handler file to manage routes
- [x] HTML Template for the frontend
    - [x] Home Page
        - [x] Placeholder for now, gonna add some screenshots and a brief description of the project once TailwindCSS is implemented
    - [x] Stats Page
    - [x] Submit URL Page
    - [x] Search Stuff
        - [x] Needs some serious styling
    - [x] 404 Page
    - [x] General Styling - Chose to go with a minimal design instead and used Water.css.
        - [x] There's still some formatting I need to fix.
    - [x] Blog Page
        - [x] Placeholder for now, going to add some blog posts about the project and other stuff once it's finished.
    - [x] About Page
        - [x] Some general styling will fix this page up.
- [x] Create routing package
- [ ] V3 rewrite lol
    - [ ] Switching to Auth0.
        - [ ] Actually not till after I figure out my own setup :)
        - [ ] Auth handling is actual finished now, but needs adjusted now that it's not just an API.
    - [x] Switching to Tailwindcss (i'll use the cdn for now, but create my own imports laters)
- [x] Removed blogs (i really don't know why I added it)
- [ ] Custom toasts for errors and success messages


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
> This is currently not working, I plan on adding a api module to the project to handle this. Originally, that's all this was going to be.

**POST**
- Set method to POST
- URL - ` http://127.0.0.1:3000/api/new-url`
- `{"long_url": "https://example.com"}`

**GET**

- Set method to GET
- Set URL to `http://127.0.0.1:3000/:id` with ':id' being the shortId of a corresponding document in your MongoDB

6. **Setup ENV variables**

On MacOS or Linux, you can set the environment variable like this:
```bash
export JWT_SECRET="your_jwt_secret"
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
