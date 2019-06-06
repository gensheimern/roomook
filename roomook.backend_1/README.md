# RaumbuchungsApp

## Getting started

Install dependencies

``` 
go mod tidy
```

If you dont have GO111MODULE=on set in your env you get an error message.
In order to fix this:

# Windows:

temporarly setting this option in CMD
```
export GO111MODULE=on
```
for permanent settings, set it to your enviroment variables in your system properties

# Mac & Linux

not tested because i´m using a windows machine

## Tests

Tests are not fully implemented yet. To run tests run:

```
go test
```

## Build

To build you have to install the dependencies first. Then run:

```
go build
```

## API Routes

### 1. Room

Fields:
*  "RoomID" is a unique String which identifies the room
*  "Name" is a optional field to set fictive names to a room like Aquarium

#### 1.1 Get Room information

##### 1.1.1 Get all rooms

``` 
GET /room
URL Params:
- none

Request body:
- none

Response body:

{
    "roomID": "33B",
    "name": "Aquarium"
},
{
    "roomID": "11",
    "name" : "11"
}

```
##### 1.1.2 Get a single room
```
GET /room/:id
URL Params:
- RoomID

Request body:
- none

Response body:
{
    "roomID": "33B",
    "name": "Aquarium"
}
```

##### 1.1.3 Get all available rooms in between 2 date times

```
POST /r/d/
URL Params:
- none

Request body:
{
    "begin" : "2019-05-27 16:05:00",
    "end" : "2019-05-27 16:10:00"
}

Response body:
{
    "roomID": "33B",
    "name": "Aquarium"
}
```

#### 1.2 Create a room

```
POST /room/
URL Params:
- none

Request body:
{
 "RoomID" : "33B",
 "Name" : "Aquarium"
}

Response body:
{
    "roomID": "33B",
    "name": "Aquarium"
}

```
#### 1.2 Update a room
```
PUT /room/:id
URL Params:
- RoomID

Request body:
{
 "roomID" : "33B",
 "name" : "Kitchen"
}

Response body:
{
    "roomID": "33B",
    "name": "Kitchen"
}

```
#### 1.2 Delete a room

```
DELETE /room/:id
URL Params:
- RoomID

Request body:
- none

Response body:
{
    "romID": "33B",
    "name": "Kitchen"
}
```

### 2. Booking

Fields:
* BookingID is a unique integer Number which identifies the booking
* ParentID, each booking has a parent booking, where some Information are set like: 
    
* Title
* UserID, this will e automaticly set to the user logged in
* RoomID
* Begin
* End
* RecurringType
    * Recurring Type can be:
        * daylie
        * 2-days
        * weekly
        * month
    

#### 2.1 Get Booking information

##### 2.1.1 Get all bookings



``` 
GET /booking
URL Params:
- none

Request body:
- none

Response body:

 {
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "mragusa",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "weekly"
},
{
    "bookingID": "5",
    "parentID": "6",
    "title": "Test",
    "userID": "gzicca",
    "roomID": "1",
    "begin": [
        "2019-05-27 16:05:00"
    ],
    "end": [
        "2019-05-27 16:10:00"
    ],
    "recurringType": "single"
},
```
##### 2.1.2 Get a single booking
``` 
GET /booking/:id
URL Params:
- bookingID

Request body:
- none

Response body:

 {
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "mragusa",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "weekly"
}
```
##### 2.1.3 Get all bookings from a specific room
``` 
GET /b/r/:id
URL Params:
- roomID

Request body:
- none

Response body:

 {
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "mragusa",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "weekly"
},
{
    "bookingID": "5",
    "parentID": "6",
    "title": "Test",
    "userID": "gzicca",
    "roomID": "1",
    "begin": [
        "2019-05-27 16:05:00"
    ],
    "end": [
        "2019-05-27 16:10:00"
    ],
    "recurringType": "single"
},
```


#### 2.2 Create a booking

You can pass multipe date-time strings inside an array to create multiple bookings

```
POST /booking
URL Params:
- none

Request body:
{
    "title": "Items",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "single"
}

Response body:
{
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "mragusa",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "single"
}
```

#### 2.3 Update a booking

```
PUT /booking/:id
URL Params:
- bookingID

Request body:
{
    "title": "Items",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ]
}

Response body: 
{  
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "ngensheimer",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "single"
}
```

#### 2.4 Delete a booking

##### 2.4.1 Delete a single booking

```
DELETE /booking/:id
URL Params:
- bookingID

Request body:
- none

Response body: 
{  
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "ngensheimer",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "single"
}
```

##### 2.4.2 Delete recurring bookings

```
DELETE /booking/:id/recurrent
URL Params:
- bookingID

Request body:
- none

Response body: 
{  
    "bookingID": "4",
    "parentID": "5",
    "title": "Items",
    "userID": "ngensheimer",
    "roomID": "1",
    "begin": [
        "2019-05-27 14:18:00"
    ],
    "end": [
        "2019-05-27 15:48:00"
    ],
    "recurringType": "single"
}
```

### 3. Websockets

In this project we are using Websockets to have a "live-synchronization".


```
GET /ws
URL Params:
- none

Request body:
- none

Response body: 
- none
```