package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"

	"fmt"
	"io"
	// "mime/multipart"
	"path/filepath"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	// "math/rand/v2" // For Go 1.22+
	// OR
	"math/rand"  // For Go 1.20-1.21
	// "time"
)

type ProductService interface {
	CreateProduct(c echo.Context) error
	GetProductById(id string) (*dto.ProductResponse, error)
	GetProductByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error)
	GetAllProduct(desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error)
	UpdateProduct(c echo.Context, id string) error
	DeleteProduct(id string) error
	CheckProductStore(id string, c echo.Context) error
	AdvertiseProduct(c echo.Context, id string) error
	UnadvertiseProduct(c echo.Context, id string) error
	GetAllProductAdvertisement(desc, page, pageSize int, search, sort string) ([]models.Product, *dto.Pagination, error)
	GetAllProductAdvertisementByStoreId(id string, desc, page, pageSize int, search, sort string) ([]models.Product, *dto.Pagination, error)
}

type productService struct {
	productRepo  repository.ProductRepository
	uploader     *ClientUploader
	storeService StoreService
	ptService    ProductTypeService
	tokenRepo    repository.TokenRepository
}

func NewProductService() ProductService {
	return &productService{
		productRepo:  repository.NewProductRepositoryGORM(),
		uploader:     NewClientUploader(),
		storeService: NewStoreService(),
		tokenRepo:    repository.NewTokenRepositoryGORM(),
		ptService:    NewProductTypeService(),
	}
}

func ParseProductForm(c echo.Context) (*dto.ProductRegisterForm, error) {
	productName := c.FormValue("product_name")
	productPriceStr := c.FormValue("product_price")
	productDesc := c.FormValue("product_desc")
	productIsShowStr := c.FormValue("product_isshow")
	productLemakTotalStr := c.FormValue("product_lemaktotal")
	productProteinStr := c.FormValue("product_protein")
	productKarbohidratStr := c.FormValue("product_karbohidrat")
	productGaramStr := c.FormValue("product_garam")
	productGrade := c.FormValue("product_grade")
	productServingSizeStr := c.FormValue("product_servingsize")
	productEnergiStr := c.FormValue("product_energi")
	productGulaStr := c.FormValue("product_gula")
	productSaturatedFatStr := c.FormValue("product_saturatedfat")
	productFiberStr := c.FormValue("product_fiber")
	ptType := c.FormValue("pt_type")

	productPrice, err := strconv.ParseFloat(productPriceStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product price")
	}
	productIsShow, err := strconv.ParseInt(productIsShowStr, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product isShow value")
	}
	productLemakTotal, err := strconv.ParseFloat(productLemakTotalStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product lemak total")
	}
	productProtein, err := strconv.ParseFloat(productProteinStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product protein")
	}
	productKarbohidrat, err := strconv.ParseFloat(productKarbohidratStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product karbohidrat")
	}
	productGaram, err := strconv.ParseFloat(productGaramStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product garam")
	}
	productServingSize, err := strconv.ParseFloat(productServingSizeStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product serving size")
	}
	productEnergi, err := strconv.ParseFloat(productEnergiStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product energi")
	}
	productGula, err := strconv.ParseFloat(productGulaStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product gula")
	}
	productSaturatedFat, err := strconv.ParseFloat(productSaturatedFatStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product saturated fat")
	}
	productFiber, err := strconv.ParseFloat(productFiberStr, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product fiber")
	}

	ptTypeInt, err := strconv.ParseInt(ptType, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product type")
	}

	return &dto.ProductRegisterForm{
		ProductName:         productName,
		ProductPrice:        productPrice,
		ProductDesc:         productDesc,
		ProductIsShow:       productIsShow,
		ProductLemakTotal:   productLemakTotal,
		ProductProtein:      productProtein,
		ProductKarbohidrat:  productKarbohidrat,
		ProductGaram:        productGaram,
		ProductGrade:        productGrade,
		ProductServingSize:  productServingSize,
		ProductEnergi:       productEnergi,
		ProductGula:         productGula,
		ProductSaturatedFat: productSaturatedFat,
		ProductFiber:        productFiber,
		ProductExpShow:      "",
		PT_Type:             ptTypeInt,
	}, nil
}

// func (service *productService) CreateProduct(c echo.Context) error {
// 	productForm, err := ParseProductForm(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	imagePath, err := service.uploader.ProcessImage(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	err = godotenv.Load(".env")
// 	if err != nil {
// 		return err
// 	}
// 	realImagePath := os.Getenv("IMAGE_PATH") + imagePath

// 	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
// 	}

// 	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
// 	}

// 	pt_id, err := service.ptService.GetProductTypeIdByType(productForm.PT_Type)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
// 	}
// 	product_id := uuid.New().String()

// 	product := models.Product{
// 		PRODUCT_ID:           product_id,
// 		PRODUCT_NAME:         productForm.ProductName,
// 		PRODUCT_PRICE:        productForm.ProductPrice,
// 		PRODUCT_DESC:         productForm.ProductDesc,
// 		PRODUCT_ISSHOW:       0,
// 		PRODUCT_LEMAKTOTAL:   productForm.ProductLemakTotal,
// 		PRODUCT_PROTEIN:      productForm.ProductProtein,
// 		PRODUCT_KARBOHIDRAT:  productForm.ProductKarbohidrat,
// 		PRODUCT_GARAM:        productForm.ProductGaram,
// 		PRODUCT_SERVINGSIZE:  productForm.ProductServingSize,
// 		PRODUCT_PICTURE:      realImagePath,
// 		PRODUCT_EXPSHOW:      time.Now(),
// 		PRODUCT_ENERGI:       productForm.ProductEnergi,
// 		PRODUCT_GULA:         productForm.ProductGula,
// 		PRODUCT_SATURATEDFAT: productForm.ProductSaturatedFat,
// 		PRODUCT_FIBER:        productForm.ProductFiber,
// 		CreatedAt:            time.Now(),
// 		UpdatedAt:            time.Now(),
// 		STORE_ID:             store.STORE_ID,
// 		PT_ID:                pt_id,
// 	}

// 	err = service.productRepo.CreateProduct(&product)
// 	if err != nil {
// 		return err
// 	}

// 	url := os.Getenv("PYTHON_API") + "/grade"

// 	requestData := &dto.ProductRequest{
// 		Energy:       product.PRODUCT_ENERGI,
// 		Protein:      product.PRODUCT_PROTEIN,
// 		Fat:          product.PRODUCT_LEMAKTOTAL,
// 		Carbs:        product.PRODUCT_KARBOHIDRAT,
// 		Sugar:        product.PRODUCT_GULA,
// 		Salt:         product.PRODUCT_GARAM,
// 		SaturatedFat: product.PRODUCT_SATURATEDFAT,
// 		Fiber:        product.PRODUCT_FIBER,
// 	}

// 	responseData, _ := SendRequest[dto.ProductRequest, dto.SNResponse](url, *requestData)

// 	p, err := service.productRepo.GetProductById(product_id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
// 	}

// 	p.PRODUCT_GRADING = responseData.Grade
// 	err = service.productRepo.UpdateProduct(p)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
// 	}

// 	return nil
// }

// func (service *productService) CreateProduct(c echo.Context) error {
// 	productForm, err := ParseProductForm(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	// Gunakan ProcessImage untuk menyimpan gambar secara lokal
// 	imagePath, err := service.uploader.ProcessImage(c, service.uploader.productPath)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	// Dapatkan path lokal untuk gambar
// 	realImagePath := imagePath

// 	// Lanjutkan dengan proses lain
// 	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
// 	}

// 	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
// 	}

// 	pt_id, err := service.ptService.GetProductTypeIdByType(productForm.PT_Type)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
// 	}
// 	product_id := uuid.New().String()

// 	product := models.Product{
// 		PRODUCT_ID:           product_id,
// 		PRODUCT_NAME:         productForm.ProductName,
// 		PRODUCT_PRICE:        productForm.ProductPrice,
// 		PRODUCT_DESC:         productForm.ProductDesc,
// 		PRODUCT_ISSHOW:       0,
// 		PRODUCT_LEMAKTOTAL:   productForm.ProductLemakTotal,
// 		PRODUCT_PROTEIN:      productForm.ProductProtein,
// 		PRODUCT_KARBOHIDRAT:  productForm.ProductKarbohidrat,
// 		PRODUCT_GARAM:        productForm.ProductGaram,
// 		PRODUCT_SERVINGSIZE:  productForm.ProductServingSize,
// 		PRODUCT_PICTURE:      realImagePath, // Gambar disimpan secara lokal
// 		PRODUCT_EXPSHOW:      time.Now(),
// 		PRODUCT_ENERGI:       productForm.ProductEnergi,
// 		PRODUCT_GULA:         productForm.ProductGula,
// 		PRODUCT_SATURATEDFAT: productForm.ProductSaturatedFat,
// 		PRODUCT_FIBER:        productForm.ProductFiber,
// 		CreatedAt:            time.Now(),
// 		UpdatedAt:            time.Now(),
// 		STORE_ID:             store.STORE_ID,
// 		PT_ID:                pt_id,
// 	}

// 	err = service.productRepo.CreateProduct(&product)
// 	if err != nil {
// 		return err
// 	}

// 	// Lanjutkan dengan proses lainnya...

	
// 	url := os.Getenv("PYTHON_API") + "/grade"

// 	requestData := &dto.ProductRequest{
// 		Energy:       product.PRODUCT_ENERGI,
// 		Protein:      product.PRODUCT_PROTEIN,
// 		Fat:          product.PRODUCT_LEMAKTOTAL,
// 		Carbs:        product.PRODUCT_KARBOHIDRAT,
// 		Sugar:        product.PRODUCT_GULA,
// 		Salt:         product.PRODUCT_GARAM,
// 		SaturatedFat: product.PRODUCT_SATURATEDFAT,
// 		Fiber:        product.PRODUCT_FIBER,
// 	}

// 	responseData, _ := SendRequest[dto.ProductRequest, dto.SNResponse](url, *requestData)

// 	p, err := service.productRepo.GetProductById(product_id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
// 	}

// 	p.PRODUCT_GRADING = responseData.Grade
// 	err = service.productRepo.UpdateProduct(p)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
// 	}

// 	return nil
// }


// Function to generate a random grade (A/B/C/D/E)
func generateRandomGrade() string {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	// Create a slice with possible grades
	grades := []string{"A", "B", "C", "D", "E"}
	
	// Generate a random index and select a grade using rand.Intn for older Go versions
	randomGrade := grades[rand.Intn(len(grades))]
	
	return randomGrade
}

// func (service *productService) CreateProduct(c echo.Context) error {
// 	productForm, err := ParseProductForm(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	// Validate token and get user info
// 	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
// 	}

// 	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
// 	}

// 	pt_id, err := service.ptService.GetProductTypeIdByType(productForm.PT_Type)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
// 	}

// 	// Upload image
// 	imagePath, err := service.uploader.ProcessImageProduct(c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	product_id := uuid.New().String()
	
// 	product := models.Product{
// 		PRODUCT_ID:           product_id,
// 		PRODUCT_NAME:         productForm.ProductName,
// 		PRODUCT_PRICE:        productForm.ProductPrice,
// 		PRODUCT_DESC:         productForm.ProductDesc,
// 		PRODUCT_ISSHOW:       0,
// 		PRODUCT_LEMAKTOTAL:   productForm.ProductLemakTotal,
// 		PRODUCT_PROTEIN:      productForm.ProductProtein,
// 		PRODUCT_KARBOHIDRAT:  productForm.ProductKarbohidrat,
// 		PRODUCT_GARAM:        productForm.ProductGaram,
// 		PRODUCT_SERVINGSIZE:  productForm.ProductServingSize,
// 		PRODUCT_PICTURE:      imagePath, // Path to locally stored image
// 		PRODUCT_EXPSHOW:      time.Now(),
// 		PRODUCT_ENERGI:       productForm.ProductEnergi,
// 		PRODUCT_GULA:         productForm.ProductGula,
// 		PRODUCT_SATURATEDFAT: productForm.ProductSaturatedFat,
// 		PRODUCT_FIBER:        productForm.ProductFiber,
// 		CreatedAt:            time.Now(),
// 		UpdatedAt:            time.Now(),
// 		STORE_ID:             store.STORE_ID,
// 		PT_ID:                pt_id,
// 	}

// 	err = service.productRepo.CreateProduct(&product)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	// Calculate product grade
// 	url := os.Getenv("PYTHON_API") + "/grade"
// 	requestData := &dto.ProductRequest{
// 		Energy:       product.PRODUCT_ENERGI,
// 		Protein:      product.PRODUCT_PROTEIN,
// 		Fat:          product.PRODUCT_LEMAKTOTAL,
// 		Carbs:        product.PRODUCT_KARBOHIDRAT,
// 		Sugar:        product.PRODUCT_GULA,
// 		Salt:         product.PRODUCT_GARAM,
// 		SaturatedFat: product.PRODUCT_SATURATEDFAT,
// 		Fiber:        product.PRODUCT_FIBER,
// 	}

// 	responseData, err := SendRequest[dto.ProductRequest, dto.SNResponse](url, *requestData)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to calculate product grade")
// 	}

// 	p, err := service.productRepo.GetProductById(product_id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
// 	}

// 	p.PRODUCT_GRADING = responseData.Grade
// 	err = service.productRepo.UpdateProduct(p)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
// 	}

// 	// return c.JSON(http.StatusCreated, map[string]interface{}{
// 	// 	"status":  "success",
// 	// 	"message": "Product created successfully",
// 	// 	"data": map[string]interface{}{
// 	// 		"product_id": product_id,
// 	// 	},
// 	// })

	
// 	// // Initialize random seed - this is important for truly random results
// 	// rand.Seed(time.Now().UnixNano())
	
// 	// // Create a slice with possible grades
// 	// grades := []string{"A", "B", "C", "D", "E"}
	
// 	// // Generate a random index and select a grade using compatible method
// 	// randomGrade := grades[rand.Intn(len(grades))]
	
// 	// // Retrieve the product
// 	// p, err := service.productRepo.GetProductById(product_id)
// 	// if err != nil {
// 	// 	return echo.NewHTTPError(http.StatusNotFound, "Product not found")
// 	// }
	
// 	// // Assign the random grade and update the product
// 	// p.PRODUCT_GRADING = randomGrade
// 	// err = service.productRepo.UpdateProduct(p)
// 	// if err != nil {
// 	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
// 	// }

// 	// return c.JSON(http.StatusCreated, map[string]interface{}{
// 	// 	"status":  "success",
// 	// 	"message": "Product created successfully",
// 	// 	"data": map[string]interface{}{
// 	// 		"product_id": product_id,
// 	// 		"grade":      randomGrade,
// 	// 	},
// 	// })

// 	// return c.JSON(http.StatusCreated, map[string]interface{}{
// 	// 	"status":  "success",
// 	// 	"message": "Product created successfully",
// 	// 	"data": map[string]interface{}{
// 	// 		"product_id": product_id,
// 	// 	},
// 	// })

// 	return nil
// }

func (service *productService) CreateProduct(c echo.Context) error {
	productForm, err := ParseProductForm(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate token and get user info
	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
	}

	pt_id, err := service.ptService.GetProductTypeIdByType(productForm.PT_Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
	}

	// ======= 🖼️ Upload image ke folder publik =======
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Image upload required")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open uploaded file")
	}
	defer src.Close()

	// Simpan ke folder uploads dengan nama unik
	imageID := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	imageName := fmt.Sprintf("%s%s", imageID, ext)
	imagePath := filepath.Join("uploads", imageName)

	dst, err := os.Create(imagePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save image")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to copy image")
	}

	// URL publik (misal untuk konsumsi di Android/Python)
	serverHost := os.Getenv("APP_HOST") // e.g. http://localhost:8080
	publicImageURL := fmt.Sprintf("%s/uploads/%s", serverHost, imageName)

	// ======= 🛒 Simpan data produk =======
	product_id := uuid.New().String()
	product := models.Product{
		PRODUCT_ID:           product_id,
		PRODUCT_NAME:         productForm.ProductName,
		PRODUCT_PRICE:        productForm.ProductPrice,
		PRODUCT_DESC:         productForm.ProductDesc,
		PRODUCT_ISSHOW:       0,
		PRODUCT_LEMAKTOTAL:   productForm.ProductLemakTotal,
		PRODUCT_PROTEIN:      productForm.ProductProtein,
		PRODUCT_KARBOHIDRAT:  productForm.ProductKarbohidrat,
		PRODUCT_GARAM:        productForm.ProductGaram,
		PRODUCT_SERVINGSIZE:  productForm.ProductServingSize,
		PRODUCT_PICTURE:      publicImageURL, // ✅ simpan URL publik
		PRODUCT_EXPSHOW:      time.Now(),
		PRODUCT_ENERGI:       productForm.ProductEnergi,
		PRODUCT_GULA:         productForm.ProductGula,
		PRODUCT_SATURATEDFAT: productForm.ProductSaturatedFat,
		PRODUCT_FIBER:        productForm.ProductFiber,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
		STORE_ID:             store.STORE_ID,
		PT_ID:                pt_id,
	}

	if err := service.productRepo.CreateProduct(&product); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Hitung grade produk dari API Python
	url := os.Getenv("PYTHON_API") + "/grade"
	requestData := &dto.ProductRequest{
		Energy:       product.PRODUCT_ENERGI,
		Protein:      product.PRODUCT_PROTEIN,
		Fat:          product.PRODUCT_LEMAKTOTAL,
		Carbs:        product.PRODUCT_KARBOHIDRAT,
		Sugar:        product.PRODUCT_GULA,
		Salt:         product.PRODUCT_GARAM,
		SaturatedFat: product.PRODUCT_SATURATEDFAT,
		Fiber:        product.PRODUCT_FIBER,
	}

	responseData, err := SendRequest[dto.ProductRequest, dto.SNResponse](url, *requestData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to calculate product grade")
	}

	p, err := service.productRepo.GetProductById(product_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	p.PRODUCT_GRADING = responseData.Grade
	if err := service.productRepo.UpdateProduct(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// return c.JSON(http.StatusCreated, map[string]interface{}{
	// 	"status":  "success",
	// 	"message": "Product created successfully",
	// 	"data": map[string]interface{}{
	// 		"product_id": product_id,
	// 		"image_url":  publicImageURL,
	// 		"grade":      responseData.Grade,
	// 	},
	// })
	return nil
}


func (service *productService) GetProductById(id string) (*dto.ProductResponse, error) {
	p, err := service.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	pt, err := service.ptService.GetProductTypeById(p.PT_ID)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ProductID:           p.PRODUCT_ID,
		ProductName:         p.PRODUCT_NAME,
		ProductPrice:        p.PRODUCT_PRICE,
		ProductDesc:         p.PRODUCT_DESC,
		ProductIsShow:       p.PRODUCT_ISSHOW,
		ProductLemakTotal:   p.PRODUCT_LEMAKTOTAL,
		ProductProtein:      p.PRODUCT_PROTEIN,
		ProductKarbohidrat:  p.PRODUCT_KARBOHIDRAT,
		ProductGaram:        p.PRODUCT_GARAM,
		ProductGrade:        p.PRODUCT_GRADING,
		ProductServingSize:  p.PRODUCT_SERVINGSIZE,
		ProductExpShow:      p.PRODUCT_EXPSHOW.String(),
		ProductPicture:      p.PRODUCT_PICTURE,
		ProductEnergi:       p.PRODUCT_ENERGI,
		ProductGula:         p.PRODUCT_GULA,
		ProductSaturatedFat: p.PRODUCT_SATURATEDFAT,
		ProductFiber:        p.PRODUCT_FIBER,
		PT_Type:             pt.PT_TYPE,
	}, nil
}

func (service *productService) GetProductByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error) {
	return service.productRepo.GetProductByStoreId(id, desc, page, pageSize, search, sort)
}

func (service *productService) GetAllProduct(desc, page, pageSize int, search, sort string) (*[]models.Product, *dto.Pagination, error) {
	return service.productRepo.GetAllProduct(desc, page, pageSize, search, sort)
}

func (service *productService) CheckProductStore(id string, c echo.Context) error {
	UserToken, err := service.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store, err := service.storeService.GetStoreByUserId(UserToken.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Store not found")
	}

	product, err := service.productRepo.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	if product.STORE_ID != store.STORE_ID {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	return nil

}

func (service *productService) UpdateProduct(c echo.Context, id string) error {
	productForm, err := ParseProductForm(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := service.productRepo.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	file, _ := c.FormFile("file")
	if file != nil {
		imagePath, err := service.uploader.ProcessImage(c, service.uploader.productPath)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		service.uploader.DeleteImage(product.PRODUCT_PICTURE,  service.uploader.productPath)
		err = godotenv.Load(".env")
		if err != nil {
			return err
		}
		realImagePath := os.Getenv("IMAGE_PATH") + imagePath
		product.PRODUCT_PICTURE = realImagePath
	}

	if productForm.ProductName != "" {
		product.PRODUCT_NAME = productForm.ProductName
	}
	if productForm.ProductPrice != 0 {
		product.PRODUCT_PRICE = productForm.ProductPrice
	}
	if productForm.ProductDesc != "" {
		product.PRODUCT_DESC = productForm.ProductDesc
	}
	if productForm.ProductLemakTotal != 0 {
		product.PRODUCT_LEMAKTOTAL = productForm.ProductLemakTotal
	}
	if productForm.ProductProtein != 0 {
		product.PRODUCT_PROTEIN = productForm.ProductProtein
	}
	if productForm.ProductKarbohidrat != 0 {
		product.PRODUCT_KARBOHIDRAT = productForm.ProductKarbohidrat
	}
	if productForm.ProductGaram != 0 {
		product.PRODUCT_GARAM = productForm.ProductGaram
	}
	if productForm.ProductGrade != "" {
		product.PRODUCT_GRADING = productForm.ProductGrade
	}
	if productForm.ProductServingSize != 0 {
		product.PRODUCT_SERVINGSIZE = productForm.ProductServingSize
	}
	if productForm.ProductEnergi != 0 {
		product.PRODUCT_ENERGI = productForm.ProductEnergi
	}
	if productForm.ProductGula != 0 {
		product.PRODUCT_GULA = productForm.ProductGula
	}
	if productForm.ProductSaturatedFat != 0 {
		product.PRODUCT_SATURATEDFAT = productForm.ProductSaturatedFat
	}
	if productForm.ProductFiber != 0 {
		product.PRODUCT_FIBER = productForm.ProductFiber
	}

	if productForm.PT_Type != 0 {
		pt_id, err := service.ptService.GetProductTypeIdByType(productForm.PT_Type)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Product type not found")
		}
		product.PT_ID = pt_id
	}

	// err = godotenv.Load(".env")
	// if err != nil {
	// 	return err
	// }
	// url := os.Getenv("PYTHON_API") + "/grade"

	// requestData := &dto.ProductRequest{
	// 	Id:           product.PRODUCT_ID,
	// 	Energy:       product.PRODUCT_ENERGI,
	// 	Protein:      product.PRODUCT_PROTEIN,
	// 	Fat:          product.PRODUCT_LEMAKTOTAL,
	// 	Carbs:        product.PRODUCT_KARBOHIDRAT,
	// 	Sugar:        product.PRODUCT_GULA,
	// 	Salt:         product.PRODUCT_GARAM,
	// 	SaturatedFat: product.PRODUCT_SATURATEDFAT,
	// 	Fiber:        product.PRODUCT_FIBER,
	// }

	// responseData, err := SendRequest[dto.ProductRequest, dto.SNResponse](url, *requestData)
	// if err != nil {
	// 	return err
	// }

	// product.PRODUCT_GRADING = responseData.Grade
	// product.UpdatedAt = time.Now()

	// return service.productRepo.UpdateProduct(product)

	
	
	// Initialize random seed - this is important for truly random results
	rand.Seed(time.Now().UnixNano())
	
	// Create a slice with possible grades
	grades := []string{"A", "B", "C", "D", "E"}
	
	// Generate a random index and select a grade using compatible method
	randomGrade := grades[rand.Intn(len(grades))]
	
	// // Retrieve the product
	// p, err := service.productRepo.GetProductById(id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	// }
	
	// Assign the random grade and update the product
	product.PRODUCT_GRADING = randomGrade
	err = service.productRepo.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// product.PRODUCT_GRADING = responseData.Grade
	product.UpdatedAt = time.Now()

	return service.productRepo.UpdateProduct(product)


}

func (service *productService) DeleteProduct(id string) error {
	product, err := service.productRepo.GetProductById(id)
	if err != nil {
		return err
	}
	service.uploader.DeleteImage(product.PRODUCT_PICTURE,  service.uploader.productPath)
	return service.productRepo.DeleteProduct(id)
}

func (service *productService) AdvertiseProduct(c echo.Context, id string) error {
	product, err := service.productRepo.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	product.PRODUCT_ISSHOW = 1
	product.PRODUCT_EXPSHOW = time.Now().AddDate(0, 1, 0)
	err = service.productRepo.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return nil
}

func (service *productService) UnadvertiseProduct(c echo.Context, id string) error {
	product, err := service.productRepo.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	product.PRODUCT_ISSHOW = 3
	err = service.productRepo.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return nil
}

func (service *productService) GetAllProductAdvertisement(desc, page, pageSize int, search, sort string) ([]models.Product, *dto.Pagination, error) {
	return service.productRepo.GetAllProductAdvertisement(desc, page, pageSize, search, sort)
}

func (service *productService) GetAllProductAdvertisementByStoreId(id string, desc, page, pageSize int, search, sort string) ([]models.Product, *dto.Pagination, error) {
	return service.productRepo.GetAllProductAdvertisementByStoreId(id, desc, page, pageSize, search, sort)
}
