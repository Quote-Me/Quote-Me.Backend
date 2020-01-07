# Quote-Me.Backend
A backend for the simplest to use motivator app. Written in golang and gin framework.

# Architecture

# Setup

First seed the data into the database by running:
```bash
mongoimport -h localhost --db quote-me --collection quotes --file data/quotes.json --jsonArray
mongoimport -h localhost --db quote-me --collection photos --file data/photos.json --jsonArray

```

You can configure the --db and --collection parameter. 

After you have done so, create a yaml config file (e.g dev.yaml). The structure should resemble the following:

```yaml
mongo:
  url: mongodb://localhost:27017
  db: quote-me
  quotesCollection: quotes
  photosCollection: photos
server:
  port: :8080
```

After that you can run the server with this command

```bash
go run main.go config.yaml
```

# Credits

Thanks to https://github.com/akhiltak/inspirational-quotes for the incredible amount of data provided.