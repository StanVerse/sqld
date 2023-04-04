// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/StanVerse/sqld/internal/sql/ast"
	"github.com/StanVerse/sqld/internal/sql/catalog"
)

var funcsCube = []*catalog.Function{
	{
		Name: "cube",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision[]"},
			},
			{
				Type: &ast.TypeName{Name: "double precision[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube_cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "cube_contained",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_contains",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_coord",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "cube_coord_llur",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "cube_dim",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "cube_distance",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "cube_enlarge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube_eq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_ge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_gt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube_inter",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube_is_point",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_le",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_ll_coord",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "cube_lt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_ne",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "cube_overlap",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "cube_send",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "cube_size",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "cube_subset",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube_union",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "cube_ur_coord",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "distance_chebyshev",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "distance_taxicab",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cube"},
			},
			{
				Type: &ast.TypeName{Name: "cube"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
}

func Cube() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsCube
	return s
}
