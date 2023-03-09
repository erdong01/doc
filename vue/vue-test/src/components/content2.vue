<template>
    <div>
        <h2>我插槽content组件内容1</h2>
        <div>
            <!-- <slot name="aaa"></slot>
            <slot name="bbb"></slot>
            <slot name="ccc"></slot> -->
        </div>
        <slot :list="list"></slot>
    </div>
</template>
<script>
export default{
    data(){
        return {
            list:[1,2,3,4]
        }
    }
}
</script>