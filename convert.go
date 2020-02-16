package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jordan-wright/email"
)

var bookExt = []string{".pdf", ".epub"}

// TODO: Convert file into mobi
func convertFileToMobi(path, outputDir string) error {
	log.Printf("Converting: %s in %s", path, outputDir)

	var fileName string

	for _, ext := range bookExt {
		if strings.HasSuffix(path, ext) {
			fileName = strings.TrimSuffix(path, ext)
			break
		}
	}

	cFileName := fmt.Sprintf("%s.mobi", filepath.Base(fileName))
	cPath := filepath.Join(outputDir, cFileName)

	// Execute ebook-convert
	binPath := "ebook-convert"
	if cfg.app.CalibrePath != "" {
		binPath = fmt.Sprintf("%s/ebook-convert", cfg.app.CalibrePath)
	}

	cmd := exec.Command(binPath, path, cPath)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Send email for Kindle.
	e := email.NewEmail()
	e.Subject = fmt.Sprintf("Curator: %s", cFileName)
	e.From = cfg.mailer.FromAddr
	e.To = cfg.app.Emails
	e.AttachFile(cPath)

	srv := fmt.Sprintf("%s:%d", cfg.mailer.ServerAddr, cfg.mailer.ServerPort)

	err = e.Send(srv, smtp.PlainAuth("", cfg.mailer.SMTPUsername, cfg.mailer.SMTPPassword, cfg.mailer.ServerAddr))
	if err != nil {
		return fmt.Errorf("error while sending email: %v", err)
	}

	log.Printf("Sent email for %s", cPath)

	return nil
}
