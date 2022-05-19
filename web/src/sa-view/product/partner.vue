<template>
    <div class="vue-box">
        <div class="c-panel">

            <div style="margin-bottom: 20px;"  v-if="category_id == 24" >
                <el-button type="primary" size="small" @click="dialogFormVisible=true">
                    添加
                </el-button>
            </div>
            <el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab" v-if="products">
                <el-tab-pane :label="item.name" :name="'id_' + item.id" v-for="(item, index) in products" :key="item.id">
                    <div class="tab-content">
                        <el-upload class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                            :data="{ fileType: 2, params: index }" :on-success="success_imgDetail"
                            :before-remove="remove_imgDetail" list-type="picture-card" :file-list="item.contentFile">
                            <i class="el-icon-plus"></i>
                        </el-upload>
                    </div>
                </el-tab-pane>
            </el-tabs>


            <div class="btn_send">
                <el-button type="primary" icon="el-icon-check" @click="ok()">确定
                </el-button>
            </div>
        </div>

        <el-dialog title="设置标题" :modal-append-to-body="false" width="500px" :visible.sync="dialogFormVisible">
            <el-form>
                <el-form-item label="标题">
                    <el-input v-model="productName" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="add()">确 定</el-button>
            </div>

        </el-dialog>
    </div>
</template>


<style >
.ql-container {
    height: 400px;
}

.tab-content {
    margin: 20px;
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
            dialogFormVisible: false,
            editableTabsValue: '2',
            content: '34343434',
            activeName: 'second',
            products: [],
            productName: '',
            constNum : 9999999, //标识自定义添加
            category_id: 24
        };
    },
    methods: {
        success_imgDetail(response, file, fileList) {
            if (response.code != 200) {
                this.sa.error(response.message);
                return;
            }
            let res = response.data;
            let index = parseInt(res.params)
            if (!this.products[index].contentFile) {
                this.products[index].contentFile = [];
            }
            this.products[index].contentFile.push(response.data);
        },
        remove_imgDetail(file, fileList) {
            console.log(file, fileList);
            let index = 0;
            if (file.params == undefined) {
                index = parseInt(file.response.data.params)
            } else {
                index = parseInt(file.params)
            }
            this.products[index].contentFile = fileList;
        },
        add() {
            let newTabName = this.constNum + Math.ceil(Math.random() * 10000);
            this.products.push({
                name: this.productName,
                id: newTabName,
                content: '',
                category_id: 24,
                contentFile: []
            });
            this.editableTabsValue = "id_" + newTabName;
            this.dialogFormVisible = false;
        },
        removeTab(targetName) {
            this.sa.confirm('是否删除，此操作不可撤销', ()=> {
                targetName = parseInt(targetName.replace("id_", ""))
                this.products = this.products.filter(tab => tab.id != targetName);
                //调用服务端删除
                if(targetName < this.constNum){
                    this.sa.delete('/product/'+targetName).then(res=>{
                        this.editableTabsValue = 'id_' + this.products[this.products.length-1].id
                        console.log(res)
                    })
                }
            })
        },

        handleClick(tab, event) {
            console.log(tab, event);
        },
        f5() {
            this.sa.get("/product/list/"+this.category_id).then(res => {
                if (res.data) {
                    this.products = res.data
                }
                if (this.products.length > 0) {
                     this.editableTabsValue = 'id_' + this.products[this.products.length-1].id
                }
                this.products = this.products.map(item => {
                    if (item.content.indexOf("[") == 0) {
                        item.contentFile = JSON.parse(item.content)
                    }
                    return item
                })
            })
        },
        ok() {
            let products = JSON.parse(JSON.stringify(this.products))
            products = products.map(item => {
                if (item.contentFile.length > 0) {
                    item.content = JSON.stringify(item.contentFile.map(img => {
                        return { url: img.url, params: img.params }
                    }))
                }
                if (item.id > this.constNum) {
                    item.id = 0
                }
                return item
            })
            this.sa.post("/product/save", products).then(res => {
               
                this.sa.ok('保存成功')
                this.f5()
            })
        }
    },
    created() {
        this.category_id = parseInt(this.sa_admin.nativeTab.description);
        this.f5()
        
    }
};
</script>