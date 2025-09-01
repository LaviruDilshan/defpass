# DefPass 🔑

**Default Password Finder CLI** – quickly fetch default passwords for network vendors.

---

## 🚀 Features

* Fetches a **vendor list** from your target site.
* Retrieves **default passwords** for each vendor.
* User-friendly CLI interface with **colorful output**.
* Randomizes **User-Agent headers** for requests.
* Fully **installable via `go install`**.

---

## 📦 Installation

Make sure you have **Go 1.18+** installed. Then run:

```bash
go install github.com/LaviruDeveloper/defpass/cmd/defpass@latest
```

> Make sure `$GOPATH/bin` or `$HOME/go/bin` is in your `PATH`.

Then you can run anywhere:

```bash
defpass
```

---

## 🖥 Usage

Run the tool:

```bash
defpass
```

You’ll see a banner and the list of available vendors:

```
1. Cisco
2. 3COM
3. MikroTik
...
```

* Enter a **vendor number** to fetch default passwords.
* Press **Enter** to go back to vendor list.
* Type `q` at any input prompt to **quit**.

---

## 🎨 Colors

* **Banner & tool name:** Cyan
* **Author:** Green
* **Social links:** Yellow

---

## 🔧 Project Structure

```
defpass/
├── cmd/defpass/main.go      # Entry point
├── internal/scraper/        # Web scraping logic
│   ├── vendors.go
│   ├── passwords.go
│   └── utils.go
├── internal/ui/banner.go    # Banner & colors
├── user_agents.txt          # Custom User-Agents
├── go.mod
```

---

## 👤 Author

**Laviru Dilshan**

* LinkedIn: [Laviru Dilshan](https://linkedin.com/in/lavirudilshan)
* X: [@LaviruDilshan](https://x.com/LaviruDilshan)
* Instagram: [lavirudilshn](https://instagram.com/lavirudilshn)

---

## ⚠️ Notes

* Designed for **educational & personal use**.
* Make sure the target site allows scraping.

---

## 💡 Example

```bash
$ defpass

Available Vendors:
1. Cisco
2. 3COM
3. MikroTik

Enter vendor number (or 'q' to quit): 2

Fetching default passwords for vendor: 3COM...

---- Entry 1 ----
Method: SNMP
Password: comcomcom
Doc: 
-----------------

Press Enter to return to vendor list, or 'q' to quit:
```

---
