<template>
  <div>
    <!-- 从数据库直接获取字段 -->
    <el-collapse v-model="activeNames">
      <el-collapse-item name="1">
        <template slot="title">
          <div :style="{fontSize:'16px',paddingLeft:'20px'}">
            点这里从现有数据库创建代码
            <i class="header-icon el-icon-thumb"></i>
          </div>
        </template>
        <el-form ref="getTableForm" :inline="true" :model="dbform" label-width="120px">
          <el-form-item label="数据库名" prop="structName">
            <el-select @change="getTable" v-model="dbform.dbName" filterable placeholder="请选择数据库">
              <el-option
                v-for="item in dbOptions"
                :key="item.database"
                :label="item.database"
                :value="item.database"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="表名" prop="structName">
            <el-select
              v-model="dbform.tableName"
              :disabled="!dbform.dbName"
              filterable
              placeholder="请选择表"
            >
              <el-option
                v-for="item in tableOptions"
                :key="item.tableName"
                :label="item.tableName"
                :value="item.tableName"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button @click="getColumn" type="primary">使用此表创建</el-button>
          </el-form-item>
        </el-form>
      </el-collapse-item>
    </el-collapse>

    <el-divider></el-divider>

    <div>
      <el-form ref="autoCodeForm" label-width="120px" :inline="true">
        <el-form-item label="名称">
          <el-input v-model="autoFillName" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="前缀">
          <el-input v-model="autoFillPrefix" placeholder="请输入前缀" />
        </el-form-item>
      </el-form>
    </div>

    <!-- 初始版本自动化代码工具 -->
    <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
      <el-form-item label="Struct名称" prop="structName">
        <el-input v-model="form.structName" placeholder="首字母自动转换大写"></el-input>
      </el-form-item>
      <el-form-item label="tableName" prop="tableName">
        <el-input v-model="form.tableName" placeholder="指定表名（非必填）"></el-input>
      </el-form-item>
      <el-form-item label="Struct简称" prop="abbreviation">
        <el-input v-model="form.abbreviation" placeholder="简称会作为入参对象名和路由group"></el-input>
      </el-form-item>
      <el-form-item label="Struct中文名称" prop="description">
        <el-input v-model="form.description" placeholder="中文描述作为自动api描述"></el-input>
      </el-form-item>
      <el-form-item label="文件名称" prop="packageName">
        <el-input v-model="form.packageName"></el-input>
      </el-form-item>
      <el-form-item label="自动创建api">
        <el-checkbox v-model="form.autoCreateApiToSql"></el-checkbox>
      </el-form-item>
      <el-form-item label="自动移动文件">
        <el-checkbox v-model="form.autoMoveFile"></el-checkbox>
      </el-form-item>
    </el-form>
    <!-- 组件列表 -->
    <div class="button-box clearflex">
      <el-button @click="editAndAddField()" type="primary">新增Field</el-button>
    </div>
    <el-table :data="form.fields" border stripe>
      <el-table-column type="index" label="序列" width="100"></el-table-column>
      <el-table-column prop="fieldName" label="Field名"></el-table-column>
      <el-table-column prop="fieldDesc" label="中文名"></el-table-column>
      <el-table-column prop="fieldJson" label="FieldJson"></el-table-column>
      <el-table-column prop="fieldType" label="Field数据类型" width="130"></el-table-column>
      <el-table-column prop="dataType" label="数据库字段类型" width="130"></el-table-column>
      <el-table-column prop="dataTypeLong" label="数据库字段长度" width="130"></el-table-column>
      <el-table-column prop="columnName" label="数据库字段" width="130"></el-table-column>
      <el-table-column prop="comment" label="数据库字段描述" width="130"></el-table-column>
      <el-table-column prop="fieldSearchType" label="搜索条件" width="130"></el-table-column>
      <el-table-column prop="dictType" label="字典" width="130"></el-table-column>
      <el-table-column label="操作" width="300">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="editAndAddField(scope.row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="scope.$index == 0"
            @click="moveUpField(scope.$index)"
          >上移</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="(scope.$index + 1) == form.fields.length"
            @click="moveDownField(scope.$index)"
          >下移</el-button>
          <el-popover placement="top" v-model="scope.row.visible">
            <p>确定删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteField(scope.$index)">确定</el-button>
            </div>
            <el-button size="mini" type="danger" icon="el-icon-delete" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-tag type="danger">id , created_at , updated_at , deleted_at 会自动生成请勿重复创建</el-tag>
    <!-- 组件列表 -->
    <div class="button-box clearflex">
      <el-button @click="enterForm" type="primary">生成代码</el-button>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog title="组件内容" :visible.sync="dialogFlag">
      <FieldDialog v-if="dialogFlag" :dialogMiddle="dialogMiddle" ref="fieldDialog" />
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
const fieldTemplate = {
  fieldName: "",
  fieldDesc: "",
  fieldType: "",
  dataType: "",
  fieldJson: "",
  columnName: "",
  dataTypeLong: "",
  comment: "",
  fieldSearchType: "",
  dictType: ""
};

const pluralize = require('pluralize');
import FieldDialog from "@/view/systemTools/autoCode/component/fieldDialog.vue";
import { toUpperCase, toHump, toSQLLine } from "@/utils/stringFun.js";
import { createTemp, getDB, getTable, getColumn } from "@/api/autoCode.js";
import { getDict } from "@/utils/dictionary";

export default {
  name: "autoCode",
  data() {
    return {
      activeNames: [""],
      dbform: {
        dbName: "",
        tableName: ""
      },
      dbOptions: [],
      tableOptions: [],
      addFlag: "",
      fdMap: {},
      form: {
        structName: "",
        tableName: "",
        packageName: "",
        abbreviation: "",
        description: "",
        autoCreateApiToSql: false,
        autoMoveFile: false,
        fields: []
      },
      rules: {
        structName: [
          { required: true, message: "请输入结构体名称", trigger: "blur" }
        ],
        abbreviation: [
          { required: true, message: "请输入结构体简称", trigger: "blur" }
        ],
        description: [
          { required: true, message: "请输入结构体描述", trigger: "blur" }
        ],
        packageName: [
          {
            required: true,
            message: "文件名称：sys_xxxx_xxxx",
            trigger: "blur"
          }
        ]
      },
      dialogMiddle: {},
      bk: {},
      dialogFlag: false,
      autoFillName: "",
      autoFillPrefix: "",
    };
  },
  components: {
    FieldDialog,
  },
  methods: {
    editAndAddField(item) {
      this.dialogFlag = true;
      if (item) {
        this.addFlag = "edit";
        this.bk = JSON.parse(JSON.stringify(item));
        this.dialogMiddle = item;
      } else {
        this.addFlag = "add";
        this.dialogMiddle = JSON.parse(JSON.stringify(fieldTemplate));
      }
    },
    moveUpField(index) {
      if (index == 0) {
        return;
      }
      const oldUpField = this.form.fields[index - 1];
      this.form.fields.splice(index - 1, 1);
      this.form.fields.splice(index, 0, oldUpField);
    },
    moveDownField(index) {
      const fCount = this.form.fields.length;
      if (index == fCount - 1) {
        return;
      }
      const oldDownField = this.form.fields[index + 1];
      this.form.fields.splice(index + 1, 1);
      this.form.fields.splice(index, 0, oldDownField);
    },
    enterDialog() {
      this.$refs.fieldDialog.$refs.fieldDialogFrom.validate(valid => {
        if (valid) {
          this.dialogMiddle.fieldName = toUpperCase(toHump(
            this.dialogMiddle.fieldName
          ));
          if (this.addFlag == "add") {
            this.form.fields.push(this.dialogMiddle);
          }
          this.dialogFlag = false;
        } else {
          return false;
        }
      });
    },
    closeDialog() {
      if (this.addFlag == "edit") {
        this.dialogMiddle = this.bk;
      }
      this.dialogFlag = false;
    },
    deleteField(index) {
      this.form.fields.splice(index, 1);
    },
    async enterForm() {
      if (this.form.fields.length <= 0) {
        this.$message({
          type: "error",
          message: "请填写至少一个field"
        });
        return false;
      }
      if (
        this.form.fields.some(item => item.fieldName == this.form.structName)
      ) {
        this.$message({
          type: "error",
          message: "存在与结构体同名的字段"
        });
        return false;
      }
      this.$refs.autoCodeForm.validate(async valid => {
        if (valid) {
          this.form.structName = toUpperCase(this.form.structName);
          if (this.form.structName == this.form.abbreviation) {
            this.$message({
              type: "error",
              message: "structName和struct简称不能相同"
            });
            return false;
          }
          const data = await createTemp(this.form);
          if (data.headers?.success == "false") {
            return;
          } else {
            this.$message({
              type: "success",
              message: "自动化代码创建成功，正在下载"
            });
          }
          const blob = new Blob([data]);
          const fileName = "ginvueadmin.zip";
          if ("download" in document.createElement("a")) {
            // 不是IE浏览器
            let url = window.URL.createObjectURL(blob);
            let link = document.createElement("a");
            link.style.display = "none";
            link.href = url;
            link.setAttribute("download", fileName);
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link); // 下载完成移除元素
            window.URL.revokeObjectURL(url); // 释放掉blob对象
          } else {
            // IE 10+
            window.navigator.msSaveBlob(blob, fileName);
          }
        } else {
          return false;
        }
      });
    },
    async getDb() {
      const res = await getDB();
      if (res.code == 0) {
        this.dbOptions = res.data.dbs;
      }
    },
    async getTable() {
      const res = await getTable({ dbName: this.dbform.dbName });
      if (res.code == 0) {
        this.tableOptions = res.data.tables;
      }
      this.dbform.tableName = "";
    },
    async getColumn() {
      const gormModelList = ["id", "created_at", "updated_at", "deleted_at"];
      const res = await getColumn(this.dbform);
      if (res.code == 0) {
        const tbHump = toHump(this.dbform.tableName);
        this.form.structName = toUpperCase(tbHump);
        this.form.tableName = this.dbform.tableName;
        this.form.packageName = tbHump;
        this.form.abbreviation = tbHump;
        this.form.description = tbHump + "表";
        this.form.autoCreateApiToSql = true;
        this.form.fields = [];
        res.data.columns &&
          res.data.columns.map(item => {
            if (!gormModelList.some(gormfd => gormfd == item.columnName)) {
              const fbHump = toHump(item.columnName);
              this.form.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + "字段",
                fieldType: this.fdMap[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong,
                columnName: item.columnName,
                comment: item.columnComment,
                fieldSearchType: "",
                dictType: ""
              });
            }
          });
      }
    },
    async setFdMap() {
      const fdTypes = ["string", "int", "bool", "float64", "time.Time"];
      fdTypes.map(async fdtype => {
        const res = await getDict(fdtype);
        res.map(item => {
          this.fdMap[item.label] = fdtype;
        });
      });
    },
    autoFillStructField(name, prefix) {
      let name_copy = "";
      if (name !== "") {
        name = toSQLLine(name.replace(name[0], name[0].toLowerCase()));
        name_copy = pluralize(name);
      }
      this.form.structName = toUpperCase(toHump(name));
      this.form.tableName = prefix + name_copy;
      this.form.abbreviation = toHump(name);
      this.form.packageName = prefix + name;
    },
    autoFillInit() {
      this.autoFillPrefix = window.localStorage.getItem('struct_auto_fill_prefix');
    },
  },
  watch: {
    autoFillName: function (val) {
      this.autoFillStructField(val, this.autoFillPrefix);
    },
    autoFillPrefix: function (val) {
      window.localStorage.setItem("struct_auto_fill_prefix", val);
      this.autoFillStructField(this.autoFillName, val);
    },
  },
  created() {
    this.getDb();
    this.setFdMap();
    this.autoFillInit();
  }
};
</script>
<style scope lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
</style>
