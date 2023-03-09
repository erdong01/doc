<template>
    <Hello></Hello>
</template>
<script>
import Hello from './Hello.vue';
export default{
    data() {
        return {
            message: " Hello Word"
        }
    },
    provide() {
        return { message: () => this.message}
    },
    components  :{
        Hello
    }
}
</script>