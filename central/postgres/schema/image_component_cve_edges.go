// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	"github.com/stackrox/rox/central/globaldb"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableImageComponentCveEdgesStmt holds the create statement for table `image_component_cve_edges`.
	CreateTableImageComponentCveEdgesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists image_component_cve_edges (
                   Id varchar,
                   IsFixable bool,
                   FixedBy varchar,
                   ImageComponentId varchar,
                   ImageCveId varchar,
                   serialized bytea,
                   PRIMARY KEY(Id, ImageComponentId, ImageCveId),
                   CONSTRAINT fk_parent_table_0 FOREIGN KEY (ImageComponentId) REFERENCES image_components(Id) ON DELETE CASCADE
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}

	// ImageComponentCveEdgesSchema is the go schema for table `image_component_cve_edges`.
	ImageComponentCveEdgesSchema = func() *walker.Schema {
		schema := globaldb.GetSchemaForTable("image_component_cve_edges")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComponentCVEEdge)(nil)), "image_component_cve_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.ImageComponent": ImageComponentsSchema,
			"storage.CVE":            ImageCvesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPONENT_VULN_EDGE, "componentcveedge", (*storage.ComponentCVEEdge)(nil)))
		globaldb.RegisterTable(schema)
		return schema
	}()
)