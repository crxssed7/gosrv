# gosrv

A simple web server written in Go with database capabilities and resourceful routing.

On first boot, the project creates an SQLite database called people.db in the current working directory. You can then navigate to http://127.0.0.1:1337 and create, read, update, delete people (CRUD).

TODO:
- API layer
- Move everything that could be moved into "global" helper methods
- Add a service directory for business logic (interacting with third party apis, helper methods etc)
- Testing framework
- Add cli tools to generate things such as models, resourceful routes, templates etc
- Add some frontend frameworks ?????
- CLI tool to generate a new project exactly like this but with a custom name and optionals such as a frontend framework, custom testing framework etc
