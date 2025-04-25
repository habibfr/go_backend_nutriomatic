package services

import (
	// "errors"
	"golang-template/dto"
	"golang-template/models"
	"golang-template/repository"
	"os"
	"time"

	"fmt"
	"io"
	// "mime/multipart"
	"path/filepath"

	"net/http"
	// "strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type ScannedNutritionService interface {
	CreateScannedNutrition(c echo.Context, user_id string) error
	GetScannedNutritionById(id string) (*models.ScannedNutrition, error)
	GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, grade, id string) ([]models.ScannedNutrition, *dto.Pagination, error)
}

type scannedNutritionService struct {
	snRepo   repository.ScannedNutritionRepository
	uploader *ClientUploader
}

func NewScannedNutritionService() ScannedNutritionService {
	return &scannedNutritionService{
		snRepo:   repository.NewScannedNutritionRepositoryGORM(),
		uploader: NewClientUploader(),
	}
}

// func (s *scannedNutritionService) CreateScannedNutrition(c echo.Context, user_id string) error {
// 	// Get form values
// 	sn_name := c.FormValue("sn_name")

// 	// Process uploaded image
// 	imagePath, err := s.uploader.ProcessImageScannedNutrition(c)
// 	if err != nil {
// 		return err
// 	}

// 	// Load environment variables
// 	err = godotenv.Load(".env")
// 	if err != nil {
// 		return errors.New("failed to load environment variables: " + err.Error())
// 	}

// 	// Construct full image path and API URL
// 	realImagePath := os.Getenv("IMAGE_PATH") + imagePath
// 	url := os.Getenv("PYTHON_API") + "/ocr"

// 	// Prepare request data
// 	requestData := &dto.SNRequest{
// 		Url: realImagePath,
// 	}

// 	// Send request and handle response
// 	responseData, err := SendRequest[dto.SNRequest, dto.SNResponse](url, *requestData)
// 	if err != nil {
// 		return errors.New("error sending request to OCR API: " + err.Error())
// 	}

// 	// Create ScannedNutrition object
// 	sn := models.ScannedNutrition{
// 		SN_ID:           uuid.New().String(),
// 		SN_PRODUCTNAME:  sn_name,
// 		SN_PRODUCTTYPE:  "", // You might want to populate this field if needed
// 		SN_INFO:         "", // Add more fields if needed
// 		SN_PICTURE:      realImagePath,
// 		SN_ENERGY:       responseData.NutritionFacts.Energy,
// 		SN_PROTEIN:      responseData.NutritionFacts.Protein,
// 		SN_FAT:          responseData.NutritionFacts.Fat,
// 		SN_CARBOHYDRATE: responseData.NutritionFacts.Carbs,
// 		SN_SUGAR:        responseData.NutritionFacts.Sugar,
// 		SN_SALT:         responseData.NutritionFacts.Sodium,
// 		SN_GRADE:        responseData.Grade,
// 		SN_SATURATEDFAT: responseData.NutritionFacts.SaturatedFat,
// 		SN_FIBER:        responseData.NutritionFacts.Fiber,
// 		CreatedAt:       time.Now(),
// 		UpdatedAt:       time.Now(),
// 		USER_ID:         user_id,
// 	}

// 	// Store the scanned nutrition data
// 	err = s.snRepo.CreateScannedNutrition(&sn)
// 	if err != nil {
// 		return errors.New("failed to store scanned nutrition data: " + err.Error())
// 	}

// 	return nil
// }

func (s *scannedNutritionService) CreateScannedNutrition(c echo.Context, user_id string) error {
	// Ambil nama produk dari form
	sn_name := c.FormValue("sn_name")

	// ===== 📷 Proses Upload Gambar Secara Manual =====
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Image file is required")
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

	// ===== 🌐 Build URL publik untuk OCR API =====
	if err := godotenv.Load(".env"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load .env")
	}

	publicImageURL := fmt.Sprintf("%s/uploads/%s", os.Getenv("APP_HOST"), imageName)
	ocrURL := os.Getenv("PYTHON_API") + "/ocr"

	// ===== 🧠 Kirim request ke OCR API =====
	requestData := &dto.SNRequest{
		Url: publicImageURL,
	}

	responseData, err := SendRequest[dto.SNRequest, dto.SNResponse](ocrURL, *requestData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "OCR API error: "+err.Error())
	}

	// ===== 🧾 Simpan data hasil OCR ke database =====
	sn := models.ScannedNutrition{
		SN_ID:           uuid.New().String(),
		SN_PRODUCTNAME:  sn_name,
		SN_PRODUCTTYPE:  "",
		SN_INFO:         "",
		SN_PICTURE:      publicImageURL,
		SN_ENERGY:       responseData.NutritionFacts.Energy,
		SN_PROTEIN:      responseData.NutritionFacts.Protein,
		SN_FAT:          responseData.NutritionFacts.Fat,
		SN_CARBOHYDRATE: responseData.NutritionFacts.Carbs,
		SN_SUGAR:        responseData.NutritionFacts.Sugar,
		SN_SALT:         responseData.NutritionFacts.Sodium,
		SN_GRADE:        responseData.Grade,
		SN_SATURATEDFAT: responseData.NutritionFacts.SaturatedFat,
		SN_FIBER:        responseData.NutritionFacts.Fiber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		USER_ID:         user_id,
	}

	if err := s.snRepo.CreateScannedNutrition(&sn); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to store scanned nutrition data: "+err.Error())
	}

	return nil

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"status":  "success",
	// 	"message": "Scanned nutrition created successfully",
	// 	"data":    sn,
	// })
}		


func (s *scannedNutritionService) GetScannedNutritionById(id string) (*models.ScannedNutrition, error) {
	return s.snRepo.GetScannedNutritionById(id)
}

func (s *scannedNutritionService) GetScannedNutritionByUserId(desc, page, pageSize int, search, sort, grade, id string) ([]models.ScannedNutrition, *dto.Pagination, error) {
	return s.snRepo.GetScannedNutritionByUserId(desc, page, pageSize, search, sort, grade, id)
}
