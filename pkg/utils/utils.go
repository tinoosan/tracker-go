package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	u := uuid.New()
	return u
}

func IsAuthorizedAndValid(w http.ResponseWriter, r *http.Request) (uuid.UUID, bool) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return userID, false
	}
	return userID, true
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ShowMenu() {
	menuOptions := []string{"Accounts", "Ledger", "Exit"}
	fmt.Println("Choose an option:")
	for i, option := range menuOptions {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func GetInputString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return input, nil
}

func GetInputFloat(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.Trim(input, "\n")
	parsedInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return parsedInput, nil
}
