# GoBank - DesignDoc

## Overview
GoBank provides a back-end API for an online banking service.  
Specifically it provides users the ability to manage how much money they owe to other people and how much money is owed to them.
Users can file claims, accept and deny claims, manage their claims and their overall balance.

## Goals
Provide an easy REST API for frontend services.
Users can:  
- Create an account
- View their claims and the claims to them
- Make a claim
- Ask to settle a claim
- Approve a settlement request
- Set an interest percentage for a claim
- Set a deadline for a claim

## Milestones
1. ~~Set up database with schema and connect to it via code in Go~~
2. Set up GoBank file and GoBankServer file with main function listening for requests
3. Set up RequestsHandler file
4. Set up UserManagementService & ClaimManagementService files
5. Implement all functionality to satisfy goals

## Proposed Solution

### Architecture - High Level
Main function accepting all requests and dispatching them to a handler in a Goroutine.
Different service classes take requests through channels from the handlers and modify the database accordingly.


### Architecture - Classes
`GoBank` calling GoBankServer's main function 
`GoBankServer` with main function accepting request and delegating to handler, initializes service classes
`RequestHandler` processing incoming requests, started as goroutine  

`UserManagementService` for managing all the users, channel passed to each handler goroutine to send requests  
`ClaimManagementService` for managing all the claims, channel passed to each handler goroutine to send requests


### API Endpoints

##### POST
- `/signup` create a new account
- `/login` authenticate with account -> get token to authenticate
- `/claim` file a new claim
- `/claim/<id>/approve` approve a claim
- `/claim/<id>/deny` deny a claim
- `/claim/<id>/settleRequest` ask to settle a claim
- `/claim/<id>/settle` settle a claim
- `/claim/<id>/deadline` set deadline
- `/claim/<id>/interest` set interest

##### GET
- `/` general account data
- `/claim/<id>` information about a claim


### Database Schema

---

##### Users
username: String - **PK**  
password: String  
email: String  

---

##### Loans  
id: Int - **PK**  
fromUser: username from UserTable - **FK**    
toUser: username from UserTable - **FK**    
date: Int    
amount: Int  
interest: Double  
status: String  

---

##### OpenBoard
openID: Int  
fromUser:  username from UserTable - **FK**  
toUser:  username from UserTable - **FK**  
loan:  id from LoanTable - **FK**  

---

##### CloseBoard
closeID: Int  
fromUser:  username from UserTable - **FK**  
toUser:  username from UserTable - **FK**  
loan:  id from LoanTable - **FK**

---


