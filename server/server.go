package main

import (
	"github.com/karthik-Prathipati/learningProto/proto/numanipb"
)

type server struct {
	numanipb.UnimplementedNumberManipulationServer
}
