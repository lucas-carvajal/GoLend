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
1. Set up main function listening for requests

## Proposed Solution

### Architecture - High Level
Main function accepting all requests and dispatching them to a handler in a Goroutine.
Different service classes take requests and modify the database accordingly.


### Architecture - Classes
`GoBank` with main function accepting request and delegating to handler  
`RequestHandler` processing incoming requests  
tba


### API Endpoints

##### POST
- `/signup` create a new account
- `/login` authenticate with account
- `/claim` file a new claim
- `/claim/<id>/approve` approve a claim
- `/claim/<id>/deny` deny a claim
- `/claim/<id>/settleRequest` ask to settle a claim
- `/claim/<id>/settle` settle a claim
- `/claim/<id>/deadline` set deadline
- `/claim/<id>/interest` set interest

##### GET
- `/dashboard` general account data
- `/claim/<id>` information about a claim


### Database Schema




