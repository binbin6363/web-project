<template>
  <el-card class="account-container">
    
  </el-card>
</template>

<script>
export default {
  name: 'Product',
  setup() {
    const goJuejin = () => {
      console.log('goJuejin')
      window.open('https://juejin.cn/book/6933939264455442444', 'target')
    }

    return {
      goJuejin
    }
  }
}
</script>
<style scoped>
  .course {
    font-weight: bold;
    display: inline-block;
  }
  .course li {
    line-height: 36px;
    color: #409eff;
    cursor: pointer;
  }
</style>