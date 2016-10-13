package stringbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

func buildConcat(itemType, clientId, ID string) string {
	return itemType + ":" + clientId + ":" + ID
}

func buildSprintf(itemType, clientId, ID string) string {
	return fmt.Sprintf("%s:%s:%s", itemType, clientId, ID)
}

func buildJoin(itemType, clientId, ID string) string {
	return strings.Join([]string{itemType, clientId, ID}, ":")
}

func buildBuffer(itemType, clientId, ID string) string {
	l := len(itemType) + len(clientId) + len(ID) + 2
	buf := make([]byte, 0, l)
	w := bytes.NewBuffer(buf)
	w.WriteString(itemType)
	w.WriteRune(':')
	w.WriteString(clientId)
	w.WriteRune(':')
	w.WriteString(ID)
	return w.String()
}
