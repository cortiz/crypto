package formaters

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func ProcessOutput(outputType string, operationResult map[string]interface{}) string {
	var expectedOutput string
	switch outputType {
	case "json":
		expectedOutput = jsonOutput(operationResult)
		break
	case "yaml":
		expectedOutput = yamlOuput(operationResult)
		break
	default:
		expectedOutput = outputText(operationResult)
		break
	}

	return expectedOutput
}

func outputText(operationResult map[string]interface{}) string {

	var sb strings.Builder
	for key, value := range operationResult {
		sb.WriteString(fmt.Sprint(value))
		sb.WriteString(" ")
		sb.WriteString(" ")
		sb.WriteString(key)
		sb.WriteString("\n")
	}

	return sb.String()

}
func jsonOutput(operationResult map[string]interface{}) string {

	jsonOutput, err := json.Marshal(operationResult)
	return processMarshall(jsonOutput, err)
}

func yamlOuput(operationResult map[string]interface{}) string {
	yamlOuput, err := yaml.Marshal(operationResult)
	return processMarshall(yamlOuput, err)
}

func processMarshall(marshallOutput []byte, err error) string {
	if err != nil {
		fmt.Printf("Unable to generate a Json Output %s \n", err)
		return ""
	}
	return string(marshallOutput)
}
