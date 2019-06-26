#!/bin/bash

# Example 1 - Basic Usage
clear

echo -n "(1 / 8) - Showing usage with ./db-csv-dump -h"
echo
echo

echo "Running: ./db-csv-dump -h"
echo

./db-csv-dump -h

echo
echo

read -p "Press [Enter] key to continue..."

# Example 2 - Use defaults from config.yml
clear

echo -n "(2 / 8) - Using defaults from config.yml"
echo
echo

echo "Running: ./db-csv-dump"
echo

./db-csv-dump

echo
echo

read -p "Press [Enter] key to continue..."

# Example 3 - override the connection to use
clear

echo -n "(3 / 8) - Override the connection"
echo
echo

echo "Running: ./db-csv-dump -connection kscope"
echo

./db-csv-dump -connection kscope

echo
echo

read -p "Press [Enter] key to continue..."

# Example 4 - override the query
clear

echo -n "(4 / 8) - Override the query using the -e parameter"
echo
echo

echo "Running: ./db-csv-dump -e \"select * from emp\""
echo

./db-csv-dump -connection kscope -e "select * from dept"

echo
echo

read -p "Press [Enter] key to continue..."

# Example 5 - override with a table name
clear

echo -n "(5 / 8) - Use just a table name"
echo
echo

echo "Running: ./db-csv-dump -e dept"
echo

./db-csv-dump -connection kscope -e "dept"

echo
echo

read -p "Press [Enter] key to continue..."

# Example 6 - change the delimiter
clear

echo -n "(6 / 8) - Change the delimiter to a colon"
echo
echo

echo "Running: ./db-csv-dump -connection kscope -e EMP -delimiter \:"
echo

./db-csv-dump -connection kscope -e EMP -delimiter \:

echo
echo

read -p "Press [Enter] key to continue..."

# Example 7 - No header
clear

echo -n "(7 / 8) - Output without a header"
echo
echo

echo "Running: ./db-csv-dump -connection kscope -e EMP -delimiter=\: -noheaders"
echo

./db-csv-dump -connection kscope -e EMP -delimiter=\: -noheaders

echo
echo

read -p "Press [Enter] key to continue..."

# Example 8 - Change the enclosed by
clear

echo -n "(8 / 8) - Change the enclosed by"
echo
echo

echo "Running: ./db-csv-dump -connection kscope -e EMP -delimiter=\: -noheaders -enclosedBy=\^"
echo

./db-csv-dump -connection kscope -e EMP -delimiter=\: -noheaders -enclosedBy=\^

echo
echo

read -p "Press [Enter] key to continue..."


