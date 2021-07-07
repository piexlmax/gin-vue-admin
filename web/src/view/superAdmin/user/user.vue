<template>
  <div>
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" icon="el-icon-plus" @click="addUser">新增用户</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="头像" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <CustomPic :pic-src="scope.row.headerImg" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="uuid" min-width="250" prop="uuid" />
      <el-table-column label="用户名" min-width="150" prop="userName" />
      <el-table-column label="昵称" min-width="150" prop="nickName" />
      <el-table-column label="用户角色" min-width="150">
        <template slot-scope="scope">
          <el-cascader
            v-model="scope.row.authority.authorityId"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            filterable
            @change="changeAuthority(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="150">
        <template slot-scope="scope">
          <el-popover v-model="scope.row.visible" placement="top" width="160">
            <p>确定要删除此用户吗</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
            </div>
            <el-button slot="reference" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      layout="total, sizes, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />

    <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" title="新增用户">
      <el-form ref="userForm" :rules="rules" :model="userInfo">
        <el-form-item label="用户名" label-width="80px" prop="username">
          <el-input v-model="userInfo.username" />
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item label="别名" label-width="80px" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item label="头像" label-width="80px">
          <div style="display:inline-block" @click="openHeaderChange">
            <img v-if="userInfo.headerImg" class="header-img-box" :src="userInfo.headerImg">
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>
        <el-form-item label="用户角色" label-width="80px" prop="authorityId">
          <el-cascader
            v-model="userInfo.authorityId"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            filterable
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeAddUserDialog">取 消</el-button>
        <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
      </div>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import {
  getUserList,
  setUserAuthority,
  register,
  deleteUser
} from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import infoList from '@/mixins/infoList'
import { mapGetters } from 'vuex'
import CustomPic from '@/components/customPic'
import ChooseImg from '@/components/chooseImg'
export default {
  name: 'Api',
  components: { CustomPic, ChooseImg },
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 5, message: '最低5位字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' },
          { min: 6, message: '最低6位字符', trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: '请输入用户昵称', trigger: 'blur' }
        ],
        authorityId: [
          { required: true, message: '请选择用户角色', trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  async created() {
    this.getTableData()
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    this.setOptions(res.data.list)
  },
  methods: {
    openHeaderChange() {
      this.$refs.chooseImg.open()
    },
    setOptions(authData) {
      this.authOptions = []
      this.setAuthorityOptions(authData, this.authOptions)
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            this.setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.ID })
      if (res.code === 0) {
        this.getTableData()
        row.visible = false
      }
    },
    async enterAddUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo)
          if (res.code === 0) {
            this.$message({ type: 'success', message: '创建成功' })
          }
          await this.getTableData()
          this.closeAddUserDialog()
        }
      })
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields()
      this.addUserDialog = false
    },
    handleAvatarSuccess(res) {
      this.userInfo.headerImg = res.data.file.url
    },
    addUser() {
      this.addUserDialog = true
    },
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authorityId: row.authority.authorityId
      })
      if (res.code === 0) {
        this.$message({ type: 'success', message: '角色设置成功' })
      }
    }
  }
}
</script>

<style lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>
