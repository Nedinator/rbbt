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
- [ ] Starting the user auth system
    - [x] Create bcrypt tools
    - [x] Create a user struct
    - [x] Signup/Login Routes
    - [ ] Some form of session management (Jwt or something)
- [ ] There's a lot of nitty gritty stuff I need to do
    - [ ] Homepage formatting
    - [ ] Everything else is anchored to the left still, I want to create a overall layout that'll keep everything the same format in general.
    - [ ] I have enough CSS that I need to move everything to a styles.css.
- [ ] First blog post about my experience with Go and Fiber for this project.
    - [ ] I need to figure out a better way to sort the blog posts. Oldest will be at the top right now.
- [x] Add search to homepage
    - [ ] I need to add a search bar to the homepage to search for links.
    - [ ] Really should just redo the /home... I don't like it.
- [ ] Long-term goals
    - [ ] Adding a user system to maintain your own links and create custom ones
    - [ ] General Security (Limiting requests mainly)
    - [ ] More in-depth statistics
        - [ ] This will be limited due to the fact I want to collect as little data as possible.


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
1. Set method to POST
2. URL - ` http://127.0.0.1:3000/api/new-url`
3. `{"long_url": "https://example.com"}`

**GET**

1. Set method to GET
2. Set URL to `http://127.0.0.1:3000/:id` with ':id' being the shortId of a corresponding document in your MongoDB
