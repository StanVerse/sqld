package compiler

import (
	"github.com/StanVerse/sqld/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
