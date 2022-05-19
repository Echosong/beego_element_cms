<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true"
        :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
                    label-width="120px">
                    <el-form-item label="分类名称" prop="name">
                        <el-input v-model="m.name"></el-input>
                    </el-form-item>
                    <el-form-item label="上级" prop="parentName" :disabled="true">
                        <el-input disabled v-model="m.parentName"></el-input>
                    </el-form-item>
                    <el-form-item label="排序" prop="sort">
                        <el-input v-model="m.sort"></el-input>
                    </el-form-item>

                    <el-form-item v-if="m.parent_id == 0 && m.type != 3">
                        <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                            :data="{ fileType: 2, params: 1 }"  :limit="m.id==1 ?10: 2"
                            :on-success="success_imgDetail" :before-remove="remove_imgDetail" list-type="picture-card"
                            :file-list="m.imgDetailFile">
                            <el-button size="small" type="primary">图片</el-button>
                            <div slot="tip" class="el-upload__tip">
                               添加两个图片： 导航图，banner图
                            </div>
                        </el-upload>
                    </el-form-item>

                    <el-form-item v-if="m.type != 3" label="链接" prop="url">
                        <el-input v-model="m.url"></el-input>
                    </el-form-item>

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
import inputEnum from "../../sa-resources/com-view/input-enum.vue";
export default {
    components: { inputEnum },
    props: ["params"],
    data() {
        return {
            m: {},
            title: "",
            isShow: false,
            rules: {
                name: [
                    {
                        required: true,
                        message: "请输分类名称",
                        trigger: "blur",
                    },
                ],
                sort: [],
                quantity: [],
                type: [],
            },
            fileList: [],
        };
    },
    methods: {
        open: function (data) {
            this.isShow = true;
            if (data.id) {
                this.title = "修改分类";
                this.m = data;
                //修改的时候

            } else {
                this.m = data;
                this.title = "添加分类";

            }
            if (this.m.img.length > 3) {
                this.m.imgDetailFile = JSON.parse(this.m.img)
            } else {
                this.imgDetailFile = []
            }

            console.log("上级id", this.m.parent_id)
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
        },
        remove_imgDetail(file, fileList) {
            this.m.imgDetailFile = fileList;
        },

        //提交商品分类信息
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.m.sort = parseInt(this.m.sort)
                    if (this.m.imgDetailFile && this.m.imgDetailFile.length > 0) {
                        let imgs = [];
                        this.m.imgDetailFile.forEach(item=>{
                            if(item.url.indexOf("upload") != -1){
                                imgs.push({url:item.url})
                            }
                        })
                        this.m.img = JSON.stringify(imgs)
                    }
                    this.sa
                        .post("/category/save", this.m)
                        .then((res) => {
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