<template>
    <el-select v-model="changeValue" filterable @change="selectTo">
        <el-option value="">请选择</el-option>
        <el-option v-for="item in projects" :key="item.id" :value="item.id" :label="item.name">
        </el-option>
    </el-select>
</template>

<script>
export default {
    props: {
        projectId: { default: 0 }
    },
    model: {
        prop: "projectId",
        event: "event1",
    },

    data() {
        return {
            projects: [],
            changeValue: this.projectId
        }
    },
    created() {
        this.sa.get("/project/getMap").then(res => {
            this.projects = res.data;
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