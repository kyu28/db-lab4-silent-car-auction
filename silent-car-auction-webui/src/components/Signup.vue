<script setup>
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { ref, defineProps } from 'vue';

const props = defineProps({
    baseURL: String
});

const baseURL = props.baseURL;

// Initialize and make Vue listen changes of these variable 
let first_name = ref("");
let last_name = ref("");
let location = ref("");
let email = ref("");
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

// Concatenate baseURL and params
function concatParams(baseURL, params) {
    if (Object.keys(params).length > 0) {
        baseURL += '?' + toQuery(params);
    }
    return baseURL;
}

// Submit customer data to server and fetch status message, if success redirect to home page
function submitCustomerData() {
    if (first_name.value.length == 0 || last_name.value.length == 0 || location.value.length == 0 || email.value.length == 0) {
        errorMsg("Please fill in all fields");
        return;
    }
    fetch(baseURL + "/customer", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: toQuery({
            first_name: first_name.value,
            last_name: last_name.value,
            location: location.value,
            email: email.value
        })
    })
        .then(response => response.text())
        .then(data => {
            if (data == "success") {
                getCustomerID();
                router.push("/");
            } else {
                errorMsg("Failed to sign up");
            }
        });
}

// Get the new customer ID
function getCustomerID() {
    fetch(concatParams(baseURL + "/customer-index", { email: email.value }), { method: "GET" })
        .then(response => {
                if (!response.ok) {
                    response.text().then(text => errorMsg(text));
                }
                return response;
            })
        .then(response => response.json())
        .then(data => {
            if (data.length == 0) {
                errorMsg("Failed to get customer ID");
            } else {
                successMsg("Sign up success, your customer ID is " + data[0].customer_id);
            }
        });
}
</script>

<template>
    <div class="container">
        <h2>Sign up</h2>
        <el-form label-position="top" size="small">
            <el-form-item label="First Name" required>
                <el-input v-model="first_name"></el-input>
            </el-form-item>
            <el-form-item label="Last Name" required>
                <el-input v-model="last_name"></el-input>
            </el-form-item>
            <el-form-item label="Location" required>
                <el-input v-model="location"></el-input>
            </el-form-item>
            <el-form-item label="Email" required>
                <el-input v-model="email"></el-input>
            </el-form-item>
            <el-form-item style="display: inline-block;" size="default">
                <el-button type="primary" @click="submitCustomerData">Sign up</el-button>
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