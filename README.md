# Blockchain Messenger

## Description

This is a messenger in which you don't have to worry about someone saying "there was no such thing", because in addition to saving messages to the Postgresql database, 
it also saves to a small blockchain. Which is known to have the property of immunability, and even its creators will face enormous difficulties if they want to delete or change something.

stack: Go, Postgres, blockchain 
## Features

-**Registration**: creating your own account in the messenger.  
-**Messaging**: you can create chats with other users by knowing their registered e-mail and messaging them.  
-**Blockchain**: All users and chats are stored in the Postgresql database, but messages are also recorded in the blockchain in order to ensure their authenticity.  

## Installation
### 1. **Clone the repository**:
    git clone https://github.com/shynn12/go-gram.git
    cd go-gram
### 2. **Install the required dependencies**:
    go get ./...    
in the each file
### 3. **Starting all up**:
First of all: We have 3 directories in the project. Each of them must be running in order to work correctly.
#### 1. **cmd-gram**:
This directory must be started first. This is the heart of the app. The basic logic is happening here. "main" is located in cmd. Then ```go build``` and execute created file.
#### 2. **cmd-gram-blockchain**:
This directory is responsible for storing and writing to the blockchain, it is removed from the main program as it can be used on blockchain nodes to create decentralization.  
To start this up go to ```cmd```. Then ```go build``` and execute created file.
## **Usage**:

