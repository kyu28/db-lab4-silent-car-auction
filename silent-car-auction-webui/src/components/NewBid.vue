<script setup>
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { ref, defineProps, inject } from 'vue';

const props = defineProps({
    url: String,
});

const url = props.url;
const customerID = inject("customerID");

// Initialize and make Vue listen changes of these variable
let autoID = ref();
let bid = ref();
const router = useRouter();

const successMsg = (msg) => {
    ElMessage({
        message: msg,
        type: 'success'
    });
}

const errorMsg = (msg) => {
    ElMessage.error(msg);
};

// Convert object to query string
function toQuery(params) {
    let str = '';
    for (let key in params) {
        str += key + '=' + params[key] + '&';
    }
    str = str.substring(0, str.length - 1);
    return str;
}

// Submit transaction to server and fetch status message, if success redirect to home page
function submitTransaction() {
    if (isNaN(bid.value) || isNaN(autoID.value) || bid.value <= 0 || autoID.value <= 0) {
        errorMsg("Invalid values");
        return;
    }
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: toQuery({
            customer_id: customerID.value,
            auto_id: autoID.value,
            bid: bid.value,
            bid_date: new Date().toISOString().slice(0, 10)
        })
    })
        .then(response => response.text())
        .then(data => {
            if (data == "success") {
                successMsg("Success");
                router.push("/");
            } else {
                errorMsg(data);
            }
        });
}

// Redirect to login page if not logged in
if (customerID.value == -1) {
    router.push("/");
    errorMsg("Please login first");
}
</script>

<template>
    <div class="container">
        <h2>Place a bid</h2>
        <el-form label-position="top">
            <el-form-item label="Auto ID" required>
                <el-input v-model="autoID"></el-input>
            </el-form-item>
            <el-form-item label="Bid" required>
                <el-input v-model="bid"></el-input>
            </el-form-item>
            <el-form-item style="display: inline-block;" size="default">
                <el-button type="primary" @click="submitTransaction">Pay</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<style scoped>
.container {
    height: 40vh;
    width: 30vw;
}
</style>