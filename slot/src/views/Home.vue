<template>
<div>
  <div class="header">
    <div class="header-logo">卍亜琉派華潘堕合法賭博遊戯卍</div>
  </div>

  <div class="home" >
    <center>
    <!-- ボタン押したら任意のリンクに遷移する -->
    <img class="alpak-img" src="https://frame-illust.com/fi/wp-content/uploads/2017/03/9694.png">
    <img class="slot-img" src="@/assets/money_slot_machine.png">
    <img class="panda-img" src="https://frame-illust.com/fi/wp-content/uploads/2017/03/9593.png">
    </center>
    <br>
    <center>
    <button @click="$router.push('login')">遊戯する</button>
    </center>
  </div>
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
   position:relative;
  background-color: #b22222;
  color: #3f3;
  height: 240px;
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
    width:400px;
    height: 400px;
    /* position:relative; */
    left: 210px;
    bottom: 250px;
}
.panda-img{
  /* position:relative; */
    width:300px;
    height:300px;
    /* position:relative; */
    /* left: 210px; */
    /* bottom: 250px; */
}

.alpak-img{
   /* position:relative; */
    width:300px;
    height:300px;
    /* position:relative; */
    /* left: 210px; */
    /* bottom: 250px; */
}

.home{
  font-size: 40px;
  /* color:#000000; */
  background-color:#000000
   /* height: 300px; */
}
</style>
