<template>
   <div class="c-panel" v-if="project.id!=-1">
            <el-row justify="center">
                <el-col :span="4">
                    <el-row>
                        <el-col>项目名称</el-col>
                        <el-col style="font-weight: bold;margin-top: 10px;">{{project.name}}</el-col>
                    </el-row>
                </el-col>
                <el-col :span="4">
                    <el-row>
                        <el-col>项目编号：</el-col>
                        <el-col style="font-weight: bold;margin-top: 10px;">{{project.code}}</el-col>
                    </el-row>
                </el-col>
                <el-col :span="4">
                    <el-row>
                        <el-col>项目地点：</el-col>
                        <el-col style="font-weight: bold;margin-top: 10px;">{{project.address}}</el-col>
                    </el-row>
                </el-col>
                <el-col :span="4">
                    <el-row>
                        <el-col>项目交付日期：</el-col>
                        <el-col style="font-weight: bold; margin-top: 10px;">{{project.finishDate}}</el-col>
                    </el-row>
                </el-col>
                <el-col :span="8" style=" display: flex;flex-direction: row;justify-content: center;margin-top: 5px;">
                    <el-button type="primary" size="small" @click="goto(project)" plain>查看更多</el-button>
                </el-col>
            </el-row>

        </div>
</template>

<script>
export default {
    props:{
        project: {default:{}}
    },
    methods:{
        goto(project){
            this.sa_admin.navigateTo("project-detail", project);
        }
    }
}
</script>

<style>

</style>