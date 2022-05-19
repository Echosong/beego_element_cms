<template>
    <el-select v-model="changeValue" filterable @change="selectTo">
        <el-option value="">请选择</el-option>
        <el-option v-for="item in products" :key="item.id" :value="item.id" :label="item.name">
        </el-option>
    </el-select>
</template>

<script>
export default {
    props: {
        productId: { default: 0 }
    },
    model: {
        prop: "productId",
        event: "event1",
    },

    data() {
        return {
            products: [],
            changeValue: this.productId
        }
    },
    created() {
        this.sa.get("/product/getMap").then(res => {
            this.products = res.data;
        })
    },
    methods: {
        selectTo() {
            this.$emit("event1", this.changeValue);
        },
    },
}
</script>
<style>
</style>