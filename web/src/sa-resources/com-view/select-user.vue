<template>
    <el-select v-model="changeValue" filterable   @change="selectTo">
        <el-option value="">请选择</el-option>
        <el-option v-for="item in users" :key="item.id" :label="item.name" :value="item.id">
        </el-option>
    </el-select>
</template>

<script>
export default {
    props: { userId: { default: 0 } },

    model: {
        prop: "userId",
        event: "event1",
    },
    data() {
        return {
            users: [],
            changeValue: this.userId
        }
    },
    created() {
        this.sa.put("/user/listPage",{page:1, pageSize:10000}).then(res => {
            this.users = res.data.content;
        });
    },
    methods: {
        selectTo() {
            this.$emit("event1", this.changeValue);
            this.$emit("userChange");   
        },
    },
}
</script>

<style>
</style>