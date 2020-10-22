package operationtypes

import "github.com/Tsuryu/arreglapp-core-operations/app/db"

var database = db.Connection.Database("arreglapp")

// Collection : collection database mongo name
var Collection = database.Collection("operation_types")
