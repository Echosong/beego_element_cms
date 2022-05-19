<template>
    <div class="vue-box">
        <!-- 参数栏 -->
        <div class="c-panel">
            <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
                label-width="120px">

                <el-form-item label="我们的服务" prop="text">
                    <el-input style="width: 400px;" type="textarea" :rows="2" placeholder="输入服务内容" v-model="m.text">
                    </el-input>
                </el-form-item>

                <el-form-item label="服务图片" prop="code">
                    <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                        :data="{ fileType: 2, params: 1 }" :multiple="false" :limit="1" :on-success="success_bgDetail"
                        :before-remove="remove_bgDetail" list-type="picture-card" :file-list="bgDetailFile">
                        <el-button size="small" type="primary">封面图</el-button>
                        <div slot="tip" class="el-upload__tip">
                        </div>
                    </el-upload>
                </el-form-item>

                <el-form-item label="实施服务内容" prop="description">
                    <div class="editor-box">
                        <div id="editor" ref="editor" style="text-align:left"></div>
                    </div>
                </el-form-item>

                <el-form-item label="底部文字" prop="foot_content">
                    <el-input style="width: 400px;" type="textarea" :rows="2" placeholder="输入文字信息"
                        v-model="m.foot_content"></el-input>
                </el-form-item>

                <el-form-item label="底部图片">
                    <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                        :data="{ fileType: 2, params: 1 }" :multiple="false" :on-success="success_footDetail"
                        :before-remove="remove_footDetail" list-type="picture-card" :file-list="footDetailFile">
                        <el-button size="small" type="primary">封面图</el-button>
                        <div slot="tip" class="el-upload__tip">
                        </div>
                    </el-upload>
                </el-form-item>

                <el-form-item>
                    <span class="c-label">&emsp;</span>
                    <el-button type="primary" icon="el-icon-plus" size="small" @click="ok('ruleForm')">确定
                    </el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
import selectData from "../../sa-resources/com-view/select-data.vue";
import E from 'wangeditor';	// 富文本编辑器 
export default {
    components: { selectData },
    //components: { inputEnum },

    props: ["params"],
    data() {
        return {
            m: {},
            title: "",
            isShow: false,
            bgDetailFile: [],
            footDetailFile: [],
            rules: {
                name: [
                    {
                        required: true,
                        message: "请输入部门名称",
                        trigger: "blur",
                    },
                ],
                description: [],
            },
            fileList: [],
            category_id: 17
        };
    },
    methods: {
        // 创建编辑器 
        create_editor: function () {
            this.$nextTick(function () {
                let editor = new E(this.$refs.editor);
                editor.customConfig.debug = true; // debug模式
                editor.customConfig.uploadFileName = 'file'; // 图片流name
                editor.customConfig.withCredentials = true; // 跨域携带cookie
                editor.customConfig.uploadImgShowBase64 = true   	// 使用 base64 保存图片
                editor.customConfig.onchange = (html) => {	// 创建监听，实时传入 
                    this.m.content = html;
                }
                editor.create();		// 创建编辑器 
                editor.txt.html(this.m.content);	// 为编辑器赋值 
            });
        },
        f5() {
            this.sa.get("/product/list/" + this.category_id).then(res => {
                if (res.data && res.data.length > 0) {
                    console.log(res.data[0].cotent)
                    this.m = JSON.parse(res.data[0].content);
                    this.m.id = res.data[0].id;
                    this.bgDetailFile = this.m.bg
                    this.footDetailFile = this.m.foot
                    this.create_editor()
                }
            })
        },

        success_bgDetail(response, file, fileList) {
            if (response.code != 200) {
                this.sa.error(response.message);
                return;
            }
            if (!this.bgDetailFile) {
                this.bgDetailFile = [];
            }
            this.bgDetailFile.push(response.data);
        },
        remove_bgDetail(file, fileList) {
            this.bgDetailFile = fileList;
        },

        success_footDetail(response, file, fileList) {
            if (response.code != 200) {
                this.sa.error(response.message);
                return;
            }
            if (!this.footDetailFile) {
                this.footDetailFile = [];
            }
            this.footDetailFile.push(response.data);
        },
        remove_footDetail(file, fileList) {
            this.footDetailFile = fileList;
        },

        //提交部门信息
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    if (this.bgDetailFile && this.bgDetailFile.length > 0) {
                        this.m.bg = this.bgDetailFile.map(item => {
                            return { url: item.url }
                        })
                    }
                    if (this.footDetailFile && this.footDetailFile.length > 0) {
                        this.m.foot = this.footDetailFile.map(item => {
                            return { url: item.url }
                        })
                    }
                    let services = { name: '技术服务', content: JSON.stringify(this.m), category_id: this.category_id, img:"", info:"" , id: this.m.id}
                    this.sa.post("/product", services).then((res) => {
                        this.sa.ok("保持成功")
                    });
                } else {
                    console.log("error submit!!");
                    return false;
                }
            });
        },
    },
    created() {
        this.f5()
     
    },
};
</script>