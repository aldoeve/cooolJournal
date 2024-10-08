package setup


import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "path/filepath"
)

var(
folderPath = "../database/"
fileName = "database.db"
schemaName = "schema.sql"
filePath = filepath.Join(folderPath, fileName)
schemaPath = filepath.Join(folderPath, schemaName)
)


func SetupDefault(flags map[string]bool) {
    //The database *does* exists
    if _, err := os.Stat(filePath); err == nil {
        fmt.Print("database exists...\n")

        //Reset database from flag
        if flags["resetdb"] {
            Resetdb()
        }

    } else { //The database does *not* exist
        fmt.Print("database does *not* exist...\n")
        Createdb()
    }

}

func Resetdb() {
    //Attempts to reset the database
    fmt.Print("reseting database...\n")
    err := os.Remove(filePath)
    fmt.Print("deleting old database...\n")
    if err != nil {
        panic(err)
    }

    Createdb()
    
}

func Createdb() {
    //Attempts to create the database
    fmt.Print("creating new database...\n")
    file, err := os.Create(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()


    //Attempts to open the database
    fmt.Print("opening database...\n")
    db, err := sql.Open("sqlite3", filePath)
    if err != nil {
        panic(err)
    }

    //Enter schema into database
    fmt.Print("opening schema...\n")
    fileContents, err := os.ReadFile(schemaPath)
    if err != nil {
        panic(err)
    }
    _, err = db.Exec(string(fileContents))

    if err != nil {
		fmt.Println("Database schema entered unsuccessfully...")
        panic(err)
	} else {
		fmt.Println("Database schema entered successfully...")
	}

    defer db.Close()
}