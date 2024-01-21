<script setup>
</script>

<template>
    <div class="container">
        <transition name="el-fade-in-linear" appear>
            <div class="titles">
                    <h1>Silent Car Auction</h1>
                    <p class="subtitle">Online car auction system</p>
            </div>  
        </transition>   
    </div>
</template>

<style scoped>
.container {
    background: transparent url(/public/bg.jpg) center center no-repeat;
    background-size: cover;
    width: 100%;
}

.el-fade-in-linear-enter-active {
    transition: opacity 2s;
}

.titles {
    backdrop-filter: blur(8px);
    padding-top: 15vh;
    width: 100%;
    height: 100%;
}

.container h1 {
    color: #409eff;
}

.container .subtitle {
    color: #cfd3dc;
}
</style>