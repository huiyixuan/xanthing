package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
)

// RequestParams 请求参数处理工具类
type RequestParams struct {
	Params map[string]any
}

// NewRequestParams 创建新的请求参数实例
func NewRequestParams(c *gin.Context) (*RequestParams, error) {
	params := &RequestParams{
		Params: make(map[string]any),
	}

	// 合并所有参数
	if err := params.mergeAllParams(c); err != nil {
		return nil, err
	}

	return params, nil
}

// mergeAllParams 合并所有参数来源
func (rp *RequestParams) mergeAllParams(c *gin.Context) error {
	// 1. 合并GET参数
	rp.mergeQueryParams(c)

	// 2. 根据请求方法合并POST参数
	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
		if err := rp.mergePostParams(c); err != nil {
			return err
		}
	}

	// 3. 合并URL参数（如果有的话）
	rp.mergeURLParams(c)

	return nil
}

// mergeQueryParams 合并GET查询参数
func (rp *RequestParams) mergeQueryParams(c *gin.Context) {
	query := c.Request.URL.Query()
	for key, values := range query {
		if len(values) == 1 {
			rp.Params[key] = values[0]
		} else {
			rp.Params[key] = values
		}
	}
}

// mergePostParams 合并POST请求参数
func (rp *RequestParams) mergePostParams(c *gin.Context) error {
	contentType := c.GetHeader("Content-Type")

	// 根据Content-Type处理不同的POST参数格式
	if strings.Contains(contentType, "application/json") {
		return rp.mergeJSONParams(c)
	} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return rp.mergeFormParams(c)
	} else if strings.Contains(contentType, "multipart/form-data") {
		return rp.mergeFormDataParams(c)
	} else {
		// 默认尝试处理为表单数据
		return rp.mergeFormParams(c)
	}
}

// mergeJSONParams 合并JSON格式参数
func (rp *RequestParams) mergeJSONParams(c *gin.Context) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return nil
	}

	var jsonParams map[string]any
	if err := json.Unmarshal(body, &jsonParams); err != nil {
		return err
	}

	// 合并JSON参数到主参数map
	for key, value := range jsonParams {
		rp.Params[key] = value
	}

	return nil
}

// mergeFormParams 合并表单参数（x-www-form-urlencoded）
func (rp *RequestParams) mergeFormParams(c *gin.Context) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}

	for key, values := range c.Request.PostForm {
		if len(values) == 1 {
			rp.Params[key] = values[0]
		} else {
			rp.Params[key] = values
		}
	}

	return nil
}

// mergeFormDataParams 合并multipart/form-data参数
func (rp *RequestParams) mergeFormDataParams(c *gin.Context) error {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32MB max memory
		return err
	}

	if c.Request.MultipartForm != nil {
		// 处理普通表单字段
		for key, values := range c.Request.MultipartForm.Value {
			if len(values) == 1 {
				rp.Params[key] = values[0]
			} else {
				rp.Params[key] = values
			}
		}

		// 处理文件字段（只记录文件名，不处理文件内容）
		for key, files := range c.Request.MultipartForm.File {
			fileNames := make([]string, len(files))
			for i, file := range files {
				fileNames[i] = file.Filename
			}
			if len(fileNames) == 1 {
				rp.Params[key] = fileNames[0]
			} else {
				rp.Params[key] = fileNames
			}
		}
	}

	return nil
}

// mergeURLParams 合并URL参数（如果有的话）
func (rp *RequestParams) mergeURLParams(c *gin.Context) {
	// 如果有URL参数（如 /users/:id），可以在这里处理
	// 例如：id := c.Param("id")
	// rp.Params["id"] = id
}

// Get 获取参数值
func (rp *RequestParams) Get(key string) any {
	if value, exists := rp.Params[key]; exists {
		return value
	}
	return nil
}

// GetString 获取字符串参数值
func (rp *RequestParams) GetString(key string) string {
	value := rp.Get(key)
	if value == nil {
		return ""
	}
	return toString(value)
}

// GetInt 获取整数参数值
func (rp *RequestParams) GetInt(key string) int {
	value := rp.Get(key)
	if value == nil {
		return 0
	}
	return toInt(value)
}

// GetFloat 获取浮点数参数值
func (rp *RequestParams) GetFloat(key string) float64 {
	value := rp.Get(key)
	if value == nil {
		return 0
	}
	return toFloat(value)
}

// GetBool 获取布尔参数值
func (rp *RequestParams) GetBool(key string) bool {
	value := rp.Get(key)
	if value == nil {
		return false
	}
	return toBool(value)
}

// GetSlice 获取切片参数值
func (rp *RequestParams) GetSlice(key string) []any {
	value := rp.Get(key)
	if value == nil {
		return nil
	}

	if slice, ok := value.([]any); ok {
		return slice
	}

	// 如果是字符串切片
	if strSlice, ok := value.([]string); ok {
		result := make([]any, len(strSlice))
		for i, v := range strSlice {
			result[i] = v
		}
		return result
	}

	// 如果是单个值，包装成切片
	return []any{value}
}

// GetAll 获取所有参数
func (rp *RequestParams) GetAll() map[string]any {
	return rp.Params
}

// Has 检查参数是否存在
func (rp *RequestParams) Has(key string) bool {
	_, exists := rp.Params[key]
	return exists
}

// toString 转换为字符串
func toString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprintf("%v", v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// toInt 转换为整数
func toInt(value any) int {
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		// 简单的字符串转数字，实际项目中可以使用更健壮的转换
		var result int
		fmt.Sscanf(v, "%d", &result)
		return result
	default:
		return 0
	}
}

// toFloat 转换为浮点数
func toFloat(value any) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case string:
		var result float64
		fmt.Sscanf(v, "%f", &result)
		return result
	default:
		return 0
	}
}

// toBool 转换为布尔值
func toBool(value any) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		lower := strings.ToLower(v)
		return lower == "true" || lower == "1" || lower == "yes" || lower == "on"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return toInt(value) != 0
	case float32, float64:
		return toFloat(value) != 0
	default:
		return false
	}
}
