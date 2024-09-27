![image](https://github.com/user-attachments/assets/e4b83f9b-300e-4233-b2f3-7fa31fd5af44)# Blockchain Messenger

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
## Usage:  
### 1. **cmd-gram**:
This directory must be started first. This is the heart of the app. The basic logic is happening here. "main" is located in ```cmd```. Then ```go build``` and execute created file.  
⚠️If you have troubles with reading ```.toml``` file you can use the ```-config-path``` command to set the configuration path manually.
❗Important note: The application uses Postgresql and the device must be running a Postgres server.
### 2. **cmd-gram-blockchain**:
This directory is responsible for storing and writing to the blockchain, it is removed from the main program as it can be used on blockchain nodes to create decentralization.  
To start this up go to ```cmd```. Then ```go build``` and execute created file.
### 3. **cmd-gram-cli**:
This is the application client. Before using this, you must be sure that all previous prerequisites listed above are met. Use the ```-ip``` flag to select the server address to connect to.
### 4. **Configurating**:
To configurate the application you should change data in ```cmd-gram/pkg/client/postgresql/api.toml``` (idk why it`s there, will be changed!).
To configurate the blockchain part you have option to change difficulty of hashing blocks to make POW (Proof of work). It can be changed in ```cmd-gram-blockchain/pkg/blockchain/pow.go``` you need "targetBits" const.
## 💡 Examples:  
For example, let's see how to make a chat with two users:  
Launch the cmd-gram  
![image](https://github.com/user-attachments/assets/aa67521e-7591-4bcd-acca-17c637692301)  
Then launch blockchain  
![image](https://github.com/user-attachments/assets/6f3422d1-77fe-4aa4-b023-2683cc341792)  
(When the blockchain was launched, the first block was created (genesis block))
Then run the cli with the flag where you specify the server address with the -ip.  
![image](https://github.com/user-attachments/assets/e08afe64-e511-4b88-965a-3617558647c2)  
Congratulations! Now we are ready to start.
Register your account  
![image](https://github.com/user-attachments/assets/be7846ff-369a-4b3f-8f66-5706f99c72d3)  
Create a chat with a user whose email address you know
![image](https://github.com/user-attachments/assets/dc6a4ea4-804e-4a33-ad32-374539aae07f)  
and then you can see all your chats with its id if type ```/all-chats```:  
![image](https://github.com/user-attachments/assets/3cc8b99f-9d68-4704-a14c-419681057123)
Select a chat by typing ```/open-chat``` and entering its ID. For example ```/open-chat 43```. If you have messages they will be shown.  
![image](https://github.com/user-attachments/assets/57f2e326-fb45-4c66-a2da-35f331809dfb)  
Since you are in a chat, all messages except the ```/exit``` command will be shown to all chat participants).








