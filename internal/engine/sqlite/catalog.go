package sqlite

import "github.com/StanVerse/sqld/internal/sql/catalog"

func NewCatalog() *catalog.Catalog {
	def := "main"
	return &catalog.Catalog{
		DefaultSchema: def,
		Schemas: []*catalog.Schema{
			defaultSchema(def),
		},
		Extensions: map[string]struct{}{},
	}
}

func newTestCatalog() *catalog.Catalog {
	return catalog.New("main")
}
