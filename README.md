# Mesmerizer

A lightweight CLI tool to send SMS messages via the Twilio API.

```
┌┬┐┌─┐┌─┐┌┬┐┌─┐┬─┐┬┌─┐┌─┐┬─┐
│││├┤ └─┐│││├┤ ├┬┘│┌─┘├┤ ├┬┘
┴ ┴└─┘└─┘┴ ┴└─┘┴└─┴└─┘└─┘┴└─
```

## Features

- Send SMS messages to any phone number via Twilio
- View message history with delivery status
- Color-coded terminal output for better readability
- Cross-platform support (Windows, Linux, macOS)
- Simple configuration with JSON file

## Prerequisites

- [Go](https://golang.org/dl/) 1.19 or higher (for building from source)
- A [Twilio](https://www.twilio.com/) account with:
  - Account SID
  - Auth Token
  - Messaging Service SID

## Installation

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/beyrakIn/mesmerizer.git
   cd mesmerizer
   ```

2. Build the binary:
   ```bash
   go build .
   ```

### Pre-built Binaries

Download the appropriate binary for your platform and create a `config.json` file in the same directory.

## Configuration

Create a `config.json` file in the same directory as the executable:

```json
{
  "url": "https://api.twilio.com/2010-04-01/Accounts/{YOUR_ACCOUNT_SID}/Messages.json",
  "account_sid": "YOUR_ACCOUNT_SID",
  "auth_token": "YOUR_AUTH_TOKEN",
  "messaging_service_sid": "YOUR_MESSAGING_SERVICE_SID"
}
```

| Field | Description |
|-------|-------------|
| `url` | Twilio Messages API endpoint (replace `{YOUR_ACCOUNT_SID}` with your actual Account SID) |
| `account_sid` | Your Twilio Account SID |
| `auth_token` | Your Twilio Auth Token |
| `messaging_service_sid` | Your Twilio Messaging Service SID |

> **Note:** All configuration fields are required. The application will exit with an error if any field is empty.

## Usage

### Send an SMS

```bash
mesmerizer -to "+1234567890" -message "Hello World"
```

### View Message History

```bash
mesmerizer -log
```

### Display Version

```bash
mesmerizer -version
```

### Display Help

```bash
mesmerizer -help
```

### Command-Line Flags

| Flag | Description |
|------|-------------|
| `-to` | Recipient phone number (e.g., `+919999999999`) |
| `-message` | SMS message body (e.g., `"Hello World"`) |
| `-log` | Display all sent messages with delivery status |
| `-version` | Print version and exit |

## Building for Multiple Platforms

Use the included Makefile to build binaries for different platforms:

```bash
make build
```

This creates the following binaries in the `bin/` directory:
- `mesmerizer-windows64.exe` - Windows 64-bit
- `mesmerizer-linux64` - Linux 64-bit
- `mesmerizer-macos64` - macOS 64-bit

## Example Output

### Sending a Message
```
┌┬┐┌─┐┌─┐┌┬┐┌─┐┬─┐┬┌─┐┌─┐┬─┐
│││├┤ └─┐│││├┤ ├┬┘│┌─┘├┤ ├┬┘
┴ ┴└─┘└─┘┴ ┴└─┘┴└─┴└─┘└─┘┴└─

Message sent successfully!
```

### Viewing Message Log
```
┌┬┐┌─┐┌─┐┌┬┐┌─┐┬─┐┬┌─┐┌─┐┬─┐
│││├┤ └─┐│││├┤ ├┬┘│┌─┘├┤ ├┬┘
┴ ┴└─┘└─┘┴ ┴└─┘┴└─┴└─┘└─┘┴└─

From: +1234567890
To: +9876543210
Status: delivered
Body: Hello World
Date Created: 2024-01-01T12:00:00Z
---
```

## Dependencies

- [fatih/color](https://github.com/fatih/color) - Terminal color output

## License

This project is open source and available under the [MIT License](LICENSE).
