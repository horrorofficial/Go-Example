---

# 🧩 AuthSecure Go SDK

A modern **Go SDK** for integrating your app with **[AuthSecure](https://authsecure.shop)** —
secure authentication made simple for loaders, applications, and tools.

---

## 🚀 Features

✅ Simple plug-and-play SDK
✅ Secure HTTPS API communication
✅ HWID binding (Windows SID)
✅ Works for: Login, Register, and License Login
✅ 100% Open Source & Easy Integration

---

## ⚙️ Setup Guide (Windows)

### 🪟 Step 1: Install Go

Download and install Go (latest version):
👉 [https://go.dev/dl/](https://go.dev/dl/)

After installation, check with:

```bash
go version
```

✅ Example output:

```
go version go1.25.4 windows/amd64
```

---

### 📁 Step 2: Clone or Create Project

Create a folder for your SDK:

```bash
mkdir authsecure_go
cd authsecure_go
```

Then initialize a Go module:

```bash
go mod init authsecure_go
```

Project structure should look like this:

```
authsecure_go/
│
├── go.mod
├── main.go
└── authsecure/
    └── authsecure.go
```

---

### 📦 Step 3: Install Dependencies

Run this command to tidy up:

```bash
go mod tidy
```

That’s it — Go will automatically install the required standard libraries.

---

## 💻 Example: Connect Your Application

Below is the **main.go** example showing how to connect your app to **AuthSecure**:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"authsecure_go/authsecure"
)

func main() {
	client := authsecure.AuthSecure{}
	client.Api(
		"XD",                                  // Application Name
		"3ezshCmkXrn",                         // Owner ID
		"7a8bfeb28afcd690812ee5de010a6860",    // Application Secret
		"1.0",                                 // Application Version
	)

	client.Init()

	fmt.Println("\n[1] Login\n[2] Register\n[3] License Login\n[4] Exit")
	fmt.Print("Choose option: ")

	var choice string
	fmt.Scan(&choice)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // clear input buffer

	switch choice {
	case "1":
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		client.Login(strings.TrimSpace(username), strings.TrimSpace(password))

	case "2":
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		fmt.Print("License: ")
		license, _ := reader.ReadString('\n')
		client.Register(strings.TrimSpace(username), strings.TrimSpace(password), strings.TrimSpace(license))

	case "3":
		fmt.Print("License: ")
		license, _ := reader.ReadString('\n')
		client.License(strings.TrimSpace(license))

	default:
		fmt.Println("Goodbye!")
	}
}
```

---

## 🔐 How It Works

1. **Initialization**

   ```go
   client.Api("AppName", "OwnerID", "Secret", "Version")
   client.Init()
   ```

   This verifies your app with AuthSecure and starts a session.

2. **Login**

   ```go
   client.Login("username", "password")
   ```

   Authenticates user credentials and HWID.

3. **Register**

   ```go
   client.Register("username", "password", "license_key")
   ```

   Registers a new user to your app with a valid license key.

4. **License Login**

   ```go
   client.License("license_key")
   ```

   Logs in using license only (no username or password required).

---

## 🧠 HWID Binding (Windows)

HWID is automatically fetched using the user’s **Windows SID**:

```powershell
[System.Security.Principal.WindowsIdentity]::GetCurrent().User.Value
```

So each system has a unique HWID, preventing sharing or spoofing of accounts.

---

## 🧩 Example Output

```
Connecting...
✅ Initialized Successfully!

[1] Login
[2] Register
[3] License Login
[4] Exit
Choose option: 1
Username: jk
Password: jk
✅ Logged in!

👤 User Info:
 Username: jk
 IP: 2a09:bac5:3c0b:1a96::2a6:65
 HWID: S-1-5-21-3116590451-4259102588-3214189088-1001
 Subscriptions:
  - default | Expires: 1762788300 | Left: 82395s
```

---

## 🧱 Integration in Your Application

You can directly embed this SDK in your Go-based software:

* Import the SDK:

  ```go
  import "your_project_path/authsecure"
  ```

* Initialize it with your app’s credentials (from your AuthSecure dashboard).

* Use the available methods:

  * `Init()` → initialize app session
  * `Login()` → login existing user
  * `Register()` → create new user
  * `License()` → login with license key

This works both for **CLI tools**, **desktop applications**, or **server-side verifications**.

---

## 🧩 Example Repo Layout for Your App

If you’re building a loader or desktop app:

```
myapp/
│
├── auth/
│   └── authsecure/     # SDK Folder (from this repo)
│
├── ui/
│   └── main_window.go  # Your application logic
│
└── main.go             # Entry point
```

Then just call:

```go
authclient := authsecure.AuthSecure{}
authclient.Api("MyApp", "ownerid", "secret", "1.0")
authclient.Init()
```

---

## 🪪 License

MIT License © 2025
Free for public & private integration.

---

## 👨‍💻 Author

Built with ❤️ by **@yourgithubusername**
For official documentation: [AuthSecure API Docs](https://authsecure.shop)

---

