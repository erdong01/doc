<template>
    <div>用户</div>
</template>
<script setup>
import { useRoute } from 'vue-router';
console.log(useRoute().params.id);
const props= defineProps({id:String})
console.log(props.id);
</script>
<script>
export default{
    // props:["id"],
    // mounted(){
    //     // console.log(this.$route.params.id)
    //     console.log(this.$route.params.id);
    //     console.log(this.id);
    // }

}
</script>
<style>

</style>