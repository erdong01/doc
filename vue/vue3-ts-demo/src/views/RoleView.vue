<template>
    <div>
        <div class="select-box">
            <el-form :inline="true" class="demo-form-inline">
                <el-button type="primary" @click="addRole">添加</el-button>
            </el-form>
        </div>

        <el-table :data="list"  style="width: 100%">
            <el-table-column prop="roleId" label="ID" width="180" />
            <el-table-column prop="roleName" label="角色名" width="180" />
            <el-table-column prop="authority" label="操作">
                <template #default="scope">
                    <el-button link size="small" @click="changeRole(scope.row)">
                        修改权限
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
    <RouterView></RouterView>
</template>

<script lang="ts">
import { ElMessage, ElMessageBox } from 'element-plus';
import { defineComponent, onMounted, reactive, toRefs } from 'vue';
import { useRouter } from 'vue-router';
import { getRoleList } from '../request/api';
import { InitData, type ListInt } from '../type/role';
export default defineComponent({
    setup() {
        const router = useRouter()
        const data = reactive(new InitData())
        onMounted(() => {
            getRoleList().then((res) => {
                console.log(res)
                data.list = res.data
            })
        })
        const addRole = () => {
            ElMessageBox.prompt('请输入角色名称', '增加', {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
            }).then(({ value }) => {
                if (value) {//判断输入框中有值，就将对应的值添加到列表中
                    data.list.push({ roleId: data.list.length + 1, roleName: value, authority: [] })
                }
                ElMessage({
                    type: 'success',
                    message: `${value}添加成功`,
                })
            })
                .catch(() => {
                    ElMessage({
                        type: 'info',
                        message: '取消操作',
                    })
                })
        }
        const changeRole = (row: ListInt) => {
            router.push({
                name: 'authority',
                params: {
                    id: row.roleId,
                    authority: row.authority
                }
            })
        }
        return { ...toRefs(data), changeRole, addRole }
    }
})
</script>

<style scoped>

</style>