# Codename: true-quasar
> A sci-fi hero-drafting game using Go and Ebiten.

### Setup
Create a MySQL server and save the connection string as `DSN="dsn"` in a file `.env` in the root project directory.

docker run --name my-mysql `
  -e MYSQL_ROOT_PASSWORD=your_password `
  -e MYSQL_DATABASE=your_database `
  -e MYSQL_USER=your_user `
  -e MYSQL_PASSWORD=your_user_password `
  -p 3306:3306 `
  -v new-mysql-data:/var/lib/mysql `
  -d mysql
