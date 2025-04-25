package services

import (
	"golang-template/dto"
	"golang-template/middleware"
	"golang-template/models"
	"golang-template/repository"
	"os"
	"time"

	"fmt"
	"io"
	// "mime/multipart"
	"path/filepath"
	"net/http"
	
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type TransactionService interface {
	CreateTransaction(c echo.Context) error
	GetTransactionById(id string) (*models.Transaction, error)
	GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	GetAllTransaction(desc, page, pageSize int, search, sort, status string) (*[]models.Transaction, *dto.Pagination, error)
	GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error)
	UpdateStatusTransaction(status string, c echo.Context, id string) error
	DeleteTransaction(c echo.Context) error
	UploadProofPayment(c echo.Context) error
	FindAllNewTransactions(desc, page, pageSize int, search, sort, status, id string) (*[]models.Transaction, *dto.Pagination, error)
}

type transactionService struct {
	tscRepo        repository.TransactionRepository
	tokenRepo      repository.TokenRepository
	storeRepo      repository.StoreRepository
	productRepo    repository.ProductRepository
	uploader       *ClientUploader
	paymentRepo    repository.PaymentRepository
	productService ProductService
}

func NewTransactionService() TransactionService {
	return &transactionService{
		tscRepo:        repository.NewTransactionRepositoryGORM(),
		tokenRepo:      repository.NewTokenRepositoryGORM(),
		storeRepo:      repository.NewStoreRepositoryGORM(),
		productRepo:    repository.NewProductRepositoryGORM(),
		uploader:       NewClientUploader(),
		paymentRepo:    repository.NewPaymentRepositoryGORM(),
		productService: NewProductService(),
	}
}

func (s *transactionService) CreateTransaction(c echo.Context) error {
	paymentMethod := &dto.PaymentTransaction{}
	err := c.Bind(paymentMethod)
	if err != nil {
		return err
	}

	id := c.Param("product_id")
	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return err
	}

	// availStore, err := s.productRepo.GetStoreByProductId(id)
	// if err != nil {
	// 	return err
	// }

	// if store.STORE_ID != availStore.STORE_ID {
	// 	return err
	// }

	payment_id, err := s.paymentRepo.GetPaymentIdByMethod(paymentMethod.PaymentMethod)
	if err != nil {
		return err
	}

	product, err := s.productRepo.GetProductById(id)
	if err != nil {
		return err
	}

	product.PRODUCT_ISSHOW = 2
	err = s.productRepo.UpdateProduct(product)
	if err != nil {
		return err
	}

	newTsc := &models.Transaction{
		TSC_ID:     uuid.New().String(),
		TSC_PRICE:  5000.00,
		TSC_START:  time.Now(),
		TSC_END:    time.Now().AddDate(0, 0, 1),
		TSC_STATUS: "pending",
		PAYMENT_ID: payment_id,
		STORE_ID:   store.STORE_ID,
		PRODUCT_ID: id,
	}

	return s.tscRepo.CreateTransaction(newTsc)
}

func (s *transactionService) GetTransactionById(id string) (*models.Transaction, error) {
	return s.tscRepo.GetTransactionById(id)
}

func (s *transactionService) GetTransactionByStoreId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.GetTransactionByStoreId(id, desc, page, pageSize, search, sort)
}

func (s *transactionService) GetAllTransaction(desc, page, pageSize int, search, sort, status string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.GetAllTransaction(desc, page, pageSize, search, sort, status)
}

func (s *transactionService) GetTransactionByUserId(id string, desc, page, pageSize int, search, sort string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.GetTransactionByUserId(id, desc, page, pageSize, search, sort)
}

func (s *transactionService) UpdateStatusTransaction(status string, c echo.Context, id string) error {
	_, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	// store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	// if err != nil {
	// 	return err
	// }

	tsc, err := s.tscRepo.GetTransactionById(id)
	if err != nil {
		return err
	}

	// if tsc.STORE_ID != store.STORE_ID {
	// 	return err
	// }

	if status == "accepted" {
		err := s.productService.AdvertiseProduct(c, tsc.PRODUCT_ID)
		if err != nil {
			return err
		}
	} else {
		err := s.productService.UnadvertiseProduct(c, tsc.PRODUCT_ID)
		if err != nil {
			return err
		}

	}

	return s.tscRepo.UpdateStatusTransaction(tsc.TSC_ID, status)
}

func (s *transactionService) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return err
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return err
	}

	tsc, err := s.tscRepo.GetTransactionById(id)
	if err != nil {
		return err
	}

	if tsc.STORE_ID != store.STORE_ID {
		return err
	}

	return s.tscRepo.DeleteTransaction(id)
}

// func (s *transactionService) UploadProofPayment(c echo.Context) error {
// 	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
// 	if err != nil {
// 		return err
// 	}

// 	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
// 	if err != nil {
// 		return err
// 	}

// 	tsc, err := s.tscRepo.FindAllNewTransactionsWithoutPagination(store.STORE_ID)
// 	if err != nil {
// 		return err
// 	}

// 	imagePath, err := s.uploader.ProcessImageProof(c)
// 	if err != nil {
// 		return err
// 	}
// 	err = godotenv.Load(".env")
// 	if err != nil {
// 		return err
// 	}
// 	realImagePath := os.Getenv("IMAGE_PATH") + imagePath

// 	// loop each transactions
// 	for _, transaction := range *tsc {
// 		transaction.TSC_BUKTI = realImagePath
// 		transaction.UpdatedAt = time.Now()
// 		if err := s.tscRepo.UpdateTransaction(&transaction); err != nil {
// 			return err
// 		}

// 	}
// 	return nil
// }

func (s *transactionService) UploadProofPayment(c echo.Context) error {
	userToken, err := s.tokenRepo.UserToken(middleware.GetToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store, err := s.storeRepo.GetStoreByUserId(userToken.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Store not found")
	}

	tsc, err := s.tscRepo.FindAllNewTransactionsWithoutPagination(store.STORE_ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve transactions")
	}

	// ===== 🖼️ Upload Bukti Pembayaran =====
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Proof of payment image is required")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open uploaded file")
	}
	defer src.Close()

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

	// ===== 🌐 Simpan sebagai URL publik =====
	if err := godotenv.Load(".env"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load .env")
	}
	publicImageURL := fmt.Sprintf("%s/uploads/%s", os.Getenv("APP_HOST"), imageName)

	// ===== 🔁 Update semua transaksi =====
	for _, transaction := range *tsc {
		transaction.TSC_BUKTI = publicImageURL
		transaction.UpdatedAt = time.Now()

		if err := s.tscRepo.UpdateTransaction(&transaction); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update transaction")
		}
	}

	return nil


	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"status":  "success",
	// 	"message": "Proof of payment uploaded",
	// 	"url":     publicImageURL,
	// })


	// 	// loop each transactions
// 	for _, transaction := range *tsc {
// 		transaction.TSC_BUKTI = realImagePath
// 		transaction.UpdatedAt = time.Now()
// 		if err := s.tscRepo.UpdateTransaction(&transaction); err != nil {
// 			return err
// 		}

// 	}
// 	return nil
}


func (s *transactionService) FindAllNewTransactions(desc, page, pageSize int, search, sort, status, id string) (*[]models.Transaction, *dto.Pagination, error) {
	return s.tscRepo.FindAllNewTransactions(desc, page, pageSize, search, sort, status, id)
}
