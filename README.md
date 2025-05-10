# Sky stacks

## Development

### 1. use vscode devcontainer
if you dont use vscode-devcontainer, you should prepare node and go.

Install npm dependences.
```bash
cd client
npm i

cd server 
go install
```

### 2. Run client process

1. Move to client direcotory  
`cd client`  

2. Run dev server  
`npm run dev`  

3. Access to `localhost:5173`

### 3. Run server process

1. Move to server directory  
`cd server`  

2. Run server
`go run .`