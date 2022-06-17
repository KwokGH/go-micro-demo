package model

type Product struct {
	ID    int64          `json:"id" gorm:"id;primaryKey;autoIncrement;not null"`
	Name  string         `json:"name" gorm:"name"`
	Sku   string         `json:"sku" gorm:"sku;unique_index:idx_sku;not null"`
	Price float32        `json:"price" gorm:"price"`
	Desc  string         `json:"desc" gorm:"desc"`
	Image []ProductImage `json:"image" gorm:"foreignKey:ProductID"`
	Size  []ProductSize  `json:"size" gorm:"foreignKey:ProductID"`
	Seo   ProductSeo     `json:"seo" gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "product"
}

type ProductImage struct {
	ID        int64  `json:"id" gorm:"id;primaryKey;autoIncrement;not null"`
	Name      string `json:"name" gorm:"name"`
	Code      string `json:"code" gorm:"code;unique_index:idx_code;not null"`
	Url       string `json:"url" gorm:"url"`
	ProductID int64  `json:"product_id"`
}

func (ProductImage) TableName() string {
	return "product_image"
}

type ProductSize struct {
	ID        int64  `json:"id" gorm:"id;primaryKey;autoIncrement;not null"`
	Name      string `json:"name" gorm:"name"`
	Code      string `json:"code" gorm:"code;unique_index:idx_code;not null"`
	ProductID int64  `json:"product_id"`
}

func (ProductSize) TableName() string {
	return "product_size"
}

type ProductSeo struct {
	ID        int64  `json:"id" gorm:"id;primaryKey;autoIncrement;not null"`
	Title     string `json:"title" gorm:"title"`
	Keywords  string `json:"keywords" gorm:"keywords"`
	Desc      string `json:"desc" gorm:"desc"`
	Code      string `json:"code" gorm:"code;unique_index:idx_code;not null"`
	ProductID int64  `json:"product_id"`
}

func (ProductSeo) TableName() string {
	return "product_seo"
}
