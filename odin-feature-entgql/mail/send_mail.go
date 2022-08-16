package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"
)

func SendArtistMail(name string, externalURL string, phoneNumber string, discord string, recommender string, createdAt time.Time) {
	auth := smtp.PlainAuth("", os.Getenv("FKT_SMTP_USERNAME"), os.Getenv("FKT_SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	from := os.Getenv("FKT_SMTP_USERNAME")
	to := []string{os.Getenv("FKT_SMTP_USERNAME")}

	headerSubject := "Subject: [FKT] 아티스트 신청\r\n"
	headerBlank := "\r\n"
	body := fmt.Sprintf("이름 - %s\nSNS/포트폴리오 링크 - %s\n연락처 - %s\n디스코드 닉네임 - %s\n디스코드 추천인 - %s\n\n신청 일시 - %s\r\n", name, externalURL, phoneNumber, discord, recommender, createdAt)
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Printf("failed sending artist information mail: %v", err)
	}
}

func SendPartnershipMail(name string, company string, email string, content string, createdAt time.Time) {
	auth := smtp.PlainAuth("", os.Getenv("KL_SMTP_USERNAME"), os.Getenv("KL_SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	from := os.Getenv("KL_SMTP_USERNAME")
	to := []string{os.Getenv("KL_SMTP_USERNAME")}

	headerSubject := "Subject: [KL] 파트너쉽 문의\r\n"
	headerBlank := "\r\n"
	body := fmt.Sprintf("이름 - %s\n회사/소속 - %s\n이메일 - %s\n문의 내용 - %s\n\n문의 일시 - %s\n", name, company, email, content, createdAt)
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Printf("failed sending partnership information mail: %v", err)
	}
}
