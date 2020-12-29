package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
)
 
// Disease schema.
type Disease struct {
   ent.Schema
}
 
// Fields of the Disease.
func (Disease) Fields() []ent.Field {
   return []ent.Field{
       field.String("DiseaseName"),
	   
   }
}
 
// Edges of the Disease.
func (Disease) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("diagnosis", Diagnosis.Type),
    }
}
 
