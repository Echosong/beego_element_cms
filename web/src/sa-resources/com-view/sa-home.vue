<style scoped>
.vue-box {
    background-color: #eee;
    font-size: 14px;
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
        "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
}
.content-box {
    margin: 1em;
    padding: 1em;
    background-color: #fff;
    color: #666;
    border-radius: 10px;
}
.el-tag {
    border-radius: 1px;
}
/*800之下*/
@media (max-width: 800px) {
    .kapian {
        width: 100% !important;
        margin-right: 0px !important;
    }
    .saoma {
        display: none;
    }
}
/* markdown格式 */
.md-contetn {
    padding-left: 1.5em;
    line-height: 26px;
}
</style>
<style>
/* 多个ul时，切换那个啥 */
.md-contetn h4 {
    margin-left: -0.5em;
}
.md-contetn ul,
.md-contetn ol {
    padding-left: 1em;
}
.md-contetn pre {
    padding: 5px;
    background-color: #eee;
    font-family: "times new roman";
}
.title {
    width: 120px;
    font-size: 18px;
    font-family: PingFangSC-Semibold, PingFang SC;
    font-weight: 600;
    color: #2d2d2d;
    line-height: 29px;
    padding-left: 20px;
    border-left: 3px solid #4983f4;
}

.header {
    margin: 20px;
    margin-bottom: 10px;
}

.project {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
}
.project-title {
    font-size: 16px;
    font-weight: 600;
    justify-content: flex-start;
    align-items: flex-start;
    display: flex;
}

.project-number {
    display: flex;
    width: 92px;
    height: 40px;
    font-size: 24px;
    font-family: PingFangSC-Semibold, PingFang SC;
    font-weight: 600;
    color: #4983f4;
    line-height: 40px;
    justify-content: center;
    align-items: center;
}

.index-task-title {
    font-size: 16px;
    font-family: PingFangSC-Semibold, PingFang SC;
    font-weight: 600;
    margin-top: 10px;
    display: flex;
    flex-direction: row;
}

.index-task-title-item {
    margin-left: 20px;
    margin-right: 20px;
}

.project-task-add {
    margin: 10px;
    height: 67px;
    background: #ffffff;
    border-radius: 8px;
    border: 1px solid #e2e2e2;
    display: flex;
    padding-left: 60px;
    justify-content: flex-start;
    align-items: center;
    color: #979797;
    flex-direction: row;
    margin-bottom: 20px;
}

.project-task-add span {
    margin-left: 20px;
    font-size: 20px;
}

.index-item {
    display: flex;
    justify-content: flex-start;
    align-items: flex-start;
    border-bottom: #eee 2px solid;
    padding: 10px;
    flex-direction: column;
    margin-top: 10px;
}

.index-item-circle {
    height: 18px;
    width: 18px;
    border-radius: 9px;
    color: #fff;
    margin-right: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.index-item-b {
    margin-top: 12px;
    display: flex;
    flex-direction: row;
    margin-left: 30px;
}

.el-radio__inner {
    height: 24px;
    width: 24px;
}

.active {
    color: #4983f4;
}

.index-item-title {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 90%;
}

.index-item-dev {
    display: flex;
    flex-direction: row;
    justify-content: left;
}
.index-item-info {
    word-break: break-all;
    white-space: pre-wrap;
    width: 90%;
}

.index-item-username {
    margin-left: 20px;
    width: 250px;
    font-weight: bold;
}
.content-card {
    display: flex;
    justify-content: center;
    align-items: center;
}

.finance {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}
.finance-item {
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

.finance-title {
    font-size: 14px;
    font-weight: bold;
    margin: 10px;
}

.finance-money {
    font-size: 16px;
    font-weight: bold;
}
.finance-info {
    margin-top: 10px;
    margin-bottom: 20px;
    color: #999;
}

.finance-ge {
    width: 1px;
    background-color: #ddd;
    height: 50px;
}

.finance-item img {
    width: 50px;
}

.kuai-list {
    width: 100%;
    min-height: 50px;
    margin-top: 10px;
    display: flex;
    justify-content: left;
    flex-wrap: wrap;
}
</style>

<template>
    <div class="vue-box">

        <div class="header">
            <el-row>
                <el-col :span="12">
                    <div class="title">工作台</div>
                </el-col>

            </el-row>
        </div>
        
    </div>
</template>

<script>
import * as echarts from "echarts";
export default {
    data() {
        return {
            enums: [],
            all: 0,
            value: 0,
            color: ["#E02020", "#F74000", "#FF7200", "#FFB103"],
            p: { pageSize: 10, page: 1, state: 0, userId: null },
            dataCount: 0,
            dataList: [],
            user: {},
            group: [],
            timer: "",
            quickLinks: [
                { name: "成品管理", link: "product-list", id: 1 },
                { name: "供应商管理", link: "supplier-list", id: 2 },
                { name: "采购单", link: "purchaseOrder-list", id: 3 },
                { name: "成品库存", link: "productStock-list", id: 4 },
                { name: "财务申请", link: "purchasePayRequest-list", id: 5 },
                { name: "我的任务", link: "task-list", id: 6 },
                { name: "入库单管理", link: "receipt-list", id: 7 },
            ],
            allData: {},
            financeOut: 0,
            financeIn: 1,
            tableData: [],
        };
    },
    created() {
        this.user = this.sa_admin.user;

        this.f5();
    },
    mounted() {
      
    },

    methods: {
        f5: function () {
        
        },
        send(index) {
            this.p.state = index;
            this.f5();
        },
        success(row) {
           
        },
        quick(row) {
            console.log(row.link);
            this.sa_admin.navigateTo(row.link);
        },
    },
};
</script>

