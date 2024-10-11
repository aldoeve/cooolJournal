package setup


import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "path/filepath"
    "encoding/json"
    "io/ioutil"
)

var(
folderPath = "../database/"
fileName = "database.db"
schemaName = "schema.sql"
demoDataName = "demo.json" 
filePath = filepath.Join(folderPath, fileName)
schemaPath = filepath.Join(folderPath, schemaName)
demoDataPath = filepath.Join(folderPath,demoDataName)

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

    LoadDemoData()
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
		fmt.Println("*Database schema entered unsuccessfully...*")
        panic(err)
	} else {
		fmt.Println("Database schema entered successfully...")
	}

    defer db.Close()
}

func LoadDemoData() {
    // Open the JSON file
    jsonFile, err := os.Open(demoDataPath)
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully opened demo JSON file...")
    defer jsonFile.Close()

    //Read the JSON file
    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result struct {
        Users   []struct {
            Email    string `json:"email"`
            Username string `json:"username"`
            Password string `json:"password"`
            Terms    string `json:"terms"`
            Bio      string `json:"bio"`
        } `json:"users"`
        Avatar  []struct {
            UserID int    `json:"user_id"`
            Avatar string `json:"avatar"`
        } `json:"avatar"`
        Topics []struct {
            TopicName string `json:"topic_name"`
        } `json:"topics"`
        Journals []struct {
            UserID  int `json:"user_id"`
            TopicID int `json:"topic_id"`
            JournalName string `json:"journal_name"`
        } `json:"journals"`
        Pages []struct {
            JournalID int    `json:"journal_id"`
            PageEntry string `json:"page_entry"`
            PageDate  string `json:"page_date"`
        } `json:"pages"`
    }

    //Parse json data
    err = json.Unmarshal(byteValue, &result)
    if err != nil {
        panic(err)
    }

    fmt.Print("opening database...\n")
    db, err := sql.Open("sqlite3", filePath)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Enters data into database
    for _, user := range result.Users {
        _, err := db.Exec("INSERT INTO users (email, username, password, terms, bio) VALUES (?, ?, ?, ?, ?)",
            user.Email, user.Username, user.Password, user.Terms, user.Bio)
        if err != nil {
            fmt.Println("*user data *not* entered*")
        }
    }

    for _, avatar := range result.Avatar {
        _, err := db.Exec("INSERT INTO avatar (user_id, avatar) VALUES (?, ?)",
            avatar.UserID, avatar.Avatar)
        if err != nil {
            fmt.Println("*avatar data *not* entered*")
        }
    }

    for _, topic := range result.Topics {
        _, err := db.Exec("INSERT INTO topics (topic_name) VALUES (?)",
            topic.TopicName)
        if err != nil {
            fmt.Println("*topic data *not* entered*")
        }
    }

    for _, journal := range result.Journals {
        _, err := db.Exec("INSERT INTO journals (user_id, topic_id, journal_name) VALUES (?, ?, ?)",
            journal.UserID, journal.TopicID, journal.JournalName)
        if err != nil {
            fmt.Println("*journal data *not* entered*")
        }
    }

    for _, page := range result.Pages {
        _, err := db.Exec("INSERT INTO pages (journal_id, page_entry, page_date) VALUES (?, ?, ?)",
            page.JournalID, page.PageEntry, page.PageDate)
        if err != nil {
            fmt.Println("*page data *not* entered*")
        }
    }
    fmt.Println("Entered correctly formated data into database...")

}