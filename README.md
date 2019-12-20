# Quote-Me.Backend
A backend for the simplest to use motivator app. Written in golang and serverless framework.

# Architecture

# Setup

First seed the data into the database by running:
```bash
mongoimport -h localhost --db quote-me --collection quotes --file data/quotes.json --jsonArray
```

You can configure the --db and --collection parameter. 

After you have done so, create a yaml config file (e.g dev.yaml). The structure should resemble the following:

```yaml
mongo:
  url: mongodb://localhost:27017
  db: quote-me
  quotesCollection: quotes
server:
  port: :8080
unsplash:
  clientId: <your app_key from unsplash>
  width: 1500
  height: 700
  keywords: 
    - mountain
    - lake
    - friends
    - christmas
    - beach
    - sea
    - landscape
    - jungle
    - waterfall
    - night-sky
    - space
    - nature
    - architecture
    - travel
    - arts-culture
```

After that you can run the server with this command

```bash
go run main.go dev
```

# Credits

Thanks to https://github.com/akhiltak/inspirational-quotes for the incredible amount of data provided.