// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/central/globaldb"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableK8sRolesStmt holds the create statement for table `k8s_roles`.
	CreateTableK8sRolesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists k8s_roles (
                   Id varchar,
                   Name varchar,
                   Namespace varchar,
                   ClusterId varchar,
                   ClusterName varchar,
                   ClusterRole bool,
                   Labels jsonb,
                   Annotations jsonb,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}

	// K8sRolesSchema is the go schema for table `k8s_roles`.
	K8sRolesSchema = func() *walker.Schema {
		schema := globaldb.GetSchemaForTable("k8s_roles")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.K8SRole)(nil)), "k8s_roles")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_ROLES, "k8srole", (*storage.K8SRole)(nil)))
		globaldb.RegisterTable(schema)
		return schema
	}()
)