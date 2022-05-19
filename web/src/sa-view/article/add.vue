<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" :fullscreen="true" :append-to-body="true"
        :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
                    label-width="120px">
                    <el-form-item label="标题" prop="title">
                        <el-input v-model="m.title" style="width: 400px;"></el-input>
                    </el-form-item>

                    <el-form-item label="分类：" v-if="m.type == 7">
                        <el-cascader :options="categorys" :props="defaultProps" v-model="m.category_id"></el-cascader>
                    </el-form-item>
                    <el-form-item label="排序">
                        <el-input v-model="m.sort"></el-input>
                    </el-form-item>

                    <el-form-item>
                        <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                            :data="{ fileType: 2, params: 1 }" :multiple="false" :limit="1"
                            :on-success="success_imgDetail" :before-remove="remove_imgDetail" list-type="picture-card"
                            :file-list="m.imgDetailFile">
                            <el-button size="small" type="primary">封面图</el-button>
                            <div slot="tip" class="el-upload__tip">
                            </div>
                        </el-upload>
                    </el-form-item>

                    <el-form-item label="简介" prop="info" v-if="m.type != 7">
                        <div style="width: 50%;">
                            <el-input type="textarea" rows="2" placeholder="输入简介" v-model="m.info"></el-input>
                        </div>
                    </el-form-item>

                    <el-form-item label="内容" prop="content">
                        <div style="width: 80%;">
                            <div class="editor-box">
                                <div id="editor" ref="editor" style="text-align:left"></div>
                            </div>
                        </div>
                    </el-form-item>
                    <el-form-item>

                        <el-button type="primary" icon="el-icon-check" size="medium" @click="ok('ruleForm')">确定
                        </el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </el-dialog>
</template>
<style >
.ql-container {
    height: 400px;
}

/* wang富文本编辑器 */
.editor-item {
    width: 100%;
    height: auto;
}

.editor-item .editor-box {
    float: left;
    width: 80%;
    margin-top: 0px;
    margin-left: 0px;
}
</style>
<script>
//import inputEnum from "../../sa-resources/com-view/input-enum.vue";
import E from 'wangeditor';	// 富文本编辑器 
export default {
    //components: { inputEnum },
    props: ["params"],
    data() {
        return {
            m: {},
            title: "",
            isShow: false,
            rules: {
                title: [
                    { required: true, message: "请输入标题", trigger: "blur" },
                ],
                content: [
                    { required: true, message: "请输入内容", trigger: "blur" },
                ],
            },
            fileList: [],
            categorys: [],
            defaultProps: {
                children: "childList",
                label: "name",
                value: "id",
                checkStrictly: true,
                emitPath: false,
            },
        };
    },
    methods: {
        open: function (data) {
            this.isShow = true;
            let titles = { 7: "公司资质", 12: "重要案例", 16: "主营业务" }
            this.m = data;
            this.categorys = data.categorys
            console.log(99888, this.categorys)
            if (data.id) {
                this.title = "修改" + titles[data.type];
                this.m.imgDetailFile = [{url: data.img}]
            } else {
                this.title = "添加 " + titles[data.type];
                this.m.category_id = data.type
            }
            this.create_editor()
        },

        // 创建编辑器 
        create_editor: function () {
            this.$nextTick(function () {
                let editor = new E(this.$refs.editor);
                editor.customConfig.debug = true; // debug模式
                editor.customConfig.uploadFileName = 'file'; // 图片流name
                editor.customConfig.withCredentials = true; // 跨域携带cookie
                editor.customConfig.uploadImgShowBase64 = true   	// 使用 base64 保存图片
                editor.customConfig.onchange = (html) => {	// 创建监听，实时传入 
                    this.m.content = html
                }
                editor.create();		// 创建编辑器 
                editor.txt.html(this.m.content);	// 为编辑器赋值 
            });
        },

        success_imgDetail(response, file, fileList) {
            if (response.code != 200) {
                this.sa.error(response.message);
                return;
            }
            if (!this.m.imgDetailFile) {
                this.m.imgDetailFile = [];
            }
            this.m.imgDetailFile.push(response.data);
            this.m.img = response.data.url;
            console.log("返回1", response.data, this.m.img);
        },
        remove_imgDetail(file, fileList) {
            this.m.imgDetailFile = fileList;
        },
        //提交公告信息
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.m.sort = parseInt(this.m.sort)
                    this.sa.post("/article/save", this.m).then((res) => {
                        console.log(res);
                        this.$parent.f5();
                        this.isShow = false;
                    });
                } else {
                    console.log("error submit!!");
                    return false;
                }
            });
        },
    },
    created() { },
};
</script>