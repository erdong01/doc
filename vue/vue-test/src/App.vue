<script setup>
import { computed, reactive, ref, watch, watchEffect } from 'vue';
import Content from './components/Content.vue';
import Content2 from './components/Content2.vue';
import HelloWorld from './components/HelloWorld.vue';
import HomeView from './components/HomeView.vue';
import TheWelcome from './components/TheWelcome.vue';
import useMainStore from './store/index.js';
const store =useMainStore();

const counter = ref(0)
const user = reactive({ name: "张三", age: 18 })

function changeCounter() {
  counter.value++
}

function changeUserName() {
  user.name = "李四"
}
watch(counter, (newValue, oldValue) => {
  console.log(newValue, oldValue);
})
watchEffect(() => {
  console.log(user.name);
})

const msg = ref("hellow word")
const reverseMsg = computed(() => {
  return msg.value.split('').reverse().join('')
})
console.log(reverseMsg.value)


</script>
<script>
export default {
  data() {
    return {
      msg: "App Data",
      childMsg: "childMsg Data"
    }
  },
  methods: {
    getChildMsg(val) {
      console.log(val)
    }
  },
  mounted() {
    console.log(this.$refs.hello.bbb);
  },
  // setup(){
  //   console.log("App setup");
  //   const counter = ref(0)
  //   const user =reactive({name:"张三",age:18})

  //   function changeCounter(){
  //     counter.value++
  //   }

  //   function changeUserName(){
  //     user.name="李四"
  //   }

  //   return {counter,user,changeCounter,changeUserName}
  // }
}
</script>
<template>

  <div>
    <h1> 商品数量：{{ store.count }}</h1>
  </div>

  <div id="app">
    <h1>Hello App!</h1>
    <p>
      <!--使用 router-link 组件进行导航 -->
      <!--通过传递 `to` 来指定链接 -->
      <!--`<router-link>` 将呈现一个带有正确 `href` 属性的 `<a>` 标签-->
      <router-link to="/">Go to Home</router-link>
      <br />
      <router-link to="/about">Go to About</router-link>
      <br />
      <router-link to="/user/123">Go to User</router-link>
      <br />
      <router-link to="/news/456">Go to News</router-link>
      <br />
      <router-link to="/parent">Go to Parent</router-link>
      <br />
      <router-link to="/page">Go to Page</router-link>
    </p>
    <!-- 路由出口 -->
    <!-- 路由匹配到的组件将渲染在这里 -->
    <router-view name="ShopTop"></router-view>
    <router-view  ></router-view>
    <router-view name="ShopFooter"></router-view>
  </div>
  <div>

    <h2>{{ counter }}</h2>
    <button @click="changeCounter">改变counter</button>
    <h2>{{ user.name }}</h2>
    <button @click="changeUserName">改变user.name</button>
  </div>
  <p>------------------------------------------------------------------</p>
  <div>
    <Content :mseeage="msg"></Content>
    <p>-------------------------------------------</p>
    <Content @injectMsg="getChildMsg"></Content>

    <p>-------------------------------------------</p>
    <Content @injectMsg="getChildMsg" ref="hello"></Content>
    <p>----------------插槽---------------------------</p>
    <Content2>
      <template v-slot:aaa>
        <button>按钮aaa</button>
        {{ childMsg }}
      </template>
      <template v-slot:bbb>
        <button>按钮bbb</button>
      </template>
      <template v-slot:ccc>
        <button>按钮ccc</button>
      </template>
    </Content2>
    <p>----------------------------------------------</p>
    <Content2>
      <template v-slot:default="slotProps">
        {{ slotProps.list }}
        <ul>
          <li v-for="(item, index) in slotProps.list" :key="index">
            {{ item }}
          </li>
        </ul>
      </template>
    </Content2>
    <Content2>
      <template v-slot:default="slotProps">
        {{ slotProps.list }}
        <ol>
          <li v-for="(item, index) in slotProps.list" :key="index">
            {{ item }}
          </li>
        </ol>
      </template>
    </Content2>
    <p>----------------------------------------------</p>
    <HomeView></HomeView>
  </div>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
    </div>
  </header>

  <main>
    <TheWelcome />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
