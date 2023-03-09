<template>
    <div><h2>我是Hello组件</h2>
    {{ message()}}
    </div>
    {{msg}}
    <button @click="msg='buhao'">按钮</button>
</template>

<script>
export default {
    data(){
        return {
            msg:"nihao"
        }
    },
    inject:["message"],
    beforeCreate(){
        console.log("beforeCreate");
    },
    created(){
        console.log("Created");
    },
    beforeMount(){
        console.log("beforeMount");
    },
    mounted(){
        console.log("mounted");
    },
    beforeUpdate(){
        console.log("beforeUpdate");
    },
    updated(){
        console.log("updated");
    },
    beforeUnmount(){
        console.log("beforeUnmount");
    },
    unmounted(){
        console.log("unmounted");
    }
}

</script>