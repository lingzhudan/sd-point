package _test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

func TestRSA(t *testing.T) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PrivateKey",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PublicKey",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
	fmt.Printf("privateKey: %+v\n", privateKey)
	fmt.Printf("privateKey: %+v\n", publicKey)
}

func TestRsaEncrypt(t *testing.T) {
	p := `-----BEGIN PublicKey-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtbdqlS6wjF8tddJSUF5f
JkJjs1R+UQa/FheUgDCNa8PzKDNVCHHD0Q3XMGfX1RD8FfbzlvRvlETj5Fc5qDS/
uoYPOAtTBqZJZ68R+927hAu59ZEhJoo6detT2HUlfVX3YC+LGIoLTWsz4Yv4WXnt
ycW4tL9J3IOq1JyTjRNI70utnLbBhlBKNs4sYk6xpUYmHt62eMbAnFAe5LO9pBUX
ojesG9oVigJL6mdgRNwJwpS3ddfgvsAggOieOAdMqjQ53OCW4uu6rIn4m9C3UZQi
FS8UGzj0fND4hkDjsbNsQQsNVD7YfX0JsaFw5sIzY/ZSQZjo1cytKxTPDQuGTbs1
XQIDAQAB
-----END PublicKey-----
`
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(p))
	if block == nil {
		panic("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte("Sdpoint958836!"))
	if err != nil {
		panic(err)
	}
	s := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("密文srs: %+v\n", s)
	var c []byte
	if c, err = base64.StdEncoding.DecodeString("B8dNoGw+kiy5WUsM79BlPs2j7Ag3ytoh6uqwxAkOILUc5TDPuxCw45nq6x8EFIEEJdq9/Tt0ZyJqFEeM+BM5okmnVBLUFlL8cxT7qWBbW54mbZKKkYFojIyQVt4wTNbD+J942xreVFNZ/ZLVWQYN/+82lsjZyXbT2X3Kvefw8gQhNxIWcsU1ksripewa0H18uxcqHXYqylpqrj7A/jp5MaalV5+WrKVXfzJNtSlEZ8MpYUQboG6DUPTcXyAoa/1e3tLpE4VFrjiSFyB5Nr3oTiKMXTcYJFJc/Ay/Z9KLT9V03AWAYWxJmU0MtIjJllyHfQX2my8xSKtVN7/nf2kkIQ=="); err != nil {
		t.Errorf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("D/z7MAEyWJZPu5+QmS5xWJ7pBjElJTs2/N835kuH6hYnNEH/KHht6qP9omu9rEPubXRY68I4+33Kvnnz/JTf2ieSabPDeevmjcGIVhGTm/IMMzrjIVYbrocsrblofG1J8I1jupW3f/o0UdpjF/T2nNeC9S930fBcNt+nSzmhH2OAVSCV6d9Lo6dysKxlgUX6swHNsI+23hdihbxDJL+VB1lHox9NeeUCHNmskrGhxVk5j3aDvA2YPXCkMRaeP1QJ+Od1cKLzCZSnl0puuxqFRC84jHZgiA18bPkpKPBe/NcErS2TF207sTeh6ge/iC3YsxVG0xexTBuIZutCDCfWJg==\n")

	pk := `-----BEGIN PrivateKey-----
MIIEogIBAAKCAQEAtbdqlS6wjF8tddJSUF5fJkJjs1R+UQa/FheUgDCNa8PzKDNV
CHHD0Q3XMGfX1RD8FfbzlvRvlETj5Fc5qDS/uoYPOAtTBqZJZ68R+927hAu59ZEh
Joo6detT2HUlfVX3YC+LGIoLTWsz4Yv4WXntycW4tL9J3IOq1JyTjRNI70utnLbB
hlBKNs4sYk6xpUYmHt62eMbAnFAe5LO9pBUXojesG9oVigJL6mdgRNwJwpS3ddfg
vsAggOieOAdMqjQ53OCW4uu6rIn4m9C3UZQiFS8UGzj0fND4hkDjsbNsQQsNVD7Y
fX0JsaFw5sIzY/ZSQZjo1cytKxTPDQuGTbs1XQIDAQABAoIBACZUiibtsk/pw60W
sEZDoc5wMpehOwPcaHJAhxKDK1GQA7p3GXiuhp+SS9HqFZzb+FCpxrgQ13hvD6Ma
ww64EblGje9EdD7y1IkKZMC8BNHVp2QlONjoT2yQNx9xlnbZq1Sesrh6ZefutQLe
g4RcM8xb6Jo3v1zEaURvwq//YWMYOCEnFbigelTz+EeHrkBy2dk9LyhiaqFl8RIq
0EEvqdRFmNHPifABrYh77SsID23BJc0enhO70kfYVOmuC8NsOnxhnI6ftNP9zK4s
dFIuwxeXUSEPlK4UXXwIaybFsVLCUFUi4ECVsEAsqyLwQPk42iqJMp+kvNt+X/Lw
feKVwAECgYEAxEFq378g6zFskOxFX6MeZLzeySiF8qx1PTqaVaEXWe+TGS8QfXYE
FY4jWuhGRjXnqCjrSgnC/l3k3nEWKRs/+9/kOI01IKG81xTM8ctHBZilbXoA0yhM
3EDeZN/P7ZrNHRlpMfvHXPmWs+etIR6UUcUoOEOrPhxHZWJ5PO1BK1UCgYEA7Qjw
yqwbduuijCMS/bWwv/lG4fN7timfR7yLu1hTdoOK2ZpDJ35x0D8Nsa1giHqXXUIs
+VSCWvUwQOzQJjKobdQrUGePbTukOsmDIf6KWNAyYNhRNYWIoxQXGmPydLcQQGNe
JKoxNvsP3LUX9j5i0/qZoZefuIj2+oflh9ahsekCgYBxYjYDtKGpH/HQPKDL4yrB
Gpk/8IiKrnZBfKUDycD69b2vdBB72336khb9A2ruT2BaN+HP8ZnVaZu7o+wgrZym
wTBSN8q04px+SocpPr72jb0tZaoqVIMrUNLrxp6TcLcGE9NBzGve3Ffj8nqEwuz/
3P8imn3JsX7SJTVULOJUzQKBgGPWpDV5kIficDC+hT0/zhbOFEEPoUf0+BXrIJ/k
cCM2/MDy7N4xsBaauFXiRQq3OBIb0X2jTtSo+Y4fMP0l0TORbj0LoanMuTvZrNYg
tHz8FHweIPZ+LDMkPybqejEW+k1kjT68QMZ7sQ6xaEpbtJy8rQEXMDs11XFo1BEH
IkwhAoGANFbtXRM0ySApQTTaQO9BVipTDjaJaewRvkbG6JJdk9UNTwP3mThaIg5q
6nAmKZF6PuH8d6E+l9yfmEMJ6FZzAvh9MFmJNuZ7wcv9bujWW3nyb9GxojB8uBxQ
yYYTyZtcjVyKL9QJLL0axqsX1Pxy1g/QVxjnAET3KpVkbXnJ8ss=
-----END PrivateKey-----
`

	//获取私钥
	block, _ = pem.Decode([]byte(pk))
	if block == nil {
		t.Errorf("private key error!\n")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, c)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	t.Errorf("明文: %+v\n", string(data))
}
