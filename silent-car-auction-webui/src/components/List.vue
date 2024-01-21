<script setup>
import { ElMessage } from 'element-plus';
import { ref, defineProps } from 'vue';
const props = defineProps({
    list: String,
    cols: Array,
    url: String,
    hasArgs: Boolean
});

const list = props.list;
const cols = props.cols;
const url = props.url;
const hasArgs = props.hasArgs;
// Initialize and make Vue listen changes of these variable
let rows = ref([]);
let arg = ref("");

const errorMsg = (msg) => {
    ElMessage.error(msg);
};

// Call db function and fetch returned list
function fetchList() {
    fetch(hasArgs ? url + "/" + arg.value : url, { method: "GET"})
        .then(response => {
                if (!response.ok) {
                    response.text().then(text => errorMsg(text));
                }
                return response;
            })
        .then(response => response.json())
        .then(data => rows.value = data);
}

// Fetch list on component mount
if (!hasArgs) {
    fetchList();
}
</script>

<template>
    <div class="container">
        <h1>{{ list }}</h1>
        <div v-if="hasArgs" class="query-area">
            <el-input v-model="arg" style="width: 40%; margin: 0 0.5vw;"></el-input>
            <el-button type="primary" v-bind:disabled="arg.length == 0" @click="fetchList">Query</el-button>
        </div>
        <el-collapse-transition>
            <el-table v-if="rows.length > 0" :data="rows" class="result" size="small" stripe=true max-height="400">
                <el-table-column v-for="col in cols" :prop="col" :label="col"></el-table-column>
            </el-table>
        </el-collapse-transition>
        <p v-if="rows.length == 0">No data fetched</p>
    </div>
</template>

<style scoped>
h1 {
    font-size: 8vh;
    margin: 2vh 0;
}

.container {
    padding-top: 3vh;
    display: block;
    width: 80%;
    color: white;
    text-align: center;
}   

.query-area {
    margin: 2vh 0;
}

.result {
    width: 100%;
    border-radius: 2vh;
}
</style>
