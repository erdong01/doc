<template>
    <h2>Page页面</h2>
    <div>
        <button @click="gopage">跳转页面</button>
    </div>
</template>
<script>
export default{
    methods:{
        gopage(){
            this.$router.push({name:'news',params:{id:123}})
            
        }
    }
}
</script>