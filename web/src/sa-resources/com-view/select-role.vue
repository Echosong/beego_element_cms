<template>
    <el-cascader :options="options" :props="defaultProps" v-model="changeValue" @change="selectTo" ></el-cascader>
</template>

<script>
export default {
    props: { roleId: { default: 1 } },

    model: {
        prop: "roleId",
        event: "event1",
    },
    data() {
        return {
            options: [],
            changeValue: this.roleId,
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
            .put("/role/listPage", { page: 1, pageSize: 10000 })
            .then((res) => {
                console.log(res);
                this.options = res.data
            });
    },
    methods: {
        selectTo() {
            this.$emit("event1", this.changeValue);
        },
        setValue(value) {
            this.changeValue = value;
        },
    },
};
</script>

<style>
</style>