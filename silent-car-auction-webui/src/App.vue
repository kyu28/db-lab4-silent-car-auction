<script setup>
import { ElMessage } from 'element-plus';
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router';
import { computed, ref, provide } from 'vue';

const key = computed(() => useRoute().fullPath);
const router = useRouter();

const tableRoutes = ref([
  {
    "display": "Auto",
    "path": "auto"
  },
  {
    "display": "Auto Index",
    "path": "auto-index"
  },
  {
    "display": "Customer",
    "path": "customer"
  },
  {
    "display": "Customer Index",
    "path": "customer-index"
  },
  {
    "display": "Transaction",
    "path": "transaction"
  },
  {
    "display": "Winner",
    "path": "winner"
  }
]);

let loggedIn = ref(false);
let customerID = ref(-1); // 0 as admin

const successMsg = (msg) => {
  ElMessage({
    message: msg,
    type: 'success',
    duration: 2000
  });
}

const loginHandler = (id) => {
  loggedIn.value = true;
  customerID.value = id;
  successMsg("Login successfully");
};

const logout = () => {
  loggedIn.value = false;
  customerID.value = -1;
  router.push("/");
  successMsg("Logout successfully");
};

provide("customerID", customerID);
</script>

<template>
  <div class="container">
    <el-menu mode="horizontal">

      <el-menu-item index="0">
        <RouterLink to="/">
          <h2>Silent Car Auction</h2>
        </RouterLink>
      </el-menu-item>

      <el-sub-menu index="1">
        <template #title>Autos</template>
        <el-menu-item index="1-1">
          <RouterLink to="/list/ford">Ford</RouterLink>
        </el-menu-item>
        <el-menu-item index="1-2">
          <RouterLink to="/list/cheverolet">Cheverolet</RouterLink>
        </el-menu-item>
        <el-menu-item index="1-3">
          <RouterLink to="/list/search-autos-by-brand">Search autos by brand</RouterLink>
        </el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="2">
        <template #title>Customer</template>
        <el-menu-item index="2-1">
          <RouterLink to="/list/winners-and-losers">Winners and Losers</RouterLink>
        </el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="3">
        <template #title>Transaction</template>
        <el-menu-item index="3-1">
          <RouterLink to="/list/max-bid">Max bid</RouterLink>
        </el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="4" v-if="loggedIn && customerID == 0">
        <template #title>Table</template>
        <el-menu-item v-for="(item, index) in tableRoutes" :index="'4-' + index">
          <RouterLink :to="'/table/' + item.path">{{ item.display }}</RouterLink>
        </el-menu-item>
      </el-sub-menu>

      <el-menu-item index="5" v-if="loggedIn && customerID != 0">
        <RouterLink to="/transaction/new">Place a bid</RouterLink>
      </el-menu-item>
      <!--Left-->

      <div style="flex-grow: 0.8;"></div>

      <!--Right-->
      <el-menu-item index="6">
        <RouterLink v-if="!loggedIn" to="/login">Login</RouterLink>
        <a v-else @click="logout">Logout</a>
      </el-menu-item>

      <el-menu-item v-if="!loggedIn" index="7">
        <RouterLink to="/signup">Sign up</RouterLink>
      </el-menu-item>

    </el-menu>
    <RouterView @login="loginHandler" :getCustomerID="getCustomerID" :key="key">
    </RouterView>
  </div>
</template>

<style scoped>
.container {
  margin: 0 auto;
  height: 100vh;
  display: flex;
  flex-direction: column;
}

/* RouterLink */
a {
  color: #409eff;
  width: 100%;
}

a:hover {
  color: #66b1ff;
}

RouterView {
  height: 80%;
  flex-grow: 1;
  width: 100%;
}
</style>
