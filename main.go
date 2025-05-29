package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. 打开CSV文件
	csvFile, err := os.Open("lawzhidao_filter.csv")
	if err != nil {
		log.Fatal("无法打开CSV文件:", err)
	}
	defer csvFile.Close()

	// 2. 读取CSV内容
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("读取CSV错误:", err)
	}

	// 3. 将CSV转换为JSON
	if len(records) < 1 {
		log.Fatal("CSV文件没有内容")
	}

	// 第一行作为键名
	headers := records[0]
	var result []map[string]string

	for i, row := range records {
		if i == 0 {
			continue // 跳过标题行
		}

		entry := make(map[string]string)
		for j, value := range row {
			if j < len(headers) {
				entry[headers[j]] = value
			}
		}
		result = append(result, entry)
	}

	// 4. 转换为JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal("JSON编码错误:", err)
	}

	// 5. 输出JSON
	fmt.Println(string(jsonData))

	// 可选：将JSON写入文件
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal("写入JSON文件错误:", err)
	}

	fmt.Println("CSV已成功转换为JSON并保存到output.json")
}
