<template>

    <div>
        <el-button @click="selectTo(item)" style=" margin-left: 30px; padding-left: 30px;padding-right: 30px;" 
            :type="item.code == val?  'primary':''"
         v-for="item in enums" :value="item.code" :key="item.code" round size="small">{{item.name}}({{item.count}})</el-button>
    </div>

</template>

<script>
export default {
    props: {
        enumName: { type: String, default: "" }
    },
    data() {
        return {
            enums: [],
            val : 1
        };
    },

    created() {
       this.sa.get("/project/getCountGroupByState").then(res => {
            this.enums = res.data;
        })
    },
    methods: {
        selectTo(item) {
            console.log(7777, item.code);
            this.val = item.code;
            this.$emit('fatherMethod', item.code);
           // this.$parent.ok(item.code)
        },
    },
};
</script>

<style>
   
</style>