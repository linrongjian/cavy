package xlsx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"cavy/common/mlog"

	"github.com/tealeg/xlsx"
)

type Xlsx struct {
	data []interface{}
	Log  *mlog.Logger
}

// 检查最小值和最大值
func checkFloatMinAndMax(tags map[string]string, value float64) error {
	minStr, has := tags["min"]
	if has {
		min, _ := strconv.ParseFloat(minStr, 64)
		if value < min {
			return fmt.Errorf("value %f less min %f", value, min)
		}
	}

	maxStr, has := tags["max"]
	if has {
		max, _ := strconv.ParseFloat(maxStr, 64)
		if value > max {
			return fmt.Errorf("value %f > max %f", value, max)
		}
	}

	return nil
}

// 检查最小值和最大值
func checkInt64MinAndMax(tags map[string]string, value int64) error {
	minStr, has := tags["min"]
	if has {
		min, _ := strconv.ParseInt(minStr, 10, 64)
		if value < min {
			return fmt.Errorf("value %d < main %d", value, min)
		}
	}

	maxStr, has := tags["max"]
	if has {
		max, _ := strconv.ParseInt(maxStr, 10, 64)
		if value > max {
			return fmt.Errorf("value %d > max %d", value, max)
		}
	}

	return nil
}

// 检查最小值和最大值
func checkIntMinAndMax(tags map[string]string, value int) error {
	minStr, has := tags["min"]
	if has {
		min, _ := strconv.Atoi(minStr)
		if value < min {
			return fmt.Errorf("value %d < min %d", value, min)
		}
	}

	maxStr, has := tags["max"]
	if has {
		max, _ := strconv.Atoi(maxStr)
		if value > max {
			return fmt.Errorf("value %d > min %d", value, max)
		}
	}

	return nil
}

// 获得字段值
func getValue(tags map[string]string, kind reflect.Kind, cell *xlsx.Cell) (interface{}, error) {
	switch kind {
	case reflect.Int, reflect.Int32:
		value, err := cell.Int()
		if err != nil {
			return 0, nil
		}

		if err := checkIntMinAndMax(tags, value); err != nil {
			return nil, fmt.Errorf("type:%s checkMinAndMax err:%v", kind.String(), err)
		}

		return value, nil
	case reflect.Int64:
		value, err := cell.Int64()
		if err != nil {
			return 0, nil
		}

		if err := checkInt64MinAndMax(tags, value); err != nil {
			return nil, fmt.Errorf("type:%s checkIntMinAndMax err:%v", kind.String(), err)
		}

		return value, nil
	case reflect.Float32, reflect.Float64:
		value, err := cell.Float()
		if err != nil {
			return 0.0, nil
		}

		if err := checkFloatMinAndMax(tags, value); err != nil {
			return nil, fmt.Errorf("Type:%s checkMinAndMax err:%v", kind.String(), err)
		}

		return value, nil
	case reflect.String:
		return cell.String(), nil
	case reflect.Bool:
		return cell.Bool(), nil
	}
	return nil, nil
}

// 获得列信息
func getColumnTag(tag string) []int {
	content := strings.Split(tag, ",")
	columns := []int{}
	for _, v := range content {
		value, _ := strconv.Atoi(v)
		columns = append(columns, value)
	}
	return columns
}

// 获得标签
func getTags(field reflect.StructField) map[string]string {
	tagstr := field.Tag.Get("xlsx")
	tags := strings.Split(tagstr, " ")
	tagFields := map[string]string{}
	for _, v := range tags {
		tmp := strings.Split(v, ":")
		if len(tmp) < 2 {
			tagFields[tmp[0]] = ""
			continue
		}
		value := tmp[1]
		tagFields[tmp[0]] = value
	}

	return tagFields
}

func (s *Xlsx) parseRow(row *xlsx.Row, tableType reflect.Type, startColumn int, endColumn int) (interface{}, int) {
	//item := map[string]interface{}{}
	item := reflect.New(tableType).Elem()
	columnCount := 0
	for i := 0; i < tableType.NumField(); i++ {
		field := tableType.Field(i)
		tag := getTags(field)
		columnStr, ok := tag["column"]
		if !ok {
			continue
		}
		columns := getColumnTag(columnStr)
		if field.Type.Kind() == reflect.Slice {
			if field.Type.Elem().Kind() == reflect.Struct {
				if len(row.Cells) <= 0 {
					continue
				}
				sliceType := reflect.MakeSlice(field.Type, len(row.Cells), len(row.Cells))
				index := 0
				for columns[0] < columns[1] {
					value, count := s.parseRow(row, field.Type.Elem(), columns[0]+startColumn, columns[1])
					columns[0] += count
					columnCount += count
					sliceValue := sliceType.Index(index)
					sliceValue.Set(reflect.Indirect(reflect.ValueOf(value)))
					index++
					if index >= len(row.Cells) {
						break
					}
				}
				tmpSliceType := reflect.MakeSlice(field.Type, index, index)
				for i := 0; i < index; i++ {
					value := tmpSliceType.Index(i)
					value.Set(sliceType.Index(i))
				}
				item.FieldByName(field.Name).Set(tmpSliceType)
			} else if len(columns) > 1 {
				valueType := field.Type.Elem().Kind()
				sliceType := reflect.MakeSlice(field.Type, len(columns), len(columns))
				for i, v := range columns {
					if v+startColumn < len(row.Cells) {
						value, err := getValue(tag, valueType, row.Cells[v+startColumn])
						if err != nil {
							s.Log.Errorf("Name:%s getValue err:%v", field.Name, err)
							continue
						}
						sliceValue := sliceType.Index(i)
						sliceValue.Set(reflect.ValueOf(value))
					}
					columnCount++
				}
				item.FieldByName(field.Name).Set(sliceType)
			} else {
				continue
			}
		} else if field.Type.Kind() == reflect.Struct {
			value, count := s.parseRow(row, field.Type, columns[0]+startColumn, columns[1])
			columnCount += count
			item.FieldByName(field.Name).Set(reflect.ValueOf(value))
		} else {
			index := columns[0] + startColumn
			if index < len(row.Cells) {
				value, err := getValue(tag, field.Type.Kind(), row.Cells[columns[0]+startColumn])
				if err != nil {
					s.Log.Errorf("Name:%s getValue err:%v", field.Name, err)
					continue
				}
				item.FieldByName(field.Name).Set(reflect.ValueOf(value))
			}
			columnCount++
		}

	}
	return item.Interface(), columnCount
}

// NewSessionBinary 通过二进制打开
func NewSessionBinary(fileData []byte, sheetName string, tableHeaderRowCount int, table interface{}) *Xlsx {
	session := &Xlsx{
		data: []interface{}{},
		Log:  mlog.NewLogger(mlog.Fields{"SheetName": sheetName}),
	}

	tableType := reflect.TypeOf(table)
	if tableType.Kind() == reflect.Ptr {
		tableType = tableType.Elem()
	}

	xlsxFile, err := xlsx.OpenBinary(fileData)
	if err != nil {
		session.Log.Errorf("OpenFile err:%v", err)
		return nil
	}

	sheet, has := xlsxFile.Sheet[sheetName]
	if !has {
		return nil
	}
	for line, row := range sheet.Rows {
		if line < tableHeaderRowCount {
			continue
		}
		data, _ := session.parseRow(row, tableType, 0, 0)
		session.data = append(session.data, data)
	}

	return session
}

// NewSessionBinary 通过二进制打开
func NewSessionBinaryReflect(fileData []byte, sheetName string, tableHeaderRowCount int, tableType reflect.Type) *Xlsx {
	session := &Xlsx{
		data: []interface{}{},
		Log:  mlog.NewLogger(mlog.Fields{"SheetName": sheetName}),
	}

	xlsxFile, err := xlsx.OpenBinary(fileData)
	if err != nil {
		session.Log.Errorf("OpenFile err:%v", err)
		return nil
	}

	sheet, has := xlsxFile.Sheet[sheetName]
	if !has {
		return nil
	}

	fmt.Println(tableType)
	for line, row := range sheet.Rows {
		if line < tableHeaderRowCount {
			continue
		}
		data, _ := session.parseRow(row, tableType, 0, 0)
		session.data = append(session.data, data)
	}

	return session
}

// NewSession 创建表格
func NewSession(path string, sheetName string, tableHeaderRowCount int, table interface{}) *Xlsx {
	session := &Xlsx{
		data: []interface{}{},
		Log:  mlog.NewLogger(mlog.Fields{"Path": path, "SheetName": sheetName}),
	}

	tableType := reflect.TypeOf(table)
	if tableType.Kind() == reflect.Ptr {
		tableType = tableType.Elem()
	}

	xlsxFile, err := xlsx.OpenFile(path)
	if err != nil {
		session.Log.Errorf("OpenFile err:%v", err)
		return nil
	}

	sheet, has := xlsxFile.Sheet[sheetName]
	if !has {
		return nil
	}
	for line, row := range sheet.Rows {
		if line < tableHeaderRowCount {
			continue
		}
		data, _ := session.parseRow(row, tableType, 0, 0)
		session.data = append(session.data, data)
	}

	return session
}

func (s *Xlsx) getArray(valueType reflect.Type) reflect.Value {
	sliceValue := reflect.MakeSlice(valueType, len(s.data), len(s.data))
	for i, v := range s.data {
		pv := reflect.New(valueType.Elem()).Elem()
		pv.Set(reflect.ValueOf(v))
		sliceValue.Index(i).Set(reflect.Indirect(pv))
	}
	return sliceValue
}

func (s *Xlsx) getMap(valueType reflect.Type) reflect.Value {
	key := ""
	itemType := valueType.Elem().Elem()
	fmt.Println(valueType)
	fmt.Println(itemType)
	for i := 0; i < itemType.NumField(); i++ {
		field := itemType.Field(i)
		tags := getTags(field)
		if _, has := tags["pk"]; has {
			key = field.Name
			break
		}
	}
	if key == "" {
		return reflect.MakeMap(valueType)
	}

	mapValue := reflect.MakeMap(valueType)
	for _, v := range s.data {
		pv := reflect.New(itemType).Elem()
		fmt.Println(pv.Type())
		fmt.Println(reflect.ValueOf(v).Type())
		pv.Set(reflect.ValueOf(v))

		mapValue.SetMapIndex(reflect.Indirect(reflect.ValueOf(v)).FieldByName(key), pv.Addr())
	}
	return mapValue
}

func (s *Xlsx) getMapBk(valueType reflect.Type) reflect.Value {
	key := ""
	itemType := valueType.Elem()
	fmt.Println(valueType)
	fmt.Println(itemType)
	for i := 0; i < itemType.NumField(); i++ {
		field := itemType.Field(i)
		tags := getTags(field)
		if _, has := tags["pk"]; has {
			key = field.Name
			break
		}
	}
	if key == "" {
		return reflect.MakeMap(valueType)
	}

	mapValue := reflect.MakeMap(valueType)
	for _, v := range s.data {
		pv := reflect.New(valueType.Elem()).Elem()
		fmt.Println(pv.Type())
		fmt.Println(reflect.ValueOf(v).Type())
		pv.Set(reflect.ValueOf(v))
		mapValue.SetMapIndex(reflect.Indirect(reflect.ValueOf(v)).FieldByName(key), reflect.Indirect(pv))
	}
	return mapValue
}

// Get 获得数据需要传入数组
func (s *Xlsx) Get(value interface{}) error {
	data := reflect.Indirect(reflect.ValueOf(value))
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}

	if data.Kind() == reflect.Slice {
		data.Set(s.getArray(data.Type()))
	} else if data.Kind() == reflect.Map {
		data.Set(s.getMap(data.Type()))
	} else {
		return fmt.Errorf("value Type is not slice or map")
	}

	return nil
}

// Get 获得数据需要传入数组
func (s *Xlsx) GetMap(value reflect.Value) error {
	data := reflect.Indirect(value)
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}

	if data.Kind() == reflect.Slice {
		data.Set(s.getArray(data.Type()))
	} else if data.Kind() == reflect.Map {
		fmt.Println(data.Type())
		data.Set(s.getMap(data.Type()))
	} else {
		return fmt.Errorf("value Type is not slice or map")
	}

	return nil
}
