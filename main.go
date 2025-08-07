package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bobesa/go-domain-util/domainutil"
)

var (
	inputDir         = "raw" // <-- change this
	emailsOutput     = "emails_cleaned.txt"
	domainsOutput    = "domains_cleaned.txt"
	seenEmailsFile   = "seen_emails.txt"
	seenDomainsFile  = "seen_domains.txt"
	emailRegex       = regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@([a-zA-Z0-9.\-]+\.[a-zA-Z]{2,})`)
)

func loadSeen(path string) map[string]struct{} {
	seen := make(map[string]struct{})
	file, err := os.Open(path)
	if err != nil {
		return seen
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seen[strings.ToLower(scanner.Text())] = struct{}{}
	}
	return seen
}

func appendLine(path, line string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to", path, ":", err)
		return
	}
	defer f.Close()
	f.WriteString(line + "\n")
}

func main() {
	// Check input folder
	files, err := os.ReadDir(inputDir)
	if err != nil {
		fmt.Println("[!] Input directory not found:", inputDir)
		return
	}

	// Load seen data
	seenEmails := loadSeen(seenEmailsFile)
	seenDomains := loadSeen(seenDomainsFile)

	emailsAdded := 0
	domainsAdded := 0

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		fullPath := filepath.Join(inputDir, file.Name())
		f, err := os.Open(fullPath)
		if err != nil {
			fmt.Println("[!] Error reading:", fullPath)
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			matches := emailRegex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				if len(match) < 2 {
					continue
				}
				email := strings.ToLower(match[0])
				domain := strings.ToLower(match[1])

				// Add email
				if _, seen := seenEmails[email]; !seen {
					appendLine(emailsOutput, email)
					appendLine(seenEmailsFile, email)
					seenEmails[email] = struct{}{}
					emailsAdded++
				}

				// Extract and add base domain
				base := domainutil.Domain(domain)
				if base != "" {
					if _, seen := seenDomains[base]; !seen {
						appendLine(domainsOutput, base)
						appendLine(seenDomainsFile, base)
						seenDomains[base] = struct{}{}
						domainsAdded++
					}
				}
			}
		}
	}

	fmt.Println("\n[✓] Emails added: ", emailsAdded)
	fmt.Println("[✓] Domains added:", domainsAdded)
	fmt.Println("[✓] Done.")
}
