package kv

import (
	"omniql.omniql.io/application"
)

function main(){
application.HandleCreate(func (*application.Aplication{}) (*application.CreateSucess,  application.CreatePipeline) ) {
	store, error := libkv.NewStore()
	// check if obeject does not exist already
	err:= store.Put(oapapa.GetId(), oapapa.getbytes(), nil)
if err{
application.FailNext
}
Succes(ddas)
})

}



