<template>
  <div>
    <div class="header">
      <div class="header-logo">卍亜琉派華潘堕合法賭博遊戯卍
      </div>
      <div class="header-list">

      </div>
    </div>

    <div class="main">
      <div class="copy-container">
        <h1>アルパカとパンダの楽しい遊び</h1>
      </div>

      <div class="contents">
        <h3 class="section-title">ログイン(入力してください)</h3>
        <div class="contents-item">
          <p>メールアドレス</p>
          <input v-model="mail" />
          <p>パスワード</p>
          <input v-model="password" />
        </div>
        <div class="check">
          <button @click="login()">ログイン</button>
        </div>
      </div>
      
      <div class="contact-form"></div>
          <p class="newlogin">新規登録</p>
          
          <a href="/singup"><li class="header-item">ようこそ！</li></a>
      </div>
      <img class="loanshark" src="https://frame-illust.com/fi/wp-content/uploads/2017/03/9593.png">
      <img class="loanshark2" src="https://frame-illust.com/fi/wp-content/uploads/2017/03/9694.png">
      <selector></selector>
  </div>
</template>

<script>
import Selector from '@/components/ImageSelector.vue';
export default {
  components: {
    Selector
  },
  data: () => {
    return {
        // 変数定義（今回はメールアドレスとパスワード）
      mail: "",
      password: "",
    };
  },
// メゾット定義
  methods: {
    async login() {
      try {
        var res = await this.axios.post("http://localhost:8080/login", {
          email: this.mail,
          password: this.password,
        });
        this.$cookie.set("token", res.data.token);
        console.log(res);
      } catch (e) {
        console.log(e);
        alert("ログインに失敗しました");
      }
    },
  },
};
</script>
<style scoped>
body {
  font-family: "Avenir Next";
}

li {
  list-style: none;
}

/*headerの設定、背景の色*/
.header {
  background-color: #b22222;
  color: #3f3;
  height: 300px;
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
  float: left;
  font-size: 100px;
  padding: 20px 40px;
}

/*heder、項目の大きさ、位置*/
.header-list li {
  float: right;
  padding: 33px 20px;
}
/*headerおわり*/
/*mainの設定*/
/*メインの背景*/
.main {
  padding: 120px 100px;
  background-color: #000;
}
/*アルパカとパンダの楽しい遊び*/
.copy-container h1 {
  font-size: 80px;
  text-align: left;
  color: #800000;
}
/*ログイン入力してください*/
.section-title {
  font-size: 50px;
  text-align: left;
  color: #800000;
}
/*メルアド、パスワード*/
.contents-item {
  font-size: 50px;
  text-align: left;
  color: #800000;
}

/*メールアドレスを打ち込むところ*/
input {
  border: 3px solid #4b0082;
  background-color: #4b0082;
}

.contents {
  height: 500px;

  margin-top: 10px;
}

.section-title {
  border-bottom: 2px solid #800000;

  font-size: 28px;
  text-align: left;

  padding-bottom: 15px 2px;

  margin-bottom: 10px;
}

.contents-item {
  float: left;

  margin-right: 400px;
  margin-top: 10px;
  margin-bottom: 5px;
}

.contents-item p {
  font-size: 24px;
  margin-top: 30px;
}

/*ログインボタン*/
button{
    position:relative;
    float: left;
    width: 120px;
    height: 80px;
    font-size: 24px;
    color: #b22222;
    background-color: #4b0082;
    bottom: 150px;
    left: 550px;

}
/*新規ログイン*/
.newlogin{
  position:relative;
  float: left;
  font-size: 24px;
  color: #b22222;
  bottom: 100px;
  left: 115px;
}
/*新規ログイン『ようこそ！』ボタン*/
.header-item{
    position:relative;
    float: left;
    width: 120px;
    height: 80px;
    font-size: 24px;
    color: #b22222;
    background-color: #4b0082;
    bottom: 60px;
    left: 10px;

}

/*アルパカの画像*/
.loanshark{
    position:relative;
    width:250px;
    height: 250px;
    position:relative;
    left: 160px;
    bottom: 250px;
}

/*パンダの画像*/
.loanshark2{
    position:relative;
    width:250px;
    height: 250px;
    position:relative;
    left: 210px;
    bottom: 250px;
}

/*フッターの設定*/
.footer {
  background-color: #2f5167;
  color: #fff;
  height: 120px;
  padding: 40px;
}

.footer-logo {
  float: left;
  font-size: 32px;
}

.footer-list {
  float: right;
}

.footer-list li {
  padding-bottom: 20px;
}
</style>
