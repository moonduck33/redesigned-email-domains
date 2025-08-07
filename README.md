# Email & Domain Extractor (Golang)

A simple Go tool that scans `.txt` files in a folder, extracts all **email addresses** and their **base domains**, and writes them to separate output files — skipping duplicates across runs.

---

## ✨ Features

- ✅ Extracts emails from `.txt` files
- 🌐 Extracts base domains using `go-domain-util`
- 🧠 Remembers previously seen emails and domains
- 📁 Outputs:
  - `emails_cleaned.txt`
  - `domains_cleaned.txt`
- 🔄 Skips duplicates using:
  - `seen_emails.txt`
  - `seen_domains.txt`

---

## 📦 Requirements

- Go 1.17+
- Module support enabled (`go.mod`)

---

## 📥 Installation

```bash
# Clone the repo
git clone https://github.com/moonduck33/emaildup.git
cd emaildup

# Initialize Go module
go mod init emaildup

# Install domain parser
go get github.com/bobesa/go-domain-util/domainutil
