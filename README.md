# ROOMOOK

ROOMOOK lets you organize your meeting rooms. You can block a chosen meeting room for your meeting or you can just searh for rooms, which are not occupied.

## How to run the system locally
There are diffrent ways to do that. You have to choose what fits best for you.

### 1.Run inside docker containers
Run the container with de docker-compose.yml locatet in the root directory.
```
docker-compose up -d
```

NOTE: the frontend will build in prod mode

### 2. Run Backend/Frontend locally



### Install dependencies
#### Backend
```
go mod tidy
```

more informations here:[ roomook.backend](https://igitlab.jacob.de/ngensheimer/roomook/tree/master/roomook.backend)


#### Frontend
```
npm install 
```

more informtaions here: [roomook.frotend](https://igitlab.jacob.de/ngensheimer/roomook/tree/master/roomook.frontend)

### Start

#### Backend

build it

```
go build
```

start it

```
./roomook.exe
```


#### Frontend

start dev Server

```
npm run serve
```