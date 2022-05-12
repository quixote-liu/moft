<template>
  <div>
    <el-dialog v-bind="$attrs" v-on="$listeners" @open="onOpen" @close="onClose" title="Dialog Titile">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-form-item label="单行文本" prop="field101">
          <el-input v-model="formData.field101" placeholder="请输入单行文本" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="field102">
          <el-input v-model="formData.field102" placeholder="请输入密码" clearable show-password
            :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="开关" prop="field103" required>
          <el-switch v-model="formData.field103"></el-switch>
        </el-form-item>
        <el-form-item label="级联选择" prop="field104">
          <el-cascader v-model="formData.field104" :options="field104Options" :props="field104Props"
            :style="{width: '100%'}" placeholder="请选择级联选择" clearable></el-cascader>
        </el-form-item>
        <el-form-item label="多选框组" prop="field105">
          <el-checkbox-group v-model="formData.field105" size="medium">
            <el-checkbox v-for="(item, index) in field105Options" :key="index" :label="item.value"
              :disabled="item.disabled">{{item.label}}</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="下拉选择" prop="field106">
          <el-select v-model="formData.field106" placeholder="请选择下拉选择" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in field106Options" :key="index" :label="item.label"
              :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="单选框组" prop="field107">
          <el-radio-group v-model="formData.field107" size="medium">
            <el-radio v-for="(item, index) in field107Options" :key="index" :label="item.value"
              :disabled="item.disabled">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer">
        <el-button @click="close">取消</el-button>
        <el-button type="primary" @click="handelConfirm">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
export default {
  inheritAttrs: false,
  components: {},
  props: [],
  data() {
    return {
      formData: {
        field101: undefined,
        field102: undefined,
        field103: false,
        field104: [],
        field105: [],
        field106: undefined,
        field107: 2,
      },
      rules: {
        field101: [{
          required: true,
          message: '请输入单行文本',
          trigger: 'blur'
        }],
        field102: [{
          required: true,
          message: '请输入密码',
          trigger: 'blur'
        }],
        field104: [{
          required: true,
          type: 'array',
          message: '请至少选择一个级联选择',
          trigger: 'change'
        }],
        field105: [{
          required: true,
          type: 'array',
          message: '请至少选择一个多选框组',
          trigger: 'change'
        }],
        field106: [{
          required: true,
          message: '请选择下拉选择',
          trigger: 'change'
        }],
        field107: [{
          required: true,
          message: '单选框组不能为空',
          trigger: 'change'
        }],
      },
      field104Options: [],
      field105Options: [{
        "label": "选项一",
        "value": 1
      }, {
        "label": "选项二",
        "value": 2
      }],
      field106Options: [{
        "label": "选项一",
        "value": 1
      }, {
        "label": "选项二",
        "value": 2
      }],
      field107Options: [{
        "label": "选项一",
        "value": 1
      }, {
        "label": "选项二",
        "value": 2
      }],
      field104Props: {
        "multiple": false,
        "label": "label",
        "value": "value",
        "children": "children"
      },
    }
  },
  computed: {},
  watch: {},
  created() {
    this.getField104Options()
  },
  mounted() {},
  methods: {
    onOpen() {},
    onClose() {
      this.$refs['elForm'].resetFields()
    },
    close() {
      this.$emit('update:visible', false)
    },
    handelConfirm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        this.close()
      })
    },
    getField104Options() {
      // 注意：this.$axios是通过Vue.prototype.$axios = axios挂载产生的
      this.$axios({
        method: 'get',
        url: 'https://www.fastmock.site/mock/f8d7a54fb1e60561e2f720d5a810009d/fg/cascaderList'
      }).then(resp => {
        var {
          data
        } = resp
        this.field104Options = data.list
      })
    },
  }
}

</script>
<style>
</style>
