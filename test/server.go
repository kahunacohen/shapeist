package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Patient represents a patient entity
type Patient struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

// PatientStore is an in-memory store for patients
type PatientStore struct {
	mu       sync.RWMutex
	patients map[int]Patient
	nextID   int
}

// NewPatientStore creates a new patient store with random data
func NewPatientStore() *PatientStore {
	store := &PatientStore{
		patients: make(map[int]Patient),
		nextID:   1,
	}
	
	// Seed with some random patients
	for i := 0; i < 5; i++ {
		store.addPatient(generateRandomPatient())
	}
	
	return store
}

// addPatient adds a patient to the store (internal use)
func (s *PatientStore) addPatient(p Patient) Patient {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	p.ID = s.nextID
	s.patients[p.ID] = p
	s.nextID++
	
	return p
}

// Create adds a new patient
func (s *PatientStore) Create(p Patient) Patient {
	return s.addPatient(p)
}

// Get retrieves a patient by ID
func (s *PatientStore) Get(id int) (Patient, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	p, exists := s.patients[id]
	return p, exists
}

// List returns all patients
func (s *PatientStore) List() []Patient {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	patients := make([]Patient, 0, len(s.patients))
	for _, p := range s.patients {
		patients = append(patients, p)
	}
	
	return patients
}

// Update modifies an existing patient
func (s *PatientStore) Update(id int, p Patient) (Patient, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.patients[id]; !exists {
		return Patient{}, false
	}
	
	p.ID = id
	s.patients[id] = p
	
	return p, true
}

// Delete removes a patient
func (s *PatientStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.patients[id]; !exists {
		return false
	}
	
	delete(s.patients, id)
	return true
}

// Random data generators
var (
	firstNames = []string{"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda", "William", "Elizabeth"}
	lastNames  = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}
	genders    = []string{"Male", "Female", "Other"}
)

func generateRandomPatient() Patient {
	rand.Seed(time.Now().UnixNano())
	
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	
	// Generate birthdate between 1940 and 2020
	minDate := time.Date(1940, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)
	delta := maxDate.Unix() - minDate.Unix()
	sec := rand.Int63n(delta) + minDate.Unix()
	birthDate := time.Unix(sec, 0)
	
	gender := genders[rand.Intn(len(genders))]
	email := fmt.Sprintf("%s.%s@example.com", strings.ToLower(firstName), strings.ToLower(lastName))
	phone := fmt.Sprintf("%03d-%03d-%04d", rand.Intn(1000), rand.Intn(1000), rand.Intn(10000))
	
	return Patient{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Gender:    gender,
		Email:     email,
		Phone:     phone,
	}
}

// HealthcareServer is the HTTP handler for the healthcare API
type HealthcareServer struct {
	store *PatientStore
}

// NewHealthcareServer creates a new healthcare server
func NewHealthcareServer() *HealthcareServer {
	return &HealthcareServer{
		store: NewPatientStore(),
	}
}

// ServeHTTP implements the http.Handler interface
func (s *HealthcareServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Route handling
	path := r.URL.Path
	
	if path == "/patients" {
		switch r.Method {
		case http.MethodGet:
			s.handleListPatients(w, r)
		case http.MethodPost:
			s.handleCreatePatient(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	
	// Handle /patients/{id}
	if strings.HasPrefix(path, "/patients/") {
		idStr := strings.TrimPrefix(path, "/patients/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid patient ID", http.StatusBadRequest)
			return
		}
		
		switch r.Method {
		case http.MethodGet:
			s.handleGetPatient(w, r, id)
		case http.MethodPut:
			s.handleUpdatePatient(w, r, id)
		case http.MethodDelete:
			s.handleDeletePatient(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	
	http.NotFound(w, r)
}

func (s *HealthcareServer) handleListPatients(w http.ResponseWriter, r *http.Request) {
	patients := s.store.List()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

func (s *HealthcareServer) handleGetPatient(w http.ResponseWriter, r *http.Request, id int) {
	patient, exists := s.store.Get(id)
	if !exists {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

func (s *HealthcareServer) handleCreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	created := s.store.Create(patient)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (s *HealthcareServer) handleUpdatePatient(w http.ResponseWriter, r *http.Request, id int) {
	var patient Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	updated, exists := s.store.Update(id, patient)
	if !exists {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (s *HealthcareServer) handleDeletePatient(w http.ResponseWriter, r *http.Request, id int) {
	if !s.store.Delete(id) {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	server := NewHealthcareServer()
	
	fmt.Println("Healthcare test server starting on :8080")
	fmt.Println("Endpoints:")
	fmt.Println("  GET    /patients      - List all patients")
	fmt.Println("  POST   /patients      - Create a new patient")
	fmt.Println("  GET    /patients/{id} - Get a patient by ID")
	fmt.Println("  PUT    /patients/{id} - Update a patient")
	fmt.Println("  DELETE /patients/{id} - Delete a patient")
	
	if err := http.ListenAndServe(":8080", server); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
