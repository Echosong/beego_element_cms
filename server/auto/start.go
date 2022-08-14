package auto

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"kaiyu-web/models"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type View struct {
	Path       string `json:"path"`
	TargetPath string `json:"target_path"`
	BaseName   string `json:"base_name"`
}

type ViewModel struct {
	model interface{}
	// 空就是生成所有，list 只生成list ，add 只生成表达
	typ string
}

const (
	LIST    = "list"
	ADD     = "add"
	SERVICE = "sa-view"
)

var (
	registerModels []ViewModel
	reName         = `description\(([\S^\)]+?)\)`
)

func init() {
	// registerModels = append(registerModels, ViewModel{model: models.Ad{}})
	registerModels = append(registerModels, ViewModel{model: models.UserLog{}})
}

func NewView() {
	v := new(View)
	viewMap, _ := config.GetSection("view")
	if viewMap["path"] == "" {
		logs.Info("请先配置element-ui 相关路径")
		return
	}
	v.Path = viewMap["path"]
	for _, modelBase := range registerModels {
		model := modelBase.model
		rv := reflect.ValueOf(model)
		typ := reflect.TypeOf(model)

		if rv.Kind() != reflect.Struct {
			logs.Info("%v 不是 结构体，不需要处理", model)
			continue
		}
		rvString := rv.String()
		baseName := strings.Split(rvString, " ")[0]
		baseName = strings.Replace(baseName, "<models.", "", -1)
		v.BaseName = strings.ToLower(baseName[:1]) + baseName[1:]
		v.TargetPath = fmt.Sprintf("%v/%v/%v", v.Path, SERVICE, v.BaseName)
		if modelBase.typ == LIST || modelBase.typ == "" {
			v.listView(rv, typ)
		}
		if modelBase.typ == ADD || modelBase.typ == "" {
			v.addView(rv, typ)
		}
	}
}

//路由配置处理
func (v *View) writeConfigJs(typeStr string) {
	perms := fmt.Sprintf("{perms:\"%v-%v\", view: () => import('@/sa-view/%v/%v.vue')},",
		v.BaseName, typeStr, v.BaseName, typeStr)
	configPath := v.Path + "/router.js"
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		logs.Info("读取配置文件router.js错误")
		return
	}
	split := strings.Split(string(bytes), "\n")
	var (
		lines strings.Builder
		flg   bool
	)
	for _, line := range split {
		item := strings.Trim(line, " ")
		if item == perms {
			flg = true
		}
		if item != "]" {
			lines.WriteString(line + "\n")
		}
	}
	if !flg {
		lines.WriteString(perms + "\n")
	}
	lines.WriteString("]")
	wLines := lines.String()

	err = ioutil.WriteFile(configPath, []byte(wLines), 0666)
	if err != nil {
		logs.Info("写入router.js 失败")
	}
}

func (v *View) listView(rv reflect.Value, typ reflect.Type) {
	targetListPath := v.TargetPath + "/list.vue"
	_, err := os.Stat(targetListPath)
	if err == nil {
		logs.Info("文件存在，不需要从新生成 %v", targetListPath)
		return
	}
	var (
		tableColumns []string
		paramsList   []string
		fromList     []string
		fileList     []string
		//exportList []string
		vueData []string
	)

	for i := 0; i < rv.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag
		orm := tag.Get("orm")
		if orm == "" {
			continue
		}
		re := regexp.MustCompile(reName)
		// -1代表取全部
		results := re.FindAllStringSubmatch(orm, -1)
		columnName := results[0][1]
		if field.Name == "SysBase" {
			continue
		}
		selectData, tableStr, fileStr := v.elTableColumn(columnName, field)
		tableColumns = append(tableColumns, tableStr)
		if selectData != "" {
			vueData = append(vueData, selectData)
		}
		if fileStr != "" {
			fileList = append(fileList, fileStr)
		}
		// query: search 查询，no 表示表单不需要
		//需要查询条件才走这里
		queryFrom := field.Tag.Get("query")
		if queryFrom != "" && strings.Contains(queryFrom, "list") {
			form, query := v.elForm(columnName, field)
			fromList = append(fromList, form)
			paramsList = append(paramsList, query)
		}

	}
	listTpl = v.replaceTpl(listTpl)
	listTpl = strings.Replace(listTpl, "//title_obj ", "{}", -1)
	//分页
	paramsList = append(paramsList, "pageSize:10,page:1")
	listTpl = strings.Replace(listTpl, "#{queryPageParams}#", "{"+strings.Join(paramsList, ",")+"}", -1)
	if len(vueData) > 0 {
		listTpl = strings.Replace(listTpl, "//#fieldEnums#", strings.Join(vueData, ","), -1)
	}
	listTpl = strings.Replace(listTpl, "//map_file", strings.Join(fileList, "\r\n"), -1)
	elContent := strings.Join(fromList, "\r\n")
	listTpl = strings.Replace(listTpl, "#{el-form-item}#", elContent, -1)
	listTpl = strings.Replace(listTpl, "#{el-table-column}#", strings.Join(tableColumns, "\r\n"), -1)

	if strings.Contains(elContent, "<input-enum") {
		listTpl = strings.Replace(listTpl, "//import inputEnum from \"../../sa-resources/com-view/input-enum.vue\";", "import inputEnum from \"../../sa-resources/com-view/input-enum.vue\";", -1)
		listTpl = strings.Replace(listTpl, "//inputEnum,", "inputEnum,", -1)
	}
	//创建文件夹
	err = os.MkdirAll(v.TargetPath, 0666)
	file, err := os.Create(targetListPath)
	defer file.Close()
	if err != nil {
		logs.Info("创建文件失败%v ", err)
	}
	_, err = file.WriteString(listTpl)
	if err != nil {
		logs.Info("写入文件失败 %v", err)
		return
	}
	v.writeConfigJs(LIST)
}

//列表查询表单
func (v *View) elForm(columnName string, field reflect.StructField) (formStr string, query string) {
	typeName := field.Tag.Get("json")
	fieldType := fmt.Sprintf("%v", field.Type)
	formStr = ""
	formStr = fmt.Sprintf("<el-form-item label=\"%v:\">", columnName)
	if fieldType == "string" {
		formStr += fmt.Sprintf("<el-input v-model=\"p.%v\" placeholder=\"模糊查询\"></el-input>\n", typeName)
	}
	query = fmt.Sprintf("%v:''", typeName)
	if fieldType == "time.Time" {
		formStr += fmt.Sprintf("<el-date-picker v-model=\"p.start%v\" type=\"datatime\"   value-format=\"yyyy-MM-dd HH:mm:ss\"  placeholder=\"开始日期\"></el-date-picker>\n", field.Name)
		formStr += fmt.Sprintf("<el-date-picker v-model=\"p.end%v\" type=\"datatime\"   value-format=\"yyyy-MM-dd HH:mm:ss\"  placeholder=\"结束日期\"></el-date-picker>\n", field.Name)
		query = fmt.Sprintf("strat%v: '', end%v: ''", typeName, typeName)
	}

	selectTag := field.Tag.Get("type")
	if selectTag == "" {
		selectTag = "input"
	}
	if len(selectTag) > 5 && selectTag[:6] == "select" {
		typeEnumName := typeName + "_enum"
		formStr += fmt.Sprintf("<input-enum :enumData=\"%v\" v-model=\"p.%v\"></input-enum> ", typeEnumName, typeName)
	}
	formStr += "</el-form-item>"
	return formStr, query
}

//列表表格
func (v *View) elTableColumn(columnName string, field reflect.StructField) (selectData string, tableStr string, fileStr string) {
	//正在表达式去列名
	tag := field.Tag
	//类型 type:"input, select([{id:0,name:'男'},{id:1,name:'女'}]), text, upload"
	typ := tag.Get("type")
	if typ == "" {
		typ = "input"
	}
	field.Tag.Get("json")
	typeName := field.Tag.Get("json")

	//列表只需要处理图片 和 下来
	if typ == "select" {
		var reSelect = `select(\s)`
		submatch := regexp.MustCompile(reSelect).FindAllStringSubmatch(typ, -1)
		typeName = typeName + "_enum"
		selectData = typeName + ":" + submatch[0][1] + ","
	}
	tableStr = fmt.Sprintf(" <el-table-column  label=\"%v\"   prop=\"%v\" ></el-table-column>", columnName, typeName)

	if typ == "upload" {
		elStr := `<el-table-column  label="#v0#"  > 
             <template slot-scope="s">
 				<el-link  v-for="item in (s.row.#v1#_file)"  :key="item.name" :href="item.url"  type="primary" >{{item.name}}</el-link>
			 </template>
			</el-table-column>
		`
		tableStr = strings.Replace(strings.Replace(elStr, "#v0#", columnName, -1), "#v1#", typeName, -1)
		fileStr = fmt.Sprintf("item.%v_file=JSON.parse(item.%v)", typeName, typeName)
	}
	return selectData, tableStr, fileStr
}

func (v *View) addView(rv reflect.Value, typ reflect.Type) {
	targetAddPath := v.TargetPath + "/add.vue"
	_, err := os.Stat(targetAddPath)
	if err == nil {
		logs.Info("文件存在，不需要从新生成 %v", targetAddPath)
		return
	}
	var (
		uploadCallback []string
		replaceFiles   []string
		replaceOld     []string
		elFormItems    []string
		ms             []string
	)

	for i := 0; i < rv.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag
		if field.Name == "SysBase" {
			continue
		}
		queryFrom := field.Tag.Get("query")
		if queryFrom != "" && strings.Contains(queryFrom, "no") {
			continue
		}
		orm := field.Tag.Get("orm")
		if orm == "" {
			continue
		}
		re := regexp.MustCompile(reName)
		// -1代表取全部
		results := re.FindAllStringSubmatch(orm, -1)
		columnName := results[0][1]

		fieldType := tag.Get("type")
		if fieldType == "" {
			fieldType = "input"
		}
		fieldName := tag.Get("json")

		if fieldType == "upload" {
			successStr := `success_{fieldName}(response, file, fileList){
				          if(response.code != 200){
				            this.sa.error(response.message);
				            return;
				          }
                          if(!this.m.{fieldName}File){
								this.m.{fieldName}File = [];
                          }
			              this.m.{fieldName}File.push(response.data);
				          console.log(fileList);
				        },`
			removeUploadFun := `remove_{fieldName}(file, fileList){
			   			  this.m.{fieldName}File = fileList;
				      },`
			successStr = strings.Replace(successStr+"\n"+removeUploadFun, "{fieldName}", fieldName, -1)
			uploadCallback = append(uploadCallback, successStr)
			replaceOlderFile := ` this.m.{fieldName} =JSON.stringify(this.m.{fieldName}_file.map(item=>{
				              let a = {};
				              a.name = item.name;
				              a.url = item.url;
				             return a;
				       }).filter(item=>{ if(!item.url.startsWith(\"blob\")){return item;}}) );`
			replaceOld = append(replaceOld, replaceOlderFile)
			replaceFiles = append(replaceFiles, fmt.Sprintf("data.%v_file=JSON.parse(data.%v)", fieldName))
		}
		formStr := fmt.Sprintf("<el-form-item label=\"%v:\">", columnName)
		if len(fieldType) > 5 && fieldType[:6] == "select" {
			typeEnumName := fieldName + "_enum"
			formStr += fmt.Sprintf("<input-enum :enumData=\"%v\" v-model=\"m.%v\"></input-enum> ", typeEnumName, typeEnumName)
			ms = append(ms, fmt.Sprintf("%v:0", fieldName))
		}

		if fieldType == "radio" {
			formStr += fmt.Sprintf("<el-switch v-model=\"m.%v \" ></el-switch>", fieldName)
			ms = append(ms, fmt.Sprintf("%v:0", fieldName))
		}
		if fieldType == "input" {
			mType := fmt.Sprintf("%v", field.Type)
			if mType == "time.Time" {
				formStr += fmt.Sprintf("<el-date-picker v-model=\"m.%v\" type=\"datatime\"   value-format=\"yyyy-MM-dd HH:mm:ss\" ></el-date-picker>\n", fieldName)
			} else {
				if mType[:3] == "int" {
					formStr += fmt.Sprintf("<el-input v-model.number=\"m.%v\"></el-input>\n", fieldName)
				} else {
					formStr += fmt.Sprintf("<el-input v-model=\"m.%v\"></el-input>\n", fieldName)
				}
			}
			ms = append(ms, fmt.Sprintf("%v:''", fieldName))
		}
		if fieldType == "text" {
			formStr += fmt.Sprintf("<el-input type=\"textarea\"  rows=\"2\" placeholder=\"%v\"  v-model=\"m.{%v}\"></el-input>", columnName, fieldName)
			ms = append(ms, fmt.Sprintf("%v:''", fieldName))
		}
		if fieldType == "upload" {
			htmlTypeStr := `<el-upload class="upload-demo" :action="sa.cfg.api_url + '/Attachment/upload?fileType=0&param=0'" :multiple="false" :limit="10"
			              :on-success="success_{fieldName}" :before-remove="remove_{fieldName}"  :file-list="m.{fieldName}_file">
				 <el-button size="mini" type="primary">点击上传</el-button>
                 <div slot="tip" class="el-upload__tip">上传{columnName}</div>
				</el-upload>`
			formStr += strings.Replace(strings.Replace(htmlTypeStr, "{fieldName}", fieldName, -1), "{columnName}", columnName, -1)
			ms = append(ms, fmt.Sprintf("%v:''", fieldName))
		}
		formStr += "</el-form-item>"
		elFormItems = append(elFormItems, formStr)

	}
	fromTpl = v.replaceTpl(fromTpl)
	elContent := strings.Join(elFormItems, "\r\n")
	if strings.Contains(elContent, "<input-enum") {
		fromTpl = strings.Replace(fromTpl, "//import inputEnum from \"../../sa-resources/com-view/input-enum.vue\";", "import inputEnum from \"../../sa-resources/com-view/input-enum.vue\";", -1)
		fromTpl = strings.Replace(fromTpl, "//inputEnum,", "inputEnum,", -1)
	}
	fromTpl = strings.Replace(fromTpl, "#{el-form-item}#", elContent, -1)
	fromTpl = strings.Replace(fromTpl, "#{data_init}#", strings.Join(ms, ",\r\n"), -1)
	fromTpl = strings.Replace(fromTpl, "//upload_functions", strings.Join(uploadCallback, "\r\n"), -1)
	fromTpl = strings.Replace(fromTpl, "//replace_file", strings.Join(replaceFiles, "\r\n"), -1)
	fromTpl = strings.Replace(fromTpl, "//replace_old", strings.Join(replaceOld, "\r\n"), -1)
	fromTpl = strings.Replace(fromTpl, "//rule_fields", "", -1)
	//logs.Info(fromTpl)
	//创建文件夹
	err = os.MkdirAll(v.TargetPath, 0666)
	file, err := os.Create(targetAddPath)
	defer file.Close()
	if err != nil {
		logs.Info("创建文件失败%v", err)
		return
	}
	_, err = file.WriteString(fromTpl)
	if err != nil {
		logs.Info("写入文件失败 %v", err)
		return
	}
	v.writeConfigJs(ADD)
}

func (v *View) replaceTpl(tplContent string) string {
	tplContent = strings.Replace(tplContent, "#{EntityName}#", v.BaseName, -1)
	tplContent = strings.Replace(tplContent, "#{tableInfo}#", "", -1)
	return tplContent
}

var listTpl = `
<template>
    <div class="vue-box">
        <div class="c-panel">
            <!-- 参数栏 -->
            <el-form :inline="true" size="mini">
				<el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-plus" @click="add">增加</el-button>
					<el-button type="primary" icon="el-icon-download" @click="exportFile()">导出</el-button>
                </el-form-item>
                #{el-form-item}#
                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-search" @click="f5();">查询 </el-button>
					
                </el-form-item>
            </el-form>
            <!-- <div class="c-title">数据列表</div> -->
            <el-table :data="dataList" size="mini" ref="multipleTable">
				<!--start-- 需要时候可以随时开启
				<el-table-column type="selection"></el-table-column>
				<el-table-column label="Id" prop="id"></el-table-column>
				-->
                #{el-table-column}#
                <el-table-column prop="address" label="操作" width="220px">
                    <template slot-scope="s">
                        <el-button class="c-btn" type="primary" icon="el-icon-edit" @click="update(s.row)">修改</el-button>
                        <el-button class="c-btn" type="danger" icon="el-icon-delete" @click="del(s.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <div class="page-box">
                <el-pagination
                        background
                        layout="total, prev, pager, next, sizes, jumper"
                        :current-page.sync="p.page"
                        :page-size.sync="p.pageSize"
                        :total="dataCount"
                        :page-sizes="[1, 10, 20, 30, 40, 50, 100]"
                        @current-change="f5(true)"
                        @size-change="f5(true)">
                </el-pagination>
            </div>
        </div>

        <!-- 增改组件 -->
        <add-or-update ref="add-or-update"></add-or-update>
    </div>
</template>

<script>
    //import inputEnum from "../../sa-resources/com-view/input-enum.vue";
	import addOrUpdate from './add.vue';
	export default {
		components: {
			addOrUpdate,
			//inputEnum,
		},
		data() {
			return {
				p:#{queryPageParams}#,
				dataCount: 0,
				dataList: [],
				//#fieldEnums#
			}
		},
		methods: {
			// 数据刷新
			f5: function() {
				this.sa.put("/#{EntityName}#/listPage", this.p).then(res=>{
					this.dataList = res.data.map((item) => {
					    //map_file
						//map_enum
                        return item;
                    });;
					this.dataCount = res.page_data.totalElements;
				});
			},
			// 删除
			del: function(data) {
				this.sa.confirm('是否删除，此操作不可撤销', function() {
					this.sa.delete("/#{EntityName}#/"+ data.id).then(res =>{
						this.sa.arrayDelete(this.dataList, data);
						this.sa.ok(res.message);
					});

				}.bind(this));
			},

			exportFile: function(){
				this.p.pageSize = 1000;
				var titleObj = //title_obj ;
				this.sa.exportFile("/#{EntityName}#/listPage", this.p, titleObj).then(res => {
					console.log(res);
				});
			},

            //更新
			update(row){
				this.$refs['add-or-update'].open(row);
			},
            //添加
			add: function(){
				this.$refs['add-or-update'].open(0);
			}
		},
		created: function(){
			this.f5();
		}
	}

</script>

<style>
</style>

`

var fromTpl = `
<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true"
        :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
                    label-width="120px">
                    #{el-form-item}#
                    <el-form-item>
                        <span class="c-label">&emsp;</span>
                        <el-button type="primary" icon="el-icon-plus" size="small" @click="ok('ruleForm')">确定
                        </el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </el-dialog>
</template>

<script>
//import inputEnum from "../../sa-resources/com-view/input-enum.vue";
export default {
  //components: { inputEnum },
  props: ["params"],
  data() {
    return {
      m: {},
      title:"",
      isShow: false,
      rules: {//rule_fields},
      fileList:[],
    }
  },
  methods: {
    open: function (data) {
      this.isShow = true;
      if (data) {
        this.title = "修改 #{tableInfo}#" ;
        //replace_file
        this.m = data;
      }else{
        this.m = {#{data_init}#}
        this.title = "添加 #{tableInfo}#" ;
      }
    },
    //upload_functions

    //提交#{tableInfo}#信息
    ok: function (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
            //replace_old
            this.sa.post("/#{EntityName}#/save", this.m).then((res) => {
              console.log(res);
              this.$parent.f5();
              this.isShow = false;
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    }
  },
  created() {
  },
};
</script>
`
