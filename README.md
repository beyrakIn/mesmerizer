# Mesmerizer

Mesmerizer is a simple CLI tool to send SMS over Twilio API.

## Installation

Change the configuration in `config.json` and then use the tool or build your binary for your OS.
If you download the tool as a binary, you should create a `config.json` file in the same directory as you can see below.
```json
{
  "url": "",
  "account_sid": "",
  "auth_token": "",
  "messaging_service_sid": ""
}
```

```bash
go build .
```

## Usage
```mesmerizer -help```
```markdown
┌┬┐┌─┐┌─┐┌┬┐┌─┐┬─┐┬┌─┐┌─┐┬─┐
│││├┤ └─┐│││├┤ ├┬┘│┌─┘├┤ ├┬┘
┴ ┴└─┘└─┘┴ ┴└─┘┴└─┴└─┘└─┘┴└─
Usage of ./mesmerizer:
  -log
    	Log messages sent, if this flag is set, mobile and message flags are ignored
  -message string
    	Message to send, eg: 'Hello World'
  -to string
    	Mobile number to send SMS, eg: +919999999999
  -version
    	Print version and exit

```

In order to send SMS, you need to pass the mobile number and message to send.

```bash
mesmerizer -to '' -message ''
```
To see messages sent by mesmerizer

```bash
mesmerizer -log
```
