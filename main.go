package main

/* db-csv-dump */

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	_ "gopkg.in/goracle.v2"
)

// TConfig - parameters in config file
type TConfig struct {
	configFile       string
	debugMode        bool
	delimiter        string
	enclosedBy       string
	noheaders        bool
	export           string
	connectionConfig string
	connectionsDir   string
}

// TConnection - parameters passed by the user
type TConnection struct {
	dbConnectionString string
	username           string
	password           string
	hostname           string
	port               int
	service            string
}

var config = new(TConfig)
var connection TConnection

var logger = loggo.GetLogger("dbcsvdump")

/********************************************************************************/
func setDebug(debugMode bool) {
	if debugMode == true {
		loggo.ConfigureLoggers("dbcsvdump=DEBUG")
		logger.Debugf("Debug log enabled")
	}
}

/********************************************************************************/
func parseFlags() {

	flag.StringVar(&config.configFile, "configFile", "config", "Configuration file for general parameters")
	flag.StringVar(&config.delimiter, "delimiter", ",", "Delimiter between fields")
	flag.StringVar(&config.enclosedBy, "enclosedBy", `"`, "Fields enclosed by")
	flag.BoolVar(&config.noheaders, "noheaders", false, "Omit Headers")
	flag.StringVar(&config.export, "export", "", "Table, View or SQL Query to export")
	flag.StringVar(&config.export, "e", "", "Table, View or SQL Query to export")

	flag.BoolVar(&config.debugMode, "debug", false, "Debug mode (default=false)")
	flag.StringVar(&config.connectionConfig, "connection", "", "Confguration file for connection")

	flag.StringVar(&connection.dbConnectionString, "db", "", "Database Connection, e.g. user/password@host:port/sid")

	flag.Parse()

	// At a minimum we either need a dbConnection or a configFile
	if (config.configFile == "") && (connection.dbConnectionString == "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

}

/********************************************************************************/
func getPassword() []byte {
	fmt.Printf("Password: ")
	pass, err := gopass.GetPasswd()
	if err != nil {
		// Handle gopass.ErrInterrupted or getch() read error
	}

	return pass
}

/********************************************************************************/
func getConnectionString(connection TConnection) string {

	if connection.dbConnectionString != "" {
		return connection.dbConnectionString
	}

	var str = fmt.Sprintf("%s/%s@%s:%d/%s", connection.username,
		connection.password,
		connection.hostname,
		connection.port,
		connection.service)

	return str
}

/********************************************************************************/
// To execute, at a minimum we need (connection && (object || sql))
func checkMinFlags() {
	// connection is required
	bHaveConnection := (getConnectionString(connection) != "")

	// check if we have either an object to export or a SQL statement
	bHaveObject := (config.export != "")

	if !bHaveConnection || !bHaveObject {
		fmt.Printf("%s:\n", os.Args[0])
	}

	if !bHaveConnection {
		fmt.Printf("  requires a DB connection to be specified\n")
	}

	if !bHaveObject {
		fmt.Printf("  requires either an Object (Table or View) or SQL to export\n")
	}

	if !bHaveConnection || !bHaveObject {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

/********************************************************************************/
func loadConfig(configFile string) {
	if config.configFile == "" {
		return
	}

	logger.Debugf("reading configFile: %s", configFile)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// need to set debug mode if it's not already set
	setDebug(viper.GetBool("debugMode"))

	config.connectionsDir = viper.GetString("connectionsDir")
	config.connectionConfig = viper.GetString("connectionConfig")

	config.debugMode = viper.GetBool("debugMode")

	if (viper.GetString("export") != "") && (config.export == "") {
		logger.Debugf("loadConfig: export loaded: %s\n", viper.GetString("export"))
		config.export = viper.GetString("export")
	}
	config.configFile = configFile
}

/********************************************************************************/
func loadConnection(connectionFile string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(config.connectionConfig)
	v.AddConfigPath(config.connectionsDir)

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if v.GetString("dbConnectionString") != "" {
		connection.dbConnectionString = v.GetString("dbConnectionString")
	}
	connection.username = v.GetString("username")
	connection.password = v.GetString("password")
	connection.hostname = v.GetString("hostname")
	connection.port = v.GetInt("port")
	connection.service = v.GetString("service")

	if (viper.GetString("export") != "") && (config.export == "") {
		logger.Debugf("loadConnection: export loaded: %s\n", v.GetString("export"))
		config.export = v.GetString("export")
	}

}

/********************************************************************************/
func debugConfig() {
	logger.Debugf("config.configFile: %s\n", config.configFile)
	logger.Debugf("config.debugMode: %s\n", strconv.FormatBool(config.debugMode))
	logger.Debugf("config.delimiter: %s\n", config.delimiter)
	logger.Debugf("config.enclosedBy: %s\n", config.enclosedBy)
	logger.Debugf("config.noheaders: %s\n", strconv.FormatBool(config.noheaders))
	logger.Debugf("config.export: %s\n", config.export)
	logger.Debugf("config.connectionConfig: %s\n", config.connectionConfig)
	logger.Debugf("connection.dbConnectionString: %s\n", connection.dbConnectionString)
}

/********************************************************************************/
func outputHeaders(cols []string) {
	// Output the headers
	for i, colName := range cols {
		fmt.Printf("%s", colName)

		// output a delimiter after each header EXCEPT for the last one
		if i < len(cols)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

/********************************************************************************/
func outputData(rows *sql.Rows) {
	cols, _ := rows.Columns()
	data := make(map[string]string)

	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
			fmt.Printf("%s%s%s", config.enclosedBy, data[colName], config.enclosedBy)

			// output a delimiter after each field EXCEPT for the last one
			if i < len(cols)-1 {
				fmt.Printf("%s", config.delimiter)
			}
		}

		fmt.Println()

	}
}

/********************************************************************************/
func main() {
	parseFlags()
	setDebug(config.debugMode)
	loadConfig(config.configFile)
	loadConnection(config.connectionConfig)

	debugConfig()
	checkMinFlags()

	if connection.password == "" {
		connection.password = string(getPassword())
	}

	db, err := sql.Open("goracle", getConnectionString(connection))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var query string

	// check if object starts with the word 'select'
	// otherwise we'll assume it's an object name
	if strings.HasPrefix(config.export, "select") {
		query = config.export
	} else {
		query = fmt.Sprintf("select * from %s", config.export)
	}

	var sql = fmt.Sprintf(query)
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	cols, _ := rows.Columns()

	if config.noheaders != true {
		outputHeaders(cols)
	}

	outputData(rows)
}
