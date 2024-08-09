package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateSUCI Generate Subscription Concealed Identifier(SUCI). (use null scheme)
func GenerateSUCI(SUPI string) string {
	// 对SUPI使用空保护策略
	// SUCI=SUPI类型取值(0表示imsi)||归属网络标识符(mcc+mnc)||路由标识符||SUPI保护算法ID(0表示null scheme)||归属网络公钥||msin
	mcc := SUPI[:3]
	mnc := SUPI[3:5]
	msin := SUPI[5:]
	SUCI := "0" + mcc + mnc + "678" + "0" + "0" + msin
	return SUCI
}

func ResolveAUTN(AUTN string) (string, string, string) {
	sqnAK := AUTN[:12]
	amf := AUTN[12:16]
	mac := AUTN[16:]
	return sqnAK, amf, mac
}

func CheckMac(xMacA, MacA string) bool {
	return xMacA == MacA
}

func GenerateResStar(ck, ik, P0, L0, rand, res string) string {
	key := []byte(ck + ik)
	P1 := rand
	L1 := fmt.Sprintf("%x", len(P1))
	P2 := res
	L2 := fmt.Sprintf("%x", len(P2))
	s := []byte("6B" + P0 + L0 + P1 + L1 + P2 + L2)
	h := hmac.New(sha256.New, key)
	h.Write(s)
	resStar := hex.EncodeToString(h.Sum(nil))[32:]
	return resStar
}
