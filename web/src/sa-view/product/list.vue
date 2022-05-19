<template>
    <div class="vue-box">
        <div class="c-panel">
            <el-tabs v-model="activeName" @tab-click="handleClick" v-if="categorys">
                <el-tab-pane :label="item.name" :name="'first_' + item.category_id" v-for="(item, index) in categorys"
                    :key="item.category_id">
                    <el-form size="mini" ref="ruleForm" class="demo-ruleForm" label-width="120px">
                        <el-form-item label="图片">
                            <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                                :data="{ fileType: 2, params: index }" :multiple="false" :limit="1"
                                :on-success="success_imgDetail" list-type="picture-card" :file-list="item.imgDetailFile">
                                <el-button size="small" type="primary">封面图</el-button>
                                <div slot="tip" class="el-upload__tip">
                                </div>
                            </el-upload>
                        </el-form-item>

                        <el-form-item label="简介" prop="info">
                            <el-input style="width: 400px;" type="textarea" :rows="2" placeholder="输入文字信息"
                                v-model="item.info"></el-input>
                        </el-form-item>
                        <el-form-item prop="info" label="内容">
                            <quill-editor v-model="item.content" ref="myQuillEditor" :options="editorOption"
                                @change="onEditorChange($event, item)"></quill-editor>
                        </el-form-item>
                    </el-form>
                </el-tab-pane>
            </el-tabs>
            <div class="btn_send">
                <el-button type="primary" icon="el-icon-check" @click="ok()">确定
                </el-button>
            </div>
        </div>
    </div>
</template>


<style >
.ql-container {
    height: 400px;
}

.btn_send {
    margin: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
<script>
export default {
    data() {
        return {
            content: '34343434',
            activeName: 'second',
            categorys: [],
            editorOption: {
                modules: {
                    toolbar: [
                        ['bold', 'italic', 'underline', 'strike'], // 加粗 斜体 下划线 删除线
                        ['blockquote', 'code-block'], // 引用  代码块
                        [{ header: 1 }, { header: 2 }], // 1、2 级标题
                        [{ list: 'ordered' }, { list: 'bullet' }], // 有序、无序列表
                        [{ script: 'sub' }, { script: 'super' }], // 上标/下标
                        [{ indent: '-1' }, { indent: '+1' }], // 缩进
                        [{ direction: 'rtl' }], // 文本方向
                        [{ size: ['12', '14', '16', '18', '20', '22', '24', '28', '32', '36'] }], // 字体大小
                        [{ header: [1, 2, 3, 4, 5, 6] }], // 标题
                        [{ color: [] }, { background: [] }], // 字体颜色、字体背景颜色
                        // [{ font: ['songti'] }], // 字体种类
                        [{ align: [] }], // 对齐方式
                        ['clean'], // 清除文本格式
                        ['image', 'video'] // 链接、图片、视频
                    ]
                },
                placeholder: '请输入正文'
            }
        };
    },
    methods: {
        handleClick(tab, event) {
            console.log(tab, event);
        },
        onEditorChange({ quill, html, text }, item) {
            console.log('editor change!', quill, html, text, item)
            item.content = html;
        },

        success_imgDetail(response, file, fileList) {
            if (response.code != 200) {
                this.sa.error(response.message);
                return;
            }
            let res = response.data;
            let index = parseInt(res.params)
            if (!this.categorys[index].imgDetailFile) {
                this.categorys[index].imgDetailFile = [];
            }
            this.categorys[index].imgDetailFile = [{ url: res.url }];
            this.categorys[index].img = res.url;
        },
       

        f5() {
            this.sa.put("/category/listPage", { page: 1, pageSize: 10000, parent_id: 20 }).then(res => {
                let categorys = res.data
                categorys = categorys.map(item => {
                    item.category_id = item.id;
                    item.id = null
                    return item
                })
                this.activeName = 'first_' + res.data[0].category_id
                let idArr = []
                categorys.forEach(item => {
                    idArr.push(item.category_id)
                })
                let ids = idArr.join(",")
                this.sa.get("/product/list/" + ids).then(res => {
                    if (res.data) {
                        res.data.forEach(item => {
                            categorys.forEach(c => {
                                if (c.category_id == item.category_id) {
                                    c.content = item.content
                                    c.id = item.id;
                                    c.imgDetailFile =[{url: item.img}]
                                    c.info = item.info
                                }
                            })
                        })
                    }
                    this.categorys = categorys
                })
            })
        },
        ok() {
            this.sa.post("/product/save", this.categorys).then(res => {
                console.log(res)
            })
        }
    },
    created() {
        this.f5()
    }
};
</script>