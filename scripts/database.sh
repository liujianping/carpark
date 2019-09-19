#!/bin/sh

# create database
mysql -uroot -p123456 -e "CREATE DATABASE IF NOT EXISTS db_carpark DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;"

# create tables
for file in $(find $1 -name "*.sql"); do
    mysql -uroot -p123456 < $file
done
