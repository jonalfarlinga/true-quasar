# Codename: true-quasar
> A sci-fi hero-drafting game using Go and Ebiten.

### Setup
Create a MySQL server and save the connection string as `DSN="dsn"` in a file `.env` in the root project directory.

docker volume create true-quasar-data

docker run --name quasar-db `
  -e MYSQL_ROOT_PASSWORD=admin `
  -e MYSQL_DATABASE=entities `
  -e MYSQL_USER=consilium `
  -e MYSQL_PASSWORD=true-quasar `
  -p 3306:3306 `
  -v true-quasar-data:/var/lib/mysql `
  -d mysql

docker exec -it quasar-db mysql -u root -p
