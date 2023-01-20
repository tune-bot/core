# database
go module for interacting with the tune-bot database

TODO: document how to use the module

When connecting to the database: you need to pass in a string like this:
> user:password@tcp(host)/database
For example:
`err := database.Connect("myUserName:MySecretPassword@tcp(my-domain.com)/myDatabase")`
