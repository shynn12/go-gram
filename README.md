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
This directory must be started first. This is the heart of the app. The basic logic is happening here. "main" is located in ```cmd```. Then ```go build``` and execute created file.  
❗Important note: The application uses Postgresql and the device must be running a Postgres server.
#### 2. **cmd-gram-blockchain**:
This directory is responsible for storing and writing to the blockchain, it is removed from the main program as it can be used on blockchain nodes to create decentralization.  
To start this up go to ```cmd```. Then ```go build``` and execute created file.
#### 3. **cmd-gram-cli**:
This is the application client. Before using this, you must be sure that all previous prerequisites listed above are met.
### 4. **Configurating**:
To configurate the application you should change data in ```cmd-gram/pkg/client/postgresql/api.toml``` (idk why it`s there, will be changed!).
To configurate the blockchain part you have option to change difficulty of hashing blocks to make POW (Proof of work). It can be changed in ```cmd-gram-blockchain/pkg/blockchain/pow.go``` you need "targetBits" const.
## 💡 Examples:  
For example, let's see how to make a chat with two users:  
Launch the cmd-gram  
![image](https://github.com/user-attachments/assets/aa67521e-7591-4bcd-acca-17c637692301)  
Then launch blockchain  
![image](https://github.com/user-attachments/assets/6f3422d1-77fe-4aa4-b023-2683cc341792)  
Then launch cli with flag where you point the addr of server  





