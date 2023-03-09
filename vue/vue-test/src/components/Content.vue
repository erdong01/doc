<template>

<div>

    <h2>
        我是组件内容；{{ mseeage }}
    </h2>
    {{ msg }}
    <button @click="msg='test22'">改变msg</button>

    <p></p>
    <button @click="sendPrent">提交数据给父组件</button>
</div>
</template>

<script>
export default{
    data(){
        return {
                msg:"test789",
                bbb:"12345689abc"
        }
    },
    props:{mseeage:{type:String,default:"App Default"}}
    ,
    methods:{
        sendPrent(){
            this.$emit("injectMsg",this.msg)
        }
    }
}
</script>