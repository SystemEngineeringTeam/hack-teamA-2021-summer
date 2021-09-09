<template>
  <div>
    <div class="header">
      <div class="header-logo">卍亜琉派華潘堕合法賭博遊戯卍</div>
      <div class="header-list"></div>
    </div>

    <div class="main">
      <div class="copy-container">
        <h1>アルパカとパンダの楽しい遊び</h1>
      </div>

      <div class="contents">
        <h3 class="section-title">ログイン(入力してください)</h3>
        <div class="contents-item3">
          <p>メールアドレス</p>
        </div>
        <input class="mail" v-model="mail" />
        <div class="contents-item2">
          <p>パスワード</p>
        </div>
        <input class="pwd" v-model="password" />
        <div class="check">
          <v-btn @click="login()">
            <a href="/main"> ログイン</a>
          </v-btn>
        </div>
      </div>

      <div class="contact-form"></div>
      <a href="/singup"><li class="header-item">新規登録</li></a>
    </div>
    <img
      class="loanshark"
      src="https://frame-illust.com/fi/wp-content/uploads/2017/03/9593.png"
    />
    <img
      class="loanshark2"
      src="https://frame-illust.com/fi/wp-content/uploads/2017/03/9694.png"
    />
  </div>
</template>

<script>
export default {
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
  position: relative;
  width: 100%;
  height: 100%;
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
/*メルアド*/
.contents-item3 {
  position: absolute;
  float: left;
  font-size: 30px;
  text-align: left;
  color: #800000;
  top: 590px;
  left: 100px;
}

/*パスワード*/
.contents-item2 {
  position: absolute;
  float: left;
  font-size: 30px;
  text-align: left;
  color: #800000;
  top: 780px;
  left: 100px;
}

/*メールアドレスを打ち込むところ*/
.mail {
  position: absolute;
  float: left;
  width: 400px;
  height: 70px;
  top: 645px;
  left: 100px;
  border: 3px solid #4b0082;
  background-color: #4b0082;
  color: #fff;
  font-size: 30px;
}

/*パスワードを打ち込むところ*/
.pwd {
  position: absolute;
  float: left;
  width: 400px;
  height: 70px;
  top: 835px;
  left: 100px;
  border: 3px solid #4b0082;
  background-color: #4b0082;
  color: #fff;
  font-size: 30px;
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
button {
  position: absolute;
  float: left;
  width: 120px;
  height: 70px;
  font-size: 24px;
  color: #b22222;
  background-color: #4b0082;
  top: 835px;
  left: 600px;
}

/*新規ログイン『ようこそ！』ボタン*/
.header-item {
  position: absolute;
  float: left;
  width: 150px;
  height: 50px;
  font-size: 30px;
  text-align: center;
  color: #b22222;
  background-color: #4b0082;
  top: 1000px;
  left: 100px;
}

/*アルパカの画像*/
.loanshark {
  position: absolute;
  width: 250px;
  height: 250px;
  left: 1100px;
  bottom: 100px;
}

/*パンダの画像*/
.loanshark2 {
  position: absolute;
  width: 250px;
  height: 250px;
  left: 1500px;
  bottom: 100px;
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
