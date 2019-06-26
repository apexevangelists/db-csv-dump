# db-csv-dump

Command to connect to an Oracle Database and quickly export data from a Table, View or SQL Query in CSV format.

## Prerequisites

Oracle Instant Client must already be installed

[Oracle Instant Client](https://www.oracle.com/database/technologies/instant-client.html)

Note - Oracle Instant Client must be configured per your environment (please follow the instructions provided by Oracle).

## Table of Contents

- [db-csv-dump](#db-csv-dump)
  - [Prerequisites](#Prerequisites)
  - [Table of Contents](#Table-of-Contents)
  - [Installation](#Installation)
  - [Building](#Building)
  - [Usage](#Usage)
  - [Support](#Support)
  - [Contributing](#Contributing)

## Installation

1) Clone this repository into a local directory, copy the dbdump executable into your $PATH

```bash
$ git clone https://github.com/apexevangelists/db-csv-dump
```

## Building

Pre-requisite - install Go

Compile the program -

```bash
$ go build
```

## Usage

```bash-3.2$ ./db-csv-dump -h
Usage of ./db-csv-dump:
  -configFile string
        Configuration file for general parameters (default "config")
  -connection string
        Confguration file for connection
  -db string
        Database Connection, e.g. user/password@host:port/sid
  -debug
        Debug mode (default=false)
  -delimiter string
        Delimiter between fields (default ",")
  -e string
        Table, View or SQL Query to export
  -enclosedBy string
        Fields enclosed by (default "\"")
  -export string
        Table, View or SQL Query to export
  -noheaders
        Omit Headers
bash-3.2$
```

## Support

Please [open an issue](https://github.com/apexevangelists/db-csv-dump/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/apexevangelists/db-csv-dump/compare).