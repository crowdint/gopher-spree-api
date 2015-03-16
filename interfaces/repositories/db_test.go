package repositories

import (
	"fmt"
	"testing"

	"github.com/crowdint/gopher-spree-api/domain"
)

func TestDB(t *testing.T) {
	err := InitDB(true)

	defer ResetDB()

	if err != nil {
		t.Error("An error has ocurred", err)
	}

	if Spree_db == nil {
		t.Error("Database helper not initialized")
	}
}

func TestAll(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	p := &[]domain.Product{}

	Spree_db.Create(p)

	err := r.All(p, map[string]interface{}{"limit": 10, "offset": 1}, nil)

	if err != nil {
		t.Errorf("DB.All %s", err)
	}

	if len(*p) == 0 {
		t.Errorf("DB.All should have found results")
	}
}

func TestAllWithConditions(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	p := &[]domain.Product{}

	tmpl := `INSERT INTO spree_products(id, name, description, available_on, deleted_at, slug, meta_description, meta_keywords, tax_category_id, shipping_category_id, created_at, updated_at, promotionable, meta_title) VALUES(%s)`

	sql1 := fmt.Sprintf(tmpl, `10,'Spree Ringer T-Shirt','Labore ut sint neque exercitationem aliquid consequuntur ea dolores.Quo asperiores eligendi ipsam officia.Autem aliquid temporibus est blanditiis','2015-02-24 17:57:13.788353',null,'spree-ringer-t-shirt',null,null,1,1,'2015-02-24 17:57:15.214292','2015-02-24 17:57:39.946429','t',null`)
	sql2 := fmt.Sprintf(tmpl, `13, 'Ruby on Rails Mug','Labore ut sint neque exercitationem aliquid consequuntur ea dolores.Quo asperiores eligendi ipsam officia.Autem aliquid temporibus est blanditiis.','2015-02-24 17:57:13.788353',null,'ruby-on-rails-mug',null,null,null,1,'2015-02-24 17:57:15.518985','2015-02-24 17:57:33.982174','t',null`)

	Spree_db.Exec(sql1)
	Spree_db.Exec(sql2)

	err := r.All(p, map[string]interface{}{"limit": 10, "offset": 1}, map[string]interface{}{"id": 1})

	if err != nil {
		t.Errorf("DB.All %s", err)
	}

	if len(*p) > 1 {
		t.Errorf("DB.All should not have found more than one result")
	}
}

func TestFindBy(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	price := 12.0
	product := &domain.Product{Name: "Test Product", Price: &price, ShippingCategoryId: 1, Slug: "test-product"}

	Spree_db.Create(product)

	err := r.FindBy(product, nil, nil)

	if err != nil {
		t.Errorf("DB.FindBy %s", err.Error())
	}
}

func TestFindByWithConditions(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	p := &domain.Product{}

	productPrice := 12.0
	Spree_db.Create(&domain.Product{Id: 1, Name: "Test Product", Price: &productPrice, ShippingCategoryId: 1, Slug: "test-product"})

	err := r.FindBy(p, nil, map[string]interface{}{"id": 1})

	if err != nil {
		t.Errorf("DB.FindBy %s", err.Error())
	}
}

func TestFindByWithOptions(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	p := &domain.Product{}

	productPrice := 12.0
	Spree_db.Create(&domain.Product{Id: 1, TaxCategoryId: 1, Name: "Test Product", Price: &productPrice, ShippingCategoryId: 1, Slug: "test-product"})

	err := r.FindBy(p, map[string]interface{}{
		"not": Not{Key: "tax_category_id", Values: []interface{}{0}},
	}, nil)

	if err != nil {
		t.Errorf("DB.FindBy %s", err.Error())
	}
}

func TestCreate(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	productPrice := 12.0
	p := &domain.Product{Name: "Test Product", Price: &productPrice, ShippingCategoryId: 1, Slug: "test-product"}

	err := r.Create(p)
	if err != nil {
		t.Errorf("DB.Create %s", err.Error())
	}

	if p.Id == 0 {
		t.Errorf("The product was not created in the DB.")
	}
}

func TestFirstOrCreate(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	productPrice := 12.0
	p := &domain.Product{Price: &productPrice, ShippingCategoryId: 1, Slug: "test-product"}

	err := r.FirstOrCreate(p, map[string]interface{}{"name": "Test Product"})
	if err != nil {
		t.Errorf("DB.FirstOrCreate %s", err.Error())
	}

	if p.Id == 0 {
		t.Errorf("The product was not created in the DB.")
	}

	if p.Name != "Test Product" {
		t.Errorf("The product's name should be Test Product, but was %s.", p.Name)
	}
}

func TestCreateWithSlug(t *testing.T) {
	InitDB(true)

	defer ResetDB()

	r := NewDatabaseRepository()

	productPrice := 12.0
	v := &domain.Variant{Sku: "test"}
	p := &domain.Product{Name: "Test Product", Price: &productPrice, ShippingCategoryId: 1, Slug: "test-product", Master: v}

	err := r.CreateWithSlug(p)
	if err != nil {
		t.Errorf("DB.CreateWithSlug %s", err.Error())
	}

	if p.Id == 0 {
		t.Errorf("The product was not created in the DB.")
	}

	if p.Slug == "" {
		t.Errorf("The product slug was not set correctly.")
	}
}
