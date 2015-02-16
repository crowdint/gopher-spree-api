package json

import (
  "github.com/crowdint/gopher-spree-api/domain/json"
  "github.com/crowdint/gopher-spree-api/domain/models"
  "github.com/crowdint/gopher-spree-api/interfaces/repositories"
  "github.com/jinzhu/copier"
  "github.com/davecgh/go-spew/spew"
)

type TaxonInteractor struct {
  Repo *repositories.TaxonRepo
}

func NewTaxonInteractor() *TaxonInteractor {
  return &TaxonInteractor{
    Repo: repositories.NewTaxonRepo(),
  }
}

type JsonTaxonsMap map[int64][]*json.Taxon

func (this *TaxonInteractor) GetJsonTaxonsMap(variantIds []int64) (JsonTaxonsMap, error) {

  taxons, err := this.Repo.FindByTaxonomyIds(variantIds)
  if err != nil {
    return JsonTaxonsMap{}, err
  }

  this.AddChildren(taxons)
  //taxonsJson := this.AddChildren(taxons)
  //taxonsJson := this.modelsToJsonTaxonsMap(taxons)

  return JsonTaxonsMap{}, nil
}

//func (this *TaxonInteractor) modelsToJsonTaxonsMap(taxonSlice []*models.Taxon) JsonTaxonsMap {
  //jsonTaxonsMap := JsonTaxonsMap{}

  //for _, taxon := range taxonSlice {
    //taxonJson := &json.Taxon{}
    //copier.Copy(taxonJson, taxon)

    //if _, exists := jsonTaxonsMap[taxon.VariantId]; !exists {
      //jsonTaxonsMap[taxon.VariantId] = []*json.Taxon{}
    //}

    //jsonTaxonsMap[taxon.VariantId] = append(jsonTaxonsMap[taxon.VariantId], taxonJson)

  //}

  //return jsonTaxonsMap
//}

func (this *TaxonInteractor) AddChildren(models []*models.Taxon) []*json.Taxon {
  all := []*json.Taxon{}
  for _, node := range models {
    jsonNode := &json.Taxon{}
    copier.Copy(jsonNode, node)

    for _, child := range models {
      this.AddChild(jsonNode, child)
    }

    all = append(all, jsonNode)
  }
  spew.Printf("All: %+v", all)
  return all
}

func (this *TaxonInteractor) AddChild(node *json.Taxon, child *models.Taxon) {
  if child.ParentId == node.Id {
    jsonChild := json.Taxon{}
    copier.Copy(jsonChild, child)
    node.Taxons = append(node.Taxons, jsonChild)
  }
}
