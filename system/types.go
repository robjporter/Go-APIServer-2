package system

import (
	"html/template"

	"bitbucket.org/roporter-mydev/auth/jwt"
	"bitbucket.org/roporter-mydev/auth/store"
	"bitbucket.org/roporter-mydev/stats"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/gorilla/sessions"
	"github.com/siddontang/ledisdb/ledis"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"github.com/blang/semver"
)

type Module struct {
	Path string
	Name string
}

type CsrfProtection struct {
	Key    string
	Cookie string
	Header string
	Secure bool
}

type Settings struct {
	Count int
	Stats *stats.Stats
}

type BoltStruct struct {
	DB          *bolt.DB
	UserDBStore *store.BoltStore
}

type LedisStruct struct {
	DB          *ledis.Ledis
	UserDBStore *store.LedisStore
}

type MongoStruct struct {
	DB          *mgo.Database
	UserDBStore *store.MongoStore
}

type VersionStruct struct {
	AppVersion     *semver.Version
	BuildRevision  string
	HeaderRevision string
	PluginVersions []string
}

type Application struct {
	Configuration  *viper.Viper
	Template       *template.Template
	Templates      []string
	Store          *sessions.CookieStore
	Versions       VersionStruct
	Ledis          LedisStruct
	Mongo          MongoStruct
	Bolt           BoltStruct
	DBOptions      *jwt.Options
	Settings       *Settings
	Log            *logrus.Logger
	CsrfProtection *CsrfProtection
}
