package json

//import (
//"encoding/json"
//"testing"

//"github.com/crowdint/gopher-spree-api/domain/models"
//"github.com/crowdint/gopher-spree-api/interfaces/repositories"
//)

//func TestTaxonInteractor_GetResponse(t *testing.T) {
//err := repositories.InitDB(true)
//if err != nil {
//t.Error("Error: An error has ocurred:", err.Error())
//}

//defer repositories.Spree_db.Close()

//taxonInteractor := NewTaxonInteractor()

//jsonTaxonSlice, err := taxonInteractor.GetResponse(1, 10, &FakeResponseParameters{})
//if err != nil {
//t.Error("Error: An error has ocurred:", err.Error())
//}

//if jsonTaxonSlice.(ContentResponse).GetCount() < 1 {
//t.Error("Error: Invalid number of rows")
//return
//}

//jsonBytes, err := json.Marshal(jsonTaxonSlice)
//if err != nil {
//t.Error("Error: An error has ocurred:", err.Error())
//}

//if string(jsonBytes) == "" {
//t.Error("Error: Json string is empty")
//}
//}
