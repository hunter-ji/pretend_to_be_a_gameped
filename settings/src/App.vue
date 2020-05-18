<template>
  <div id="app">
    <div class="container">
      <!--项目名字-->
      <Project-title />
      <!--设置客户端数量-->
      <el-card shadow="always" class="box">
        <div class="client-number">
          <i class="el-icon-mobile-phone" />
          <div class="client-number__title">客户端数量</div>
          <el-input-number
            v-model="num"
            @change="handleChange"
            :min="1" :max="10"
            label="描述文字" />
        </div>
      </el-card>
      <!--客户端设置-->
      <div v-for="(item, index) in form" :key="index">
        <handle :num="index + 1" :form="item" @removehandle="remove_handle" />
      </div>

      <!--按钮-->
      <div class="handle-btn-group">
        <el-button type="primary" style="width: 50%;" class="shadow" @click="handle_submit">启动</el-button>
        <el-button style="width: 50%;" class="shadow" @click="handle_reset">重置</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import("./style.css");
import ProjectTitle from "./components/ProjectTitle";
import handle from "./components/handle";

export default {
  name: "App",
  components: {
    handle,
    ProjectTitle
  },
  data() {
    return {
      num: 1,
      handle_object: {
        left1: "",
        left2: "",
        left3: "",
        left4: "",
        right1: "",
        right2: "",
        right3: "",
        right4: ""
      },
      form: []
    }
  },
  mounted() {
    this.fetch_data();
  },
  watch: {
    form(val) {
      this.num = val.length;
    }
  },
  methods: {
    fetch_data() {
      this.add_handle();
    },
    handleChange(currentValue, oldValue) {
      if (currentValue > oldValue) {
        this.add_handle();
      } else {
        this.remove_handle(currentValue + 1);
      }
    },
    add_handle() {
      this.form.push(
        JSON.parse(JSON.stringify(this.handle_object))
      );
    },
    remove_handle(num) {
      new Promise(resolve => {
        this.form.splice(num - 1, 1);
        resolve()
      })
        .then(() => {
          this.$message({
            message: `删除${ num }号手柄成功`,
            type: "success"
          })
        })
        .catch(() => {
          this.$message({
            message: `删除${ num }手柄失败`,
            type: "error"
          })
        })

    },
    handle_submit() {
      console.log("待开发...");
    },
    handle_reset() {
      this.form = [];
      this.add_handle();
    }
  }
}
</script>

<style>
body {
  padding: 0;
  margin: 0;
  background: #E4E7ED;
}

.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 100px 0;
}

.client-number {
  display: flex;
  flex-direction: row;
  font-size: large;
  align-items: center;
  justify-content: center;
}

.client-number__title {
  margin-right: 50px;
}

.handle-btn-group {
  display: flex;
  flex-direction: row;
  width: 800px;
  margin: auto;
  padding-top: 50px;
}


</style>
