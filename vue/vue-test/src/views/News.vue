<template>
    <div>新闻</div>
</template>
<script setup>
import { useRoute } from 'vue-router';
console.log(useRoute().params.id);
</script>
<script>
export default{
    mounted(){
        // console.log(this.$route.params.id)
    }
}
</script>
<style>

</style>