<script setup>
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { ref, defineEmits } from 'vue';

// Initialize and make Vue listen changes of these variable
let customerID = ref("");
// Emit results to App.vue to update login status
const emits = defineEmits(["login"]);

const router = useRouter();

const errorMsg = (msg) => {
    ElMessage.error(msg);
};

function sendLoginData() {
    if (isNaN(customerID.value) || customerID.value < 0) {
        errorMsg("User ID must be a positive integer");
        return;
    }
    emits("login", customerID.value);
    router.push("/");
}
</script>

<template>
    <div class="container">
        <h2>Login</h2>
        <el-form >
            <el-form-item label="User ID">
                <el-input v-model="customerID"></el-input>
            </el-form-item>
            <el-form-item style="display: inline-block;">
                <el-button type="primary" @click="sendLoginData" :disabled="customerID.length == 0">Login</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<style scoped>
.container {
    padding-top: 15vh;
    height: 40vh;
    width: 30vw;
}
</style>