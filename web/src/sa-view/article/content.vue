<template>
    <div class="vue-box">
        <!-- 参数栏 -->
        <div class="c-panel">
            <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
                label-width="120px">
                <el-form-item label="标题" prop="title">
                    <el-input v-model="m.title" style="width: 400px;"></el-input>
                </el-form-item>

                <el-form-item>
                    <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                        :data="{ fileType: 2, params: 0 }" :multiple="false" :limit="1" :on-success="success_imgDetail"
                        :before-remove="remove_imgDetail" list-type="picture-card" :file-list="m.imgDetailFile">
                        <el-button size="small" type="primary">导图</el-button>
                        <div slot="tip" class="el-upload__tip">
                        </div>
                    </el-upload>
                </el-form-item>

                <el-form-item label="内容" prop="content">
                    <div style="width: 80%;">
                        <div class="editor-box">
                            <div id="editor" ref="editor" style="text-align:left"></div>
                        </div>
                    </div>

                </el-form-item>

                <el-form-item label="视频地址" prop="video">
                    <el-input v-model="m.video" style="width: 400px;"></el-input>
                </el-form-item>

                <el-form-item label="中部业务" prop="bs">
                    <el-row v-for="( item, index) in bs" :key="item.id">
                        <el-col :span="24">
                            <el-input placeholder="输入图标地址" v-model="item.icon" class="about-ico" />
                            <el-input placeholder="输入文字，span着色" v-model="item.text" class="about-text" />
                            <el-button type="text"
                                :icon="index == 0 ? 'el-icon-circle-plus-outline' : 'el-icon-remove-outline'"
                                @click="outLine(index)"></el-button>
                        </el-col>
                    </el-row>
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
<style >
.about-ico {
    margin: 10px;
}

.about-text {
    margin: 10px;
}

.el-input-length {
    width: 300px;
}

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
    props: ["params"],
    data() {
        return {
            m: { code: 'about' },
            bs: [{ ico: "", text: "", id: 0 }],
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
        };
    },
    
    methods: {
        f5() {
            this.sa.get("/content/about").then(res => {
                this.m = res.data
                this.m.imgDetailFile = [{ url: this.m.img }]
                this.m.topImgDetailFile = [{ url: this.m.top_img }]
                this.bs = JSON.parse(this.m.bs)
                this.create_editor()
            })
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
        outLine(index) {
            if (index === 0) {
                this.bs.push({ icon: "", text: "", id: this.bs.length + 1 })
            } else {
                this.bs.splice(index, 1)
            }
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
        },
        remove_imgDetail(file, fileList) {
            this.m.imgDetailFile = fileList;
        },

        success_topImgDetail(response, file, fileList) {
            if (response.code != 200) {
                this.sa.error(response.message);
                return;
            }
            if (!this.m.topImgDetailFile) {
                this.m.topImgDetailFile = [];
            }
            this.m.topImgDetailFile.push(response.data);
            this.m.top_img = response.data.url;
            console.log("返回", response.data, this.m.top_img);
        },
        remove_topImgDetail(file, fileList) {
            this.m.topImgDetailFile = fileList;
        },

        //提交公告信息
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.m.bs = JSON.stringify(this.bs)
                    this.sa.post("/content/save", this.m).then((res) => {
                        this.sa.ok("更新成功")
                        this.isShow = false;
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