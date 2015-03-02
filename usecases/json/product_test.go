package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductInteractor_GetMergedResponse(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	tmpl := `INSERT INTO spree_products(id, name, description, available_on, deleted_at, slug, meta_description, meta_keywords, tax_category_id, shipping_category_id, created_at, updated_at, promotionable, meta_title) VALUES(%s)`

	sql1 := fmt.Sprintf(tmpl, `1,'Spree Ringer T-Shirt','Labore ut sint neque exercitationem aliquid consequuntur ea dolores.Quo asperiores eligendi ipsam officia.Autem aliquid temporibus est blanditiis','2015-02-24 17:57:13.788353',null,'spree-ringer-t-shirt',null,null,1,1,'2015-02-24 17:57:15.214292','2015-02-24 17:57:39.946429','t',null`)
	sql2 := fmt.Sprintf(tmpl, `2, 'Ruby on Rails Mug','Labore ut sint neque exercitationem aliquid consequuntur ea dolores.Quo asperiores eligendi ipsam officia.Autem aliquid temporibus est blanditiis.','2015-02-24 17:57:13.788353',null,'ruby-on-rails-mug',null,null,null,1,'2015-02-24 17:57:15.518985','2015-02-24 17:57:33.982174','t',null`)

	repositories.Spree_db.Exec(sql1)
	repositories.Spree_db.Exec(sql2)

	productInteractor := NewProductInteractor()

	jsonProductSlice, err := productInteractor.GetResponse(1, 10, &FakeResponseParameters{})
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if jsonProductSlice.(ContentResponse).GetCount() < 1 {
		t.Error("Error: Invalid number of rows")
		return
	}

	jsonBytes, err := json.Marshal(jsonProductSlice)
	if err != nil {
		t.Error("Error: An error has ocurred:", err.Error())
	}

	if string(jsonBytes) == "" {
		t.Error("Error: Json string is empty")
	}
}

func TestProductInteractor_getIdSlice(t *testing.T) {
	products := []*domain.Product{
		&domain.Product{
			Id: 1,
		},
		&domain.Product{
			Id: 2,
		},
		&domain.Product{
			Id: 3,
		},
	}

	productInteractor := NewProductInteractor()

	ids := productInteractor.getIdSlice(products)

	if len(ids) != 3 {
		t.Error("Incorrect number of ids")
	}

	if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Error("Incorrect id value")
	}
}

func TestProductInteractor_mergeVariants(t *testing.T) {
	jsonProductSlice := []*domain.Product{
		&domain.Product{
			Id:   99991,
			Name: "Product1",
		},
		&domain.Product{
			Id:   99992,
			Name: "Product2",
		},
	}

	jsonVariantsMap := JsonVariantsMap{
		99991: []*domain.Variant{
			{
				Id: 99991,
			},
		},
		99992: []*domain.Variant{
			{
				Id:       99992,
				IsMaster: true,
			},
		},
	}

	productInteractor := NewProductInteractor()

	productInteractor.mergeVariants(jsonProductSlice, jsonVariantsMap)

	p2 := jsonProductSlice[0]

	if p2.Variants == nil {
		t.Error("Product variants are nil")
		return
	}

	if len(p2.Variants) == 0 {
		t.Error("No product variants found")
		return
	}

	v1 := p2.Variants[0]

	if v1.Id != 99991 || v1.Name != "Product1" || v1.IsMaster {
		t.Errorf("Incorrect variant values %d %s %b", v1.Id, v1.Name, v1.IsMaster)
	}
}

func TestProductInteractor_mergeOptionTypes(t *testing.T) {
	jsonProductSlice := []*domain.Product{
		&domain.Product{
			Id: 3,
		},
	}

	jsonOptionTypesMap := JsonOptionTypesMap{
		3: []*domain.OptionType{
			{
				Id:           1,
				Name:         "tshirt-size",
				Presentation: "Size",
			},
			{
				Id:           2,
				Name:         "tshirt-color",
				Presentation: "Color",
			},
		},
	}

	productInteractor := NewProductInteractor()

	productInteractor.mergeOptionTypes(jsonProductSlice, jsonOptionTypesMap)

	product := jsonProductSlice[0]

	if product.OptionTypes == nil {
		t.Error("Product OptionTypes are nil")
		return
	}

	if len(product.OptionTypes) == 0 {
		t.Error("No product optionTypes found")
		return
	}

	optionType1 := product.OptionTypes[0]

	if optionType1.Id != 1 || optionType1.Name != "tshirt-size" || optionType1.Presentation != "Size" {
		t.Errorf("Incorrect optionType values: \n Id -> %d, Name -> %s, Presentation -> %d", optionType1.Id, optionType1.Name, optionType1.Presentation)
	}
}

func TestProductInteractor_mergeClassifications(t *testing.T) {
	jsonProductSlice := []*domain.Product{
		&domain.Product{
			Id: 3,
		},
		&domain.Product{
			Id: 5,
		},
	}

	jsonOptionTypesMap := JsonClassificationsMap{
		3: []*domain.Classification{
			{
				TaxonId:  1,
				Position: 5,
				Taxon: domain.Taxon{
					Id:   1,
					Name: "taxonName",
				},
			},
		},
	}

	productInteractor := NewProductInteractor()

	productInteractor.mergeClassifications(jsonProductSlice, jsonOptionTypesMap)

	product1 := jsonProductSlice[0]
	product2 := jsonProductSlice[1]

	if product1.Classifications == nil || product2.Classifications == nil {
		t.Error("Product.Classifications should be and empty slice [] at least")
	}

	classification := product1.Classifications[0]

	if classification.TaxonId != 1 || classification.Taxon.Id != 1 {
		t.Error("Wrong assignment of classifications")
	}

	if len(product2.Classifications) > 0 {
		t.Error("Wrong assignment of classficiations")
	}

	if product1.TaxonIds[0] != 1 {
		t.Error("Wrong assignment of taxon ids")
	}

	if len(product2.TaxonIds) != 0 {
		t.Error("Wrong assignment of taxon ids")
	}

}
