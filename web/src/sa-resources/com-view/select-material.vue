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
        productId: { default: '' },
        type: 0,
    },
    model: {
        prop: "productId",
        event: "event1",
    },
    

    data() {
        return {
            datas: [],
            changeValue: this.productId,
            listData:[],
        };
    },
    created() {
        this.sa.get("/material/getMap").then((res) => {
            this.listData = res.data;
            this.datas = res.data.filter((t) => t.type == this.type);
        });
    },

    methods: {
        selectTo() {
            this.$emit("event1", this.changeValue);
            this.$emit("change");
        },
        changeType: function (type) {
            this.datas = this.listData.filter((t) => t.type == type);
            this.changeValue = this.datas[0].id;
        },
    },
};
</script>
<style>
</style>