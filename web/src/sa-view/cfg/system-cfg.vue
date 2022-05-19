/* eslint-disable */
<style scoped>
</style>

<template>
    <div class="vue-box">
        <!-- 参数栏 -->
        <div class="c-panel">
            <el-form size="mini" label-width="120px">
                <template v-for="(item, index) in configList">
                    <el-form-item :label="item.name + ':'" :key="item.id">
                        <el-input v-if="item.type == 0" v-model="item.value" :placeholder="item.description"></el-input>

                        <el-select v-if="item.type == 1" v-model="item.value">
                            <el-option label="请选择" value=""></el-option>
                            <el-option v-for="op in item.options" :key="op.v" :label="op.n" :value="op.n">{{ op.n }}
                            </el-option>
                        </el-select>

                        <el-switch v-if="item.type == 2" v-model="item.value" active-color="#13ce66">
                        </el-switch>

                        <el-input v-if="item.type == 5" type="textarea" :rows="4" style="width: 60%"
                            :placeholder="item.description" v-model="item.value">
                        </el-input>

                        <template v-if="item.type == 3">
                            <el-checkbox-group v-model="item.value">
                                <el-checkbox v-for="op in item.options" :key="op.v" :label="op.n">{{ op.n }}
                                </el-checkbox>
                            </el-checkbox-group>
                        </template>

                        <el-upload v-if="item.type == 4" class="upload-demo" :action="sa.cfg.api_url + '/main/upload'"
                            :data="{ fileType: 2, params: index }" :multiple="false" :limit="1"
                            :on-exceed="handleExceed" :on-success="success" list-type="picture-card"
                            :file-list="item.fileList">
                            <el-button size="small" type="primary">{{
                                item.name
                            }}</el-button>
                            <div slot="tip" class="el-upload__tip">
                                {{ item.description }}
                            </div>
                        </el-upload>
                    </el-form-item>
                </template>

                <el-form-item>
                    <el-button type="primary" size="mini" icon="el-icon-check" @click="ok">保存修改</el-button>
                    <el-button type="primary" size="mini" icon="el-icon-close" @click="f5">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            configList: [],
            m: null,
            checkboxs: [],
            file: "",
            activeName:'tb_2'
        };
    },
    methods: {
        handleExceed() { },
        // f5
        f5: function () {
            this.sa.get("/config/list").then((res) => {
               console.log(res)
                this.configList = res.data.map((item) => {
                    if (item.description.indexOf(",") != -1) {
                        var arrStr = item.description.split(",");
                        var options = [];
                        arrStr.forEach((element, index) => {
                            options.push({ n: element, v: index });
                        });
                        item.options = options;
                    }
                    if (item.type == 3) {
                        item.values = (item.value || "").split(",");
                    }

                    if (item.type == 2) {
                        item.value = item.value == "true" ? true : false;
                    }
                    if (item.type == 4) {
                        item.fileList = [{ name: item.name, url: item.value }];
                    }

                    return item;
                });
                console.log(56565, this.configList)
            });
        },

        success(response, file, fileList) {
            console.log(response, file, fileList);
            //服务端返回做了个传递
            let index = parseInt(response.data.params);
            this.configList[index].fileList[0].url = response.data.url;
        },
        // ok
        ok: function () {
            console.log(this.configList);
            this.configList.map((item) => {
                if (item.type == 3) {
                    item.value = item.values ? item.values.join(",") : "";
                }
                if (item.type == 4) {
                    item.value = item.fileList[0].url;
                }
            });
            this.sa.post("/config/save", this.configList).then((res) => {
                console.log(res);
                this.sa.ok("保存成功");
            });
        },
    },
    created() {
        this.f5();
    },
};
</script>
