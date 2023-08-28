# Server

This will be `r-ver` server with hosted `plumber` api for scripts and server api for managing plumber

- starting/restarting plumber?
- downloading script results?
- installing R packages
- adding and removing plumber api
- etc

Q: will i need database access here? might be enough just using some local configuration files?

## Server API

| NAME           | DESCRIPTION
|----------------|-----------------------
| add            | Adds new plumber API
| delete         | Removes existing plumber API
| download       | Download result of a script
| package/add    | Install R package?
| package/delete | Removes R package?

## Environment variables

| NAME    | DESCRIPTION
|---------|----------------