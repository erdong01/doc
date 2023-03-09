<template>
    <div>About</div>
    <button @click="goBack">后退</button>
</template>
<script setup>
</script>
<script>
export default{
    mounted(){
        console.log(this.$route.query.name);
    },
    methods:{
        goBack(){
            this.$router.go(-1)
        }
    }
}
</script>
<style>

</style>