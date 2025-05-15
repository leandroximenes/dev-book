# Go Getting Started App

This is a simple Go web application using [Gorilla Mux](https://github.com/gorilla/mux), prepared for deployment on [Heroku](https://heroku.com).

## 🚀 Deployment Notes

Before deploying to Heroku, **you must build the project locally**. Heroku expects a compiled binary (e.g. `./main`) to run as defined in the `Procfile`.

### ✅ Steps to Build and Deploy

1. **Build the binary:**

   ```bash
   go build -o main
   ```

2. **Commit the compiled binary:**

   ```bash
   git add main
   git commit -m "chore(module): v1.00.00"
   ```


3. **Push to Heroku:**

   ```bash
   git push heroku master
   ```

   This command will deploy the application to Heroku, using the compiled binary.

