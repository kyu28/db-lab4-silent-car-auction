<script setup>
import { ElMessage } from 'element-plus';
import { ref, defineProps } from 'vue';
const props = defineProps({
    table: String,
    immutableCols: Array,
    mutableCols: Array,
    baseURL: String
});

const GET = 0;
const POST = 1;
const PATCH = 2;
const DELETE = 3;

// Initialize and make Vue listen changes of these variable 
let table = props.table;
let immutableCols = props.immutableCols;
let mutableCols = props.mutableCols;
const baseURL = props.baseURL;
const rows = ref([]);
const crud = ref(GET);
const updateColumn = ref('');
const updateValue = ref('');

const successMsg = (msg) => {
    ElMessage({
        message: msg,
        type: 'success'
    });
};
const errorMsg = (msg) => {
    ElMessage.error(msg);
};

// Concatenate url and params
function concatParams(url, params) {
    if (Object.keys(params).length > 0) {
        url += '?' + toQuery(params);
    }
    return url;
}

// Convert object to query string
function toQuery(params) {
    let str = '';
    for (let key in params) {
        str += key + '=' + params[key] + '&';
    }
    str = str.substring(0, str.length - 1);
    return str;
}

// Do CRUD operation and fetch results
function submitToDB() {
    let url = baseURL + '/' + table;
    const methods = ['GET', 'POST', 'PATCH', 'DELETE'];
    // UPDATE (PATCH) has special URL
    if (crud.value == PATCH) {
        url += '/' + updateColumn.value + '/' + updateValue.value;
    }
    // Construct params
    let params = {};
    document.getElementById("conditions").querySelectorAll("input").forEach(function (input) {
        if (input.value != '') {
            params[input.name] = input.value;
        }
    });
    // Params in URI except for INSERT (POST)
    if (crud.value != POST) {
        url = concatParams(url, params);
    } else {
        if (!isInsertFormLegal(params)) {
            errorMsg('Please fill in all fields');
            return false;
        }
    }
    fetch(url, { method: methods[crud.value],
        // POST needs body and headers
        body: crud.value == POST ? toQuery(params) : null,
        headers: crud.value == POST ? { 'Content-Type': 'application/x-www-form-urlencoded' } : {},
    })
        .then(response => {
                if (!response.ok) {
                    response.text().then(text => errorMsg(text));
                }
                return response;
            })
        // Only parse JSON for SELECT
        .then(crud.value == GET ? response => response.json() : response => response.text())
        .then(crud.value == GET ?
        data => {
            rows.value = data;
        } :
        data => {
            if (data == 'success') {
                successMsg('Success, please GET to refresh the result');
                rows.value = [];
            } else {
                errorMsg(data);
            }
        });
}

// Check if all fields are filled in INSERT (POST) form
function isInsertFormLegal(params) {
    for (let i = 0; i < mutableCols.length; i++) {
        if (!(mutableCols[i] in params)) {
            return false;
        }
    }
    return true;
}

// Clear all input fields
function clearForm() {
    document.getElementById("conditions").querySelectorAll("input").forEach(function (input) {
        input.value = '';
    });
    document.getElementById("update").querySelectorAll("input").forEach(function (input) {
        input.value = '';
    });
}
</script>

<template>
    <div class="table-container">
        <div class="left-panel">
            <h1>Table {{ table }}</h1>
            <el-radio-group v-model="crud">
                <el-radio-button :label="GET">Select</el-radio-button>
                <el-radio-button :label="POST">Insert</el-radio-button>
                <el-radio-button :label="PATCH">Update</el-radio-button>
                <el-radio-button :label="DELETE">Delete</el-radio-button>
            </el-radio-group>
            <form id="update" v-if="crud == PATCH">
                <div class="col_row">
                    <span class="col_name">column</span>&nbsp;
                    <el-select class="col_input" placeholder="Select a column" v-model="updateColumn">
                        <el-option v-for="col in mutableCols" :key="col" :value="col" :label="col"/>
                    </el-select>
                </div>
                <div class="col_row">
                    <span class="col_name" id="tag_value">value</span>&nbsp;
                    <input type="text" class="col_input" v-model="updateValue"/><br/>
                </div>
            </form>
            <hr v-if="crud == PATCH" width="85%"/>
            <form id="conditions">
                <div v-if="crud != POST">
                    <div v-for="col in immutableCols" class="col_row">
                        <span class="col_name" v-bind:name="col" v-bind:id="'tag_' + col">{{ col }}</span>&nbsp;
                        <input type="text" class="col_input" v-bind:name="col" v-bind:id="'input_' + col"/><br/>
                    </div>
                </div>
                <div v-for="col in mutableCols" class="col_row">
                    <span class="col_name" v-bind:name="col" v-bind:id="'tag_' + col">{{ col }}</span>&nbsp;
                    <input type="text" class="col_input" v-bind:name="col" v-bind:id="'input_' + col"/><br/>
                </div>
            </form>
            <el-button type="primary" @click="clearForm()">Clear</el-button>
            <el-button type="success" @click="submitToDB()">Submit</el-button>
        </div>
        <div class="right-panel">
            <el-collapse-transition>
                <el-table v-if="rows.length > 0" :data="rows" class="result" size="small" stripe=true max-height="400">
                    <el-table-column v-for="col in immutableCols" :prop="col" :label="col"></el-table-column>
                    <el-table-column v-for="col in mutableCols" :prop="col" :label="col"></el-table-column>
                </el-table>
            </el-collapse-transition>
            <p v-if="rows.length == 0">No data fetched</p>
        </div>
    </div>
</template>

<style scoped>
h1 {
    font-size: 8vh;
    margin: 2vh 0;
}

form {
    width: 100%;
    padding: 2vh 0;
}

.table-container {
    padding-top: 5vh;
}

.col_row {
    margin: 1vh auto;
    display: block;
    width: 80%;
}

.col_name {
    display: inline-block;
    width: 40%;
}

.col_input {
    display: inline-block;
    width: 50%;
}

.left-panel {
    display: block;
    float: left;
    width: 30%;
    color: white;
    text-align: center;
}

.right-panel {
    display: block;
    float: left;
    width: 65%;
    color: white;
    text-align: center;
}   

.result {
    width: 100%;
    border-radius: 2vh;
}
</style>
