package main

import (
	"bytes"
	"crypto/sha256"
	"io"
	"log"
	"os"

	"github.com/edgelesssys/ego/eclient"
)

func main() {
	file, err := os.OpenFile("report.bin", os.O_RDONLY, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reportBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	data := []byte("some kind of data signed by the report")

	report, err := eclient.VerifyRemoteReport(reportBytes)
	if err != nil {
		log.Fatalln(err)
	}

	// Report data can contain up to 64 bytes but the data signed is usually longer than that,
	// so it's usually hashed.
	// In case of this example, the report was created for "some kind of data signed by the report" string and hashed with SHA256, which is 32 bytes long
	dataSum := sha256.Sum256(data)
	if !bytes.Equal(dataSum[:], report.Data[:32]) {
		log.Fatalln("report data does not match the certificate's hash")
	}

	log.Println("Report signature verified")

	log.Println("Data:", report.Data)
	log.Println("SecurityVersion:", report.SecurityVersion)
	log.Println("Debug:", report.Debug)
	log.Println("UniqueID:", report.UniqueID)
	log.Println("SignerID:", report.SignerID)
	log.Println("ProductID:", report.ProductID)
	log.Println("TCBStatus:", report.TCBStatus)

	// While VerifyRemoteReport verifies the report's integrity and signature, the third party must additionally verify the content of the remote report
	// if report.SecurityVersion < 2 {
	// 	log.Fatalln("invalid security version")
	// }
	// var expectedProductId uint16 = 1234
	// // verifying code version
	// if binary.LittleEndian.Uint16(report.ProductID) != expectedProductId {
	// 	log.Fatalln("invalid product")
	// }
	// // verifying signer
	// if !bytes.Equal(report.SignerID, []byte("enclave's public key")) {
	// 	log.Fatalln("invalid signer")
	// }
	// if report.Debug {
	// 	log.Fatalln("debug enclave not allowed")
	// }
}
