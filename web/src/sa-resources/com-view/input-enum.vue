<template>
    <el-select v-model="changeValue" @change="selectTo">
        <el-option label="请选择" value=""></el-option>
        <el-option
            :key="item.code"
            v-for="item in enums"
            :value="item.code"
            :label="item.name"
        >
        </el-option>
    </el-select>
</template>

<script>
export default {
    props: {
        enumName: { type: String, default: "" },
        selectValue: { default: 0},
    },


    model: {
        prop: "selectValue",
        event: "event1",
    },
    data() {
        return {
            enums: [],
            changeValue: this.selectValue
        };
    },

    created() {
        this.sa.get("/Home/getEnums?enumName=" + this.enumName).then((res) => {
            this.enums = res.data;
        });
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
};
</script>

<style>
</style>