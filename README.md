# atomic

## Installation

Clone this repository ``git clone https://github.com/ehardi19/atomic.git``

or

``go get github.com/ehardi19/atomic``

## Running

``go run main.go``


## Usage

1. Class

1.1. Creating New Class

POST localhost:8080/class

1.2. Getting All Classes

GET localhost:8080/class/

1.3. Getting Class by id

GET localhost:8080/class/:id

1.4. Getting All Classes by topic

GET localhost:8080/topic/
1.5. Updating a Class by id

PUT localhost:8080/class/:id

1.6. Deleting a Class by id

DELETE localhost:8080/class/:id

2. Registrant

1.1. Creating New Registrant

POST localhost:8080/registrant

1.2. Getting All Registrants

GET localhost:8080/registrant/

1.3. Getting Registrant by id

GET localhost:8080/registrant/:id

1.4. Updating a Registrant by id

UPDATE localhost:8080/registrant/:id

1.5. Deleting a Registrant by id

DELETE localhost:8080/registrant/:id
