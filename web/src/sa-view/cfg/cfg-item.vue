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

        <el-input v-if="item.type == 5" type="textarea" :rows="4" style="width: 60%" :placeholder="item.description"
            v-model="item.value">
        </el-input>

        <template v-if="item.type == 3">
            <el-checkbox-group v-model="item.values">
                <el-checkbox v-for="op in item.options" :key="op.v" :label="op.n">{{ op.n }}
                </el-checkbox>
            </el-checkbox-group>
        </template>

        <el-upload v-if="item.type == 4" class="upload-demo" :action="sa.cfg.api_url + '/file/upload'"
            :data="{ fileType: 2, params: index }" :multiple="false" :limit="1" :on-exceed="handleExceed"
            :on-success="success" list-type="picture-card" :file-list="item.fileList">
            <el-button size="small" type="primary">{{
                    item.name
            }}</el-button>
            <div slot="tip" class="el-upload__tip">
                {{ item.description }}
            </div>
        </el-upload>
    </el-form-item>
</template>

<script>
export default {
    props: {
        configList: { type: Array, default: "" }
    },
    created() {
        console.log("-------", this.configList)
    }
}
</script>