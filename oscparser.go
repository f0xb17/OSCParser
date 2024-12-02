package oscparser

import (
	"bufio"
	"os"
	"regexp"
)

// he read_file function reads an .osc file and extracts all the variable names present in the file.
// It returns a map where the keys are the variable names and the values are boolean true.
func Read_File(filePath string, pattern string) (map[string]bool, error) {
	// Open the file and return an error when the file cannot be opened!
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// Close the file after reading
	defer file.Close()

	// Create a new map - String, Bool
	variables := make(map[string]bool)

	// Compile the regex for (S.L.variablename), (L.L.variablename) or (S.$.variablename), (L.$.variablename)
	regex := regexp.MustCompile(pattern)

	// Create a Scanner to read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Store current line in line
		line := scanner.Text()
		// Find match in current line
		matches := regex.FindAllStringSubmatch(line, -1)
		// Add the extracted value to map
		for _, match := range matches {
			variables[string(match[2])] = true
		}
	}

	// Check for errors that occured during the scanning process and return error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Return everything
	return variables, err
}
