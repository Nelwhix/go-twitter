package twitter

import (
	"fmt"
	"os"
	"encoding/json"
)

func ddJSON(value interface{}) {
	jsonData, _ := json.Marshal(value)

	fmt.Fprintln(os.Stdout, string(jsonData))
}