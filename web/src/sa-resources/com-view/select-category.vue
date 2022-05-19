<template>
    <el-cascader :options="categorys" :props="defaultProps" v-model="changeValue" @change="selectTo"></el-cascader>
</template>

 
<script>
import sa_admin_code_util from "../index/admin-util.js"; // admin代码util
export default {
    props: {
        categoryId: { default: "" },
        type: { default: 1 },
    },
    model: {
        prop: "categoryId",
        event: "event1",
    },

    data() {
        return {
            categorys: [],
            changeValue: this.categoryId,
            defaultProps: {
                children: "childList",
                label: "name",
                value: "id",
                checkStrictly: true,
                emitPath: false,
            },
        };
    },
    created() {
        this.sa
            .put("/ProductCategory/ListPage", {
                page: 1,
                pageSize: 1000,
                type: this.type,
            })
            .then((res) => {
                this.categorys = sa_admin_code_util.arrayToTree(
                    res.data.content
                ); // 一维转tree
            });
    },
    methods: {
        selectTo() {
            this.$emit("event1", this.changeValue);
        },
        setValue(value) {
            this.changeValue = value;
        }
    },
};
</script>
<style>
</style>