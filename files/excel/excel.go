package excel

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ReadeExcel(address string) ([]string, []map[string]string, error) {
	cols := make([]string, 0, 100)
	ret := make([]map[string]string, 0, 100)
	f, err := excelize.OpenFile(address)
	if err != nil {
		return cols, ret, err
	}
	defer f.Close()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return cols, ret, err
	}
	for i, row := range rows {
		if i == 0 { //取得第一行的所有数据---execel表头
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
		} else {
			theRow := map[string]string{}
			for j, colCell := range row {
				k := cols[j]
				theRow[k] = colCell
			}
			ret = append(ret, theRow)
		}
	}
	return cols, ret, nil
}

type Config struct {
	Id         int64  `db:"id" json:"id,omitempty" goqu:"pk,skipinsert,skipupdate"`
	EventType  string `db:"Event_Type" json:"eventType,omitempty" goqu:"defaultifempty"`
	Category   string `db:"Category" json:"category,omitempty" goqu:"defaultifempty"`
	Status     string `db:"Status" json:"status,omitempty" goqu:"defaultifempty"`
	Area       string `db:"Area" json:"area,omitempty" goqu:"defaultifempty"`
	Node       string `db:"Node" json:"node,omitempty" goqu:"defaultifempty"`
	Module     string `db:"Module" json:"module,omitempty" goqu:"defaultifempty"`
	Parameter  string `db:"Parameter" json:"parameter,omitempty" goqu:"defaultifempty"`
	State      string `db:"State" json:"state,omitempty" goqu:"defaultifempty"`
	UpdateTime string `db:"Update_Time" json:"updateTime,omitempty" goqu:"defaultifempty"`
}

func Export(path string, dbData []*Config) error {
	f := excelize.NewFile()
	sheet := "报表1"
	// Create a new sheet.
	index, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	_ = f.SetCellValue(sheet, "A1", "Id")
	_ = f.SetCellValue(sheet, "B1", "EventType")
	_ = f.SetCellValue(sheet, "C1", "Category")
	_ = f.SetCellValue(sheet, "D1", "Area")
	_ = f.SetCellValue(sheet, "E1", "Node")
	_ = f.SetCellValue(sheet, "F1", "Module")
	_ = f.SetCellValue(sheet, "G1", "Parameter")
	_ = f.SetCellValue(sheet, "H1", "State")
	line := 1
	for _, row := range dbData {
		line++
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", line), row.Id)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", line), row.EventType)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", line), row.Category)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", line), row.Area)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", line), row.Node)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", line), row.Module)
		_ = f.SetCellValue(sheet, fmt.Sprintf("G%d", line), row.Parameter)
		_ = f.SetCellValue(sheet, fmt.Sprintf("H%d", line), row.State)
	}
	f.SetActiveSheet(index)

	sheet2 := "报表2"
	// Create a new sheet.
	index2, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	_ = f.SetCellValue(sheet2, "A1", "Id")
	_ = f.SetCellValue(sheet2, "B1", "EventType")
	_ = f.SetCellValue(sheet2, "C1", "Category")
	_ = f.SetCellValue(sheet2, "D1", "Area")
	_ = f.SetCellValue(sheet2, "E1", "Node")
	_ = f.SetCellValue(sheet2, "F1", "Module")
	_ = f.SetCellValue(sheet2, "G1", "Parameter")
	_ = f.SetCellValue(sheet2, "H1", "State")
	line2 := 1
	for _, row := range dbData {
		line++
		_ = f.SetCellValue(sheet2, fmt.Sprintf("A%d", line2), row.Id)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("B%d", line2), row.EventType)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("C%d", line2), row.Category)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("D%d", line2), row.Area)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("E%d", line2), row.Node)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("F%d", line2), row.Module)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("G%d", line2), row.Parameter)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("H%d", line2), row.State)
	}
	_ = f.DeleteSheet("Sheet1")
	f.SetActiveSheet(index2)
	if err := f.SaveAs(path); err != nil {
		return err
	}
	return nil
}

func ExportAndDownload(ctx *gin.Context, path string, dbData []*Config) error {
	f := excelize.NewFile()
	sheet := "报表1"
	// Create a new sheet.
	index, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	_ = f.SetCellValue(sheet, "A1", "Id")
	_ = f.SetCellValue(sheet, "B1", "EventType")
	_ = f.SetCellValue(sheet, "C1", "Category")
	_ = f.SetCellValue(sheet, "D1", "Area")
	_ = f.SetCellValue(sheet, "E1", "Node")
	_ = f.SetCellValue(sheet, "F1", "Module")
	_ = f.SetCellValue(sheet, "G1", "Parameter")
	_ = f.SetCellValue(sheet, "H1", "State")
	line := 1
	for _, row := range dbData {
		line++
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", line), row.Id)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", line), row.EventType)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", line), row.Category)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", line), row.Area)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", line), row.Node)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", line), row.Module)
		_ = f.SetCellValue(sheet, fmt.Sprintf("G%d", line), row.Parameter)
		_ = f.SetCellValue(sheet, fmt.Sprintf("H%d", line), row.State)
	}
	f.SetActiveSheet(index)

	sheet2 := "报表2"
	// Create a new sheet.
	index2, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	_ = f.SetCellValue(sheet2, "A1", "Id")
	_ = f.SetCellValue(sheet2, "B1", "EventType")
	_ = f.SetCellValue(sheet2, "C1", "Category")
	_ = f.SetCellValue(sheet2, "D1", "Area")
	_ = f.SetCellValue(sheet2, "E1", "Node")
	_ = f.SetCellValue(sheet2, "F1", "Module")
	_ = f.SetCellValue(sheet2, "G1", "Parameter")
	_ = f.SetCellValue(sheet2, "H1", "State")
	line2 := 1
	for _, row := range dbData {
		line++
		_ = f.SetCellValue(sheet2, fmt.Sprintf("A%d", line2), row.Id)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("B%d", line2), row.EventType)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("C%d", line2), row.Category)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("D%d", line2), row.Area)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("E%d", line2), row.Node)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("F%d", line2), row.Module)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("G%d", line2), row.Parameter)
		_ = f.SetCellValue(sheet2, fmt.Sprintf("H%d", line2), row.State)
	}
	_ = f.DeleteSheet("Sheet1")

	f.SetActiveSheet(index2)

	disposition := fmt.Sprintf("attachment; filename=%s %s.xlsx", url.PathEscape("解决中文乱码"), "2020-01-01 12:12:12")
	ctx.Writer.Header().Set("Content-Disposition", disposition)
	ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition") //允许暴露给客户端
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")

	f.SetActiveSheet(index)
	//ctx.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	_ = f.Write(ctx.Writer)

	return nil
}

const statusDeviceStatisticBaseTableName = "导出excel"

//导出excel
//func (c *communicationStatisticBlockService) ExportStatistic(ctx gin.Context, req *v1.ExportStatisticReq) (res v1.ExportStatisticListApiRes, err error) {
//	//根据时间区间查询数据
//	var beginYm = ""
//	var endYm = ""
//	//如果未传递时间区间 默认本月
//	if req.BeginAt == "" || req.EndAt == "" {
//		beginYm = gtime.Now().Format("ym")
//		endYm = gtime.Now().Format("ym")
//	} else {
//		beginYm = gtime.NewFromStr(req.BeginAt).Format("ym")
//		endYm = gtime.NewFromStr(req.EndAt).Format("ym")
//	}
//	var beginTableName = statusDeviceStatisticBaseTableName + "_" + beginYm
//	var endTableName = statusDeviceStatisticBaseTableName + "_" + endYm
//	f := excelize.NewFile()
//	// 创建一个工作表
//	index, err := f.NewSheet("Sheet1")
//	if err != nil {
//		return nil, err
//	}
//	//合并单元格
//	err = f.MergeCell("Sheet1", "A1", "E1")
//	if err != nil {
//		return nil, err
//	}
//	err = f.MergeCell("Sheet1", "A2", "E2")
//	if err != nil {
//		return nil, err
//	}
//	//设置列宽
//	err = f.SetColWidth("Sheet1", "A", "E", 25)
//	if err != nil {
//		return nil, err
//	}
//	//设置表头内容
//	err = f.SetCellValue("Sheet1", "A1", "电表数据明细")
//	if err != nil {
//		return nil, err
//	}
//	err = f.SetCellValue("Sheet1", "A2", ""+req.BeginAt+"至"+req.EndAt+"")
//	if err != nil {
//		return nil, err
//	}
//	err = f.SetCellValue("Sheet1", "A3", "电表编号")
//	if err != nil {
//		return nil, err
//	}
//	err = f.SetCellValue("Sheet1", "B3", "起始读数")
//	if err != nil {
//		return nil, err
//	}
//	err = f.SetCellValue("Sheet1", "C3", "抄表读数")
//	if err != nil {
//		return nil, err
//	}
//	err = f.SetCellValue("Sheet1", "D3", "实际用电量")
//	if err != nil {
//		return nil, err
//	}
//	err = f.SetCellValue("Sheet1", "E3", "所属区域")
//	if err != nil {
//		return nil, err
//	}
//	//获取设备列表
//	lists := ([]*entity.Device)(nil)
//	devicelist := dao.Device.Ctx(ctx)
//	err = devicelist.Scan(&lists)
//	if err != nil {
//		return nil, err
//	}
//	//遍历设备
//	for i, v := range lists {
//		curIndex := i + 4
//		strIndex := strconv.Itoa(curIndex)
//		//设置表内内容
//		err = f.SetCellValue("Sheet1", "A"+strIndex, v.Name)
//		if err != nil {
//			return nil, err
//		}
//		result, _ := g.DB().GetAll(ctx, "select total_active_energy from "+beginTableName+" where device_sn = "+v.DeviceSn+" and created_at >= "+req.BeginAt+" order by total_active_energy  ASC")
//		result2, _ := g.DB().GetAll(ctx, "select total_active_energy from "+endTableName+" where device_sn = "+v.DeviceSn+" and created_at >= "+req.EndAt+" order by total_active_energy  DESC")
//		if len(result) > 0 && len(result2) > 0 {
//			first := result[0]
//			//设置起始读数
//			err = f.SetCellValue("Sheet1", "B"+strIndex, first["total_active_energy"])
//			if err != nil {
//				return nil, err
//			}
//			last := result2[0]
//			//设置抄表读数
//			err = f.SetCellValue("Sheet1", "C"+strIndex, last["total_active_energy"])
//			if err != nil {
//				return nil, err
//			}
//			//
//			//cha := last["total_active_energy"].Float64() - first["total_active_energy"].Float64()
//			//f.SetCellValue("Sheet1", "D"+strIndex, cha)
//		}
//		//设置设备所属范围
//		result3, _ := g.DB().GetOne(ctx, "select name from meter_group where group_id = ?", v.GroupId)
//		re := result3["name"]
//		err = f.SetCellValue("Sheet1", "E"+strIndex, re)
//		if err != nil {
//			return nil, err
//		}
//		//设置表格样式，垂直水平居中
//		styleOne, err := f.NewStyle(
//			`{"alignment":{"horizontal":"center","vertical":"center"}}`,
//		)
//		if err != nil {
//			fmt.Println(err)
//		}
//		// 设置样式范围
//		err = f.SetCellStyle("Sheet1", "A1", "E"+strIndex, styleOne)
//		if err != nil {
//			return nil, err
//		}
//		//设置表格样式，垂直水平居中，字体加粗
//		styleTwo, err := f.NewStyle(
//			`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`,
//		)
//		if err != nil {
//			fmt.Println(err)
//		}
//		// 设置样式范围
//		err = f.SetCellStyle("Sheet1", "A1", "E3", styleTwo)
//		if err != nil {
//			return nil, err
//		}
//	}
//	r := ghttp.RequestFromCtx(ctx).Response
//	// 设置工作簿的默认工作表
//	f.SetActiveSheet(index)
//	//设置表格名
//	filename := "work.xlsx"
//	// 返回数据到客户端
//	r = ghttp.RequestFromCtx(ctx).Response
//	r.Header().Add("Content-Type", "application/octet-stream")
//	r.Header().Add("Content-Disposition", "attachment; filename="+filename)
//	r.Header().Add("Content-Transfer-Encoding", "binary")
//	//下载表格
//	_ = f.Write(r.Writer)
//	return
//}
