package system

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/roporter-mydev/auth/jwt"
	"bitbucket.org/roporter-mydev/auth/store"
	"bitbucket.org/roporter-mydev/stats"
	"github.com/gorilla/sessions"

	"github.com/boltdb/bolt"
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
	mgo "gopkg.in/mgo.v2"
)

func (application *Application) Startup() {
}

func (application *Application) LoadConfig(locations []string) {
	application.Configuration = LoadConfig("config", locations)
}

func (application *Application) Init() {
	application.Settings = &Settings{Count: 0}
	application.Settings.Stats = stats.New()
	application.Store = sessions.NewCookieStore([]byte(application.Configuration.GetString("Secret")))
}

func (application *Application) DisplayIntro() {
	ClearTerminal()
	configRuntime(application)
}

func (application *Application) SetupSecurity() {
	timeout, _ := strconv.Atoi(application.Configuration.GetString("certificate.Timeout"))
	options := jwt.Options{
		SigningMethod: application.Configuration.GetString("certificate.signing"),
		PrivateKey:    ReadFileAsString(application.Configuration.GetString("certificate.Private")), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     ReadFileAsString(application.Configuration.GetString("certificate.Public")),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    time.Duration(timeout) * time.Minute,
	}
	application.DBOptions = &options
}

func (application *Application) LoadTemplates() error {
	var templates []string
	fn := func(path string, f os.FileInfo, err error) error {

		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	var err error
	if Exists(application.Configuration.GetString("TemplatePath")) {
		err = filepath.Walk(application.Configuration.GetString("TemplatePath"), fn)
	} else {
		panic("TEMPLATE DIRECTORY DOES NOT EXIST")
		os.Exit(1)
	}

	if err != nil {
		return err
	}

	application.Template = template.Must(template.ParseFiles(templates...))
	return nil
}

func (application *Application) ConnectToDatabase() {
	if application.Configuration.GetString("Database.Type") == "bolt" {
		application.ConnectToBoltDatabase()
	} else if application.Configuration.GetString("Database.Type") == "ledis" {
		application.ConnectToLedisDatabase()
	} else if application.Configuration.GetString("Database.Type") == "mongo" {
		application.ConnectToMongoDatabase()
	}
}

func (application *Application) ConnectToMongoDatabase() {
	session, err := mgo.Dial(application.Configuration.GetString("Database.Hosts"))

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, application.Configuration.GetBool("Database.Mongo.Monotonic"))
	application.Mongo.UserDBStore = store.NewMongoStore(session.DB(application.Configuration.GetString("Database.Mongo.DBName")).C(application.Configuration.GetString("Database.Mongo.UsersTable")))

	_, err = application.createDefaultAdminUser()
	if err == nil {

	}
}

func (application *Application) ConnectToLedisDatabase() {
	//create := false

	//TEMP
	cfg := lediscfg.NewConfigDefault()
	cfg.DBPath = application.Configuration.GetString("Database.Ledis.Path")
	cfg.Databases = application.Configuration.GetInt("Database.Ledis.Count")
	cfg.DBName = application.Configuration.GetString("Database.Ledis.Backend")
	cfg.ConnReadBufferSize = application.Configuration.GetInt("Database.Ledis.readbuffer")
	cfg.ConnWriteBufferSize = application.Configuration.GetInt("Database.Ledis.writebuffer")
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)

	application.Ledis.DB = l
	application.Ledis.UserDBStore = store.NewLedisStore(db)

	_, err := application.createDefaultAdminUser()
	if err == nil {

	}
}

func (application *Application) ConnectToBoltDatabase() {
	create := false
	path := application.Configuration.GetString("Database.Bolt.Path") + "/" + application.Configuration.GetString("Database.Name")
	if !Exists(path) {
		create = true
	}
	db, err := bolt.Open(path, 0600, &bolt.Options{})
	if err != nil {
		//glog.Fatalf("Can't connect to or create the database: %v", err)
		panic(err)
	}
	application.Bolt.DB = db
	store2, err2 := store.NewBoltStore(application.Bolt.DB, "users")
	application.Bolt.UserDBStore = store2
	if err2 != nil {
		panic("Can not create bolt store")
	}
	if create {
		application.createDefaultAdminUser()
	}
}

func (application *Application) createDefaultAdminUser() (string, error) {
	if application.Configuration.GetString("Database.Type") == "bolt" {
		return application.Bolt.UserDBStore.Signin(application.Configuration.GetString("Admin.Email"), application.Configuration.GetString("Admin.Password"), application.Configuration.GetStringSlice("Admin.Scopes"))
	} else if application.Configuration.GetString("Database.Type") == "ledis" {
		found, email, err := application.Ledis.UserDBStore.AdminExists(application.Configuration.GetString("Admin.Email"))
		if found {
			return email, err
		} else {
			return application.Ledis.UserDBStore.Signin(application.Configuration.GetString("Admin.Email"), application.Configuration.GetString("Admin.Password"), application.Configuration.GetStringSlice("Admin.Scopes"))
		}
	} else if application.Configuration.GetString("Database.Type") == "mongo" {
		//found, email, err := application.Mongo.UserDBStore.AdminExists(application.Configuration.GetString("Admin.Email"))
		//if found {
		//	return email, err
		//} else {
		return application.Mongo.UserDBStore.Signin(application.Configuration.GetString("Admin.Email"), application.Configuration.GetString("Admin.Password"), application.Configuration.GetStringSlice("Admin.Scopes"))
		//}
	}
	return "", errors.New("")
}

func (application *Application) Close() {
	fmt.Println("Terminated Application, connections and sessions.")
	if application.Configuration.GetString("Database.Type") == "bolt" {
		application.Bolt.DB.Close()
	} else if application.Configuration.GetString("Database.Type") == "ledis" {
		application.Ledis.DB.Close()
	} else if application.Configuration.GetString("Database.Type") == "mongo" {

	}
}
