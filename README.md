# ifttt-webhook

IFTTT Webhook library for Go

## Usage

```go
i := iftttWebhook.New("<your-api-key>")
i.Emit("<your-event-name>", "value1", "value2", "value3")
```
