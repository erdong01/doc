<script lang="ts">
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router';
export default defineComponent({
  name: "HomeView",
  setup() {
    const router = useRouter()
    console.log(router.getRoutes());
    const list =router.getRoutes().filter(v=>v.meta.isShow)
    console.log(list);
    return {list}
  },
  components: {},
})
</script>
<template>
  <div class="home">
    <div class="common-layout">
      <el-container>
        <el-header height="80px"> <el-row :gutter="20">
            <el-col :span="4">
              <img src="../assets/logo.svg" class="logo" />
            </el-col>
            <el-col :span="16">
              <h2>后台管理系统</h2>
            </el-col>
            <el-col :span="4">
              <span class="quit-login">退出登录</span>
            </el-col>
          </el-row>
        </el-header>
        <el-container>
          <el-aside width="200px"><el-menu 
            active-text-color="#ffd04b" background-color="#545c64"
              class="el-menu-vertical-demo" default-active="2" text-color="#fff" router>
              <el-menu-item :index="item.path" v-for="item in list" :key="item.path">
                <span>{{item.meta.title}}</span>
              </el-menu-item>
            </el-menu>
          </el-aside>
          <el-main>
            <RouterView></RouterView>
          </el-main>
        </el-container>
      </el-container>
    </div>

  </div>
</template>
<style lang="scss" scoped>
.el-header {
  height: 80px;
  background-color: #666;

  .logo {
    height: 80px;
  }
}

h2 {
  text-align: center;
  height: 80px;
  color: azure;
}

.quit-login {
  line-height: 80px;
  color: azure;
}

.el-aside {
  .el-menu {
    height: calc(100vh - 80px);
  }
}
</style>