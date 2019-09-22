carpark
===

simple carpark service for Singapo.

## 1. Installation

````bash

$: git clone github.com/liujianping/carpark
$: cd carpark
$: make image
# make sure 127.0.0.1:8080 is available
$: make up 
# wait for a short time when mysql is already
$: make initdb

# check service
$: curl -X GET http://localhost:8080/carparks/nearest?latitude=1.37326&longitude=103.897&page=0&per_page=3

# stop service
$: make down
````

## 2. Source from scratch

````bash
$: git clone github.com/liujianping/carpark
$: cd carpark 
$: tree -L 1 .
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── dataset
├── docker-compose.yml
├── etc
├── go.mod
├── go.sum
├── http
├── job
├── main.go
├── orm
├── prepare
├── scripts
└── svy21
````

### 2.1 dependence

please install dependence tools firstly,

- [github.com/auto-program/db-orm](https://github.com/auto-program/db-orm)


### 2.2 data schema design

use `yaml` to define the database schema, you can find the only two database model in  folder `orm/yaml`.

use `db-orm` to generate the sql scripts & model code.

````bash
$: make clean
$: make gen
````

then you can create the database `db_carpark` in your dev mysql server(MySQL 5.7). please execute the `orm/sql` in database `db_carpark`.

### 2.3 preprocess carpark dataset from csv file

In `csv` file, some fields need to be preprocessed. I do the following fields:

- "x_coord","y_coord"
  
  convert this two field in `svy21` format to `latitude`, `longitude` format.

- "short_term_parking"

  anaylze this field to <from, to> pair format

- string type `YES/NO` to boolean type

then save the csv items to the database table.


````bash

$: go run main.go prepare -f dataset/hdb-carpark-information.csv 
````

before running, make sure the `etc/config.toml` is configured properly. 

### 2.4 schedule the job to updating the carpark info regularly.

````bash
$: go run main.go job
````

before running, make sure the `etc/config.toml` is configured properly. 

And you can configure the job schedule by flag `--crontab`, it's default setting is `* * * * *`, means every minute.


### 2.5 starting the http service

````bash
$: go run main.go http
````

before running, make sure the `etc/config.toml` is configured properly and `127.0.0.1:8080` not be occupied by other programs. 

then you can try the service by the command:

````
$: curl -X GET \
   http://localhost:8080/carparks/nearest?latitude=1.37326&longitude=103.897&page=3&per_page=3
````

