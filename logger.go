package logger
// Import ImpressionListener interface
import (
    "fmt"
    "github.com/splitio/go-client/v6/splitio/impressionListener"
)

// Implementation Sample for a Custom Impression Listener
type CustomImpressionListener struct {
}

func (i *CustomImpressionListener) LogImpression(data impressionlistener.ILObject) {
    // Custom behaviour
    fmt.Println("Inside custom split impression logger")
}
