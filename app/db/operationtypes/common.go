package operationtypes

import "github.com/Tsuryu/arreglapp-commons/app/db"

var database = db.Connection.Database("arreglapp")

// Collection : collection database mongo name
var Collection = database.Collection("operation_types")
