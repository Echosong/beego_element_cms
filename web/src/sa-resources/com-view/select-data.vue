<template>
    <el-select v-model="changeValue" filterable @change="selectTo">
        <el-option value="">请选择</el-option>
        <el-option v-for="item in datas" :key="item.id" :value="item.id" :label="item.name">
        </el-option>
    </el-select>
</template>

<script>
export default {
    props: {
        indexId: { default: ''},
        routeName: {default: ""},
        type: {default: 0}
    },
    model: {
        prop: "indexId",
        event: "event1",
    },

    data() {
        return {
            datas: [],
            changeValue: this.indexId
        }
    },
    created() {
        this.sa.get("/"+this.routeName+"/getMap?type="+this.type).then(res => {
            this.datas = res.data;
        })
    },
    methods: {
        selectTo() {
            this.$emit("event1", this.changeValue);
            this.$emit("change");
        },

        setValue(value) {
            this.changeValue = value;
        }
    },
}
</script>
<style>
</style>