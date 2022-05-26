# Sparrow Blockchain

## Requirements
* GoLang
## To Run
First we need to ensure the wallet server is working properly. In a terminal, navigate into `wallet_server` and run the following command:
```
go run main.go wallet_server.go
```

After this command is run, open a new terminal and navigate to `blockchain_server`. From here run the following command:
```
go run main.go blockchain_server.go -port [PORT]
```

Each time you run this command you'll need a new terminal, and this will act as a Node on the blockchain. Simply replace `[PORT]` with a port from 5001 to 5005 (each instance must have a unique port).

## To Access

With the above services running, open a website to `localhost:8080` to view a wallet. Whenever you reload the page, or open a new tab to the page, you'll be given a new wallet ID.

To check the current status of a node, simply copy the output from when you ran the blockchain_server command to the wallet details on the page. After a few seconds you should see the current count update to the new current value of that node.