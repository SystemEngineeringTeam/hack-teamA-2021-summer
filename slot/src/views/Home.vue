<template>
  <div class="header">
    <center>
    <div class="header-logo">卍亜琉派華潘堕合法賭博遊戯卍</div>
     </center>
     <center>
       <div class="home" >
      <!-- ボタン押したら任意のリンクに遷移する -->
      <img class="slot-img" src="@/assets/money_slot_machine.png">
      <br>
      <button @click="$router.push('login')">遊戯する</button>
       </div>
    </center>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'Home',
  components: {
    // HelloWorld
  }
}
</script>
<style scoped>
/*headerの設定、背景の色*/
.header {
  background-color: #b22222;
  color: #3f3;
  height: 300px;
}
/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
  float: left;
  font-size: 80px;
  padding: 20px 40px;
}
/* 遊戯するボタン */
button{
    /* position:relative;
    float: left; */
    width: 120px;
    height: 80px;
    font-size: 24px;
    color: #b22222;
    background-color: #4b0082;
    bottom: 150px;
    left: 550px;

}
/* スロットの画像 */
.slot-img{
  /* position:relative; */
    width:250px;
    height: 250px;
    /* position:relative; */
    left: 210px;
    bottom: 250px;
}
.home{
  font-size: 40px;
  /* color:#000000; */
  background-color:#000000
}
</style>
