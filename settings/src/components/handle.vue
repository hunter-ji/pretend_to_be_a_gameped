<template>
  <el-card shadow="always" class="box handle-box">
    <div slot="header">
      <span>{{ num }}号手柄</span>
    </div>
    <div class="handle-body">
      <div class="handle-item">
        <el-form label-width="50px" :model="form">
          <el-form-item label="左1">
            <el-input v-model="form.left1" />
          </el-form-item>
          <el-form-item label="左2">
            <el-input v-model="form.left2" />
          </el-form-item>
          <el-form-item label="左3">
            <el-input v-model="form.left3" />
          </el-form-item>
          <el-form-item label="左4">
            <el-input v-model="form.left4" />
          </el-form-item>
        </el-form>
      </div>
      <div class="handle-item">
        <el-form label-width="50px" :model="form">
          <el-form-item label="右1">
            <el-input v-model="form.right1" />
          </el-form-item>
          <el-form-item label="右2">
            <el-input v-model="form.right2" />
          </el-form-item>
          <el-form-item label="右3">
            <el-input v-model="form.right3" />
          </el-form-item>
          <el-form-item label="右4">
            <el-input v-model="form.right4" />
          </el-form-item>
        </el-form>
      </div>
    </div>
  </el-card>
</template>

<script>
export default {
  props: {
    num: Number,
    form: Object
  },
  data() {
    return {}
  }
}
</script>

<style scoped>
.handle-body {
  display: flex;
  justify-content: space-around;
}

.handle-box {
  padding: 12px 0;
  margin-top: 30px;
}

.handle-item {
  width: 40%;
}
</style>