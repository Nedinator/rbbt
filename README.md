# 🐸 Rbbt

_A simple URL shortener implemented in Go using Fiber and MongoDB and a frontend with GO (using Fiber) HTML templates._

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
- [x] Auth handling is actual finished now, but needs adjusted now that it's not just an API.
- [x] Save referer data for a little bit more insight into clicks
- [x] Switching to Tailwindcss (i'll use the cdn for now, but create my own imports laters)
- [x] Removed blogs (i really don't know why I added it)
- [x] Currently /stats/:id loads referers whether you're logged in on the owner account or not.
- [x] Customize toasts for errors and success messages
- [x] Setup middleware for dashboard and retrieve URLs associated with user
- [x] Found a bug with the stats.html page. When you search for the link it works fine but going from dashboard uses the wrong url for the link to the short url. I think I may just change this to a button that copies the link to the clipboard.
- [x] Decided on server side for graphs.
  - [x] I have to decide whether static and server-side with go or dynamic and client-side with JS
- [ ] Signed up user upgrades
  - [x] Custom short URLs
    - [ ] Need to check this. Right now it's just whatever they type which could cause duplicates.
  - [ ] Custom referer data
    - [ ] Custom tags to categorize for analytics
  - [ ] Delete URLs
- [ ] Redo the home page. Finally add screenshots and more details.
- [x] Switch to GORM with with postgres

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

3. **Configure app:**
   There's some environment variables you'll need. Maybe someone can write a package to handle this better who knows.

   ```bash
   export DOMAIN=127.0.0.1:3000
   export MONGO_URI=mongodb://localhost:27017
   export JWT_SECRET=this_can_be_anything_just_keep_it_the_same_or_users_cant_login
   ```

4. **Run the Application:**

   ```bash
   go run main.go
   ```

5. **Or Build the Application**

   ```bash
   go build -o ribbit .
   ./ribbit
   ```

## License

[MIT](https://choosealicense.com/licenses/mit/)

```

```
