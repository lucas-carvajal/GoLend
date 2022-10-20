## GoBank API

### What is it and what does it do?
GoBank provides a back-end API for an online banking service.  
Specifically it provides users the ability to manage how much money they owe 
to other people and how much money is owed to them. 
Users can file claims, accept and deny claims, manage their claims 
and their overall balance.

### How to use it?
1. Create a MySQL database named golend
2. Apply the `create-tables.sql` file to it e.g.  
    `source /path/to/file/on/my/computer/create-tables.sql`
3. Have {user: 'root'} with {password: ''} in MySQL, otherwise change the configuration in the
   `setUpAndTestDBConnection()` function in the `main.go` file

### What did I learn?
* Go
* MySQL
* Go `database/sql` library

### Disclaimers
Security was not a focus in this project as the main purpose
was to build an API in Go and have some fun along the way.
As such, for example arguments for DB queries are not 
sanitized (and thus subsceptible to SQL injections),
so it is not advised to use this in any kind of serious application.
