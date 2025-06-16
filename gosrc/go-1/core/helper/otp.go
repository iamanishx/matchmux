package helper

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "time"
)

func GenerateOTP() string {
    return generateCryptoOTP()
}

func generateCryptoOTP() string {
    min := big.NewInt(100000)
    max := big.NewInt(999999)
    rangeVal := new(big.Int).Sub(max, min)
    rangeVal = rangeVal.Add(rangeVal, big.NewInt(1))
    n, err := rand.Int(rand.Reader, rangeVal)
    if err != nil {
        return generateTimeBasedOTP()
    }
    result := new(big.Int).Add(n, min)
    return result.String()
}

func generateTimeBasedOTP() string {
    now := time.Now().UnixNano()
    otp := (now % 900000) + 100000 
    return fmt.Sprintf("%06d", otp)
}

func GenerateAlphanumericOTP(length int) string {
    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)
    for i := range result {
        num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        result[i] = charset[num.Int64()]
    }
    return string(result)
}

func ValidateOTPFormat(otp string) bool {
    if len(otp) != 6 {
        return false
    }
    for _, char := range otp {
        if char < '0' || char > '9' {
            return false
        }
    }
    return true
}