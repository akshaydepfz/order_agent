package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"order_agent/internal/models"
	"order_agent/internal/repository"
	"order_agent/pkg/s3"
	"order_agent/pkg/utils"
)

type ShopHandler struct {
	shopRepo *repository.ShopRepo
	s3Client *s3.Client
}

func NewShopHandler(shopRepo *repository.ShopRepo, s3Client *s3.Client) *ShopHandler {
	return &ShopHandler{shopRepo: shopRepo, s3Client: s3Client}
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseBool(s string) bool {
	return strings.ToLower(s) == "true" || s == "1"
}

func parseDate(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}
	return t
}

func (h *ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid multipart form")
		return
	}

	name := r.FormValue("name")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	if name == "" || phone == "" || address == "" {
		utils.Error(w, http.StatusBadRequest, "Name, phone, and address are required")
		return
	}

	var logoURL string
	if h.s3Client != nil {
		file, fileHeader, err := r.FormFile("brand_image")
		if err != nil {
			utils.Error(w, http.StatusBadRequest, "Brand image is required")
			return
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "Failed to read image")
			return
		}

		contentType := http.DetectContentType(fileBytes)
		if !strings.HasPrefix(contentType, "image/") {
			utils.Error(w, http.StatusBadRequest, "Only image files are allowed")
			return
		}

		key := h.s3Client.GenerateShopLogoKey(fileHeader.Filename)
		logoURL, err = h.s3Client.UploadImage(r.Context(), key, fileBytes, contentType)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "Image upload failed")
			return
		}
	} else {
		utils.Error(w, http.StatusServiceUnavailable, "Image upload service unavailable")
		return
	}

	shop := &models.Shop{
		ID:               uuid.New().String(),
		Name:             name,
		Description:      r.FormValue("description"),
		LogoURL:          logoURL,
		OwnerName:        r.FormValue("owner_name"),
		Phone:            phone,
		Email:            r.FormValue("email"),
		WhatsAppNumber:   r.FormValue("whatsapp_number"),
		PhoneNumberID:    r.FormValue("phone_number_id"),
		AccessToken:      r.FormValue("access_token"),
		VerifyToken:      r.FormValue("verify_token"),
		Address:          address,
		City:             r.FormValue("city"),
		State:            r.FormValue("state"),
		Pincode:          r.FormValue("pincode"),
		Latitude:         parseFloat(r.FormValue("latitude")),
		Longitude:        parseFloat(r.FormValue("longitude")),
		OpeningTime:      r.FormValue("opening_time"),
		ClosingTime:     r.FormValue("closing_time"),
		DeliveryRadiusKM: parseInt(r.FormValue("delivery_radius_km")),
		DeliveryFee:      parseFloat(r.FormValue("delivery_fee")),
		MinimumOrder:     parseFloat(r.FormValue("minimum_order")),
		DeliveryTimeMin:  parseInt(r.FormValue("delivery_time_min")),
		CODEnabled:       parseBool(r.FormValue("cod_enabled")),
		UPIEnabled:       parseBool(r.FormValue("upi_enabled")),
		AIContext:        r.FormValue("ai_context"),
		SelectedPlan:     r.FormValue("selected_plan"),
		PlanStartDate:    parseDate(r.FormValue("plan_start_date")),
		PlanExpireDate:   parseDate(r.FormValue("plan_expire_date")),
		Status:           r.FormValue("status"),
	}
	if shop.Status == "" {
		shop.Status = "active"
	}

	if err := h.shopRepo.Create(r.Context(), shop); err != nil {
		if errors.Is(err, repository.ErrDBNotConfigured) {
			utils.Error(w, http.StatusServiceUnavailable, "Database not configured")
			return
		}
		log.Printf("Create shop error: %v", err)
		utils.Error(w, http.StatusInternalServerError, "Failed to create shop")
		return
	}
	utils.Success(w, shop)
}

func (h *ShopHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	shops, err := h.shopRepo.List(r.Context())
	if err != nil {
		if errors.Is(err, repository.ErrDBNotConfigured) {
			utils.Error(w, http.StatusServiceUnavailable, "Database not configured")
			return
		}
		log.Printf("List shops error: %v", err)
		utils.Error(w, http.StatusInternalServerError, "Failed to get shops")
		return
	}
	if shops == nil {
		shops = []*models.Shop{}
	}
	utils.Success(w, shops)
}

func (h *ShopHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := r.PathValue("id")
	if id == "" {
		utils.Error(w, http.StatusBadRequest, "Shop ID is required")
		return
	}

	shop, err := h.shopRepo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrDBNotConfigured) {
			utils.Error(w, http.StatusServiceUnavailable, "Database not configured")
			return
		}
		log.Printf("GetByID shop error: %v", err)
		utils.Error(w, http.StatusInternalServerError, "Failed to get shop")
		return
	}
	if shop == nil {
		utils.Error(w, http.StatusNotFound, "Shop not found")
		return
	}
	utils.Success(w, shop)
}
