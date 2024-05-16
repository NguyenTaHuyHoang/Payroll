package func_support

import (
	"crypto/rand"
	"encoding/hex"
)

func generateRandom160BitAddress() (string, error) {
	// Tạo một mảng byte có độ dài 20 byte (160 bit)
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// Chuyển đổi mảng byte thành chuỗi hex
	address := hex.EncodeToString(bytes)
	return address, nil
}

// func main() {
// 	address, err := generateRandom160BitAddress()
// 	if err != nil {
// 		log.Fatalf("Error generating random address: %v", err)
// 	}
// 	fmt.Printf("Random 160-bit address: %s\n", address)
// }
