package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/olekukonko/tablewriter"
)

const (
	apiURL        = "https://indiandata.shop/search.php"
	configFile    = "INDIAN_DATA_SHOP_CONFIG"
	defaultAPIKey = ""
)

type ApiResponse struct {
	Status  string   `json:"status"`
	Persons []Person `json:"data"`
}

type Person struct {
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	FatherName string `json:"father_name"`
	Address    string `json:"address"`
	AltMobile  string `json:"alt_mobile"`
	Circle     string `json:"circle"`
	IDNumber   string `json:"id_number"`
	Email      string `json:"email"`
}

type Config struct {
	APIKey      string `json:"api_key"`
	DisplayType string `json:"display_type"`
}

func loadConfig() (*Config, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}

func saveConfig(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configFile, data, 0644)
}

func configure() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter API_KEY: ")
	apiKey, _ := reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)

	fmt.Print("Enter DISPLAY_TYPE (TABLE or PLAIN): ")
	displayType, _ := reader.ReadString('\n')
	displayType = strings.ToUpper(strings.TrimSpace(displayType))

	cfg := &Config{
		APIKey:      apiKey,
		DisplayType: displayType,
	}

	if err := saveConfig(cfg); err != nil {
		fmt.Println("Failed to save config:", err)
	} else {
		fmt.Println("Configuration saved successfully.")
	}
}

func PrintPeopleTable(people []Person) {
	if len(people) == 0 {
		println("No records to display.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name", "Email", "Mobile", "Alt Mobile",
		"Father", "Circle", "ID Number", "Address",
	})

	for _, p := range people {
		row := []string{
			p.Name,
			p.Email,
			p.Mobile,
			p.AltMobile,
			p.FatherName,
			p.Circle,
			p.IDNumber,
			p.Address,
		}
		table.Append(row)
	}

	table.SetRowLine(true)
	table.Render()
}

func PrintPeopleTablev2(people []Person) {
	if len(people) == 0 {
		fmt.Println("No records to display.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	headers := []string{
		"Name", "Email", "Mobile", "Alt Mobile", "Father", "Circle", "ID Number", "Address",
	}
	fmt.Fprintln(w, strings.Join(headers, "\t"))

	var separator []string
	for _, h := range headers {
		separator = append(separator, strings.Repeat("-", len(h)))
	}
	fmt.Fprintln(w, strings.Join(separator, "\t"))

	for _, p := range people {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			p.Name,
			p.Email,
			p.Mobile,
			p.AltMobile,
			p.FatherName,
			p.Circle,
			p.IDNumber,
			p.Address,
		)
	}
	w.Flush()
}

func main() {
	configureFlag := flag.Bool("configure", false, "Run configuration")
	searchType := flag.String("type", "", "Type of search (email, mobile, aadhar)")
	query := flag.String("query", "", "Search input")
	masked := flag.Bool("masked", false, "Masked result (true/false)")
	flag.Parse()

	if *configureFlag {
		configure()
		return
	}

	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading config. Run with --configure first.")
		os.Exit(1)
	}

	if *searchType == "" || *query == "" {
		fmt.Println("Error: --type and --query are required.")
		flag.Usage()
		os.Exit(1)
	}

	form := url.Values{}
	form.Set("type", *searchType)
	form.Set("query", *query)
	form.Set("masked", fmt.Sprintf("%v", *masked))

	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("API_KEY", cfg.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	var apiResp ApiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(apiResp.Persons) == 0 {
		fmt.Println("No results found.")
		return
	}

	if cfg.DisplayType == "PLAIN" {
		PrintPeopleTablev2(apiResp.Persons)
	} else {
		PrintPeopleTable(apiResp.Persons)
	}
}
