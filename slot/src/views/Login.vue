<template>
  <div>
    <div class="header">
      <div class="header-logo">チームA</div>
      <div class="header-list">
        <ul>
          <li>ログイン</li>
          <li>新規ログイン</li>
          <li>ゲーム概要</li>
        </ul>
      </div>
    </div>

    <div class="main">
      <div class="copy-container">
        <h1>スロットゲーム（仮）</h1>
      </div>

      <div class="contents">
        <h3 class="section-title">ログイン(入力してください)</h3>
        <div class="contents-item">
          <p>メールアドレス</p>
          <input v-model="mail" />
          <p>パスワード</p>
          <input v-model="password" />
        </div>
        <button @click="login()">確認</button>
      </div>
      <div class="contact-form"></div>
    </div>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      mail: "",
      password: "",
    };
  },
  methods: {
    login() {
      this.axios
        .post("http://localhost:8080/login", {
          email: this.mail,
          password: this.password,
        },)
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  actions: {},
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
  background-color: #26d0c9;
  color: #fff;
  height: 90px;
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
  float: left;
  font-size: 36px;
  padding: 20px 40px;
}

/*heder、項目の大きさ、位置*/
.header-list li {
  float: left;
  padding: 33px 20px;
}
/*headerおわり*/
/*mainの設定*/
.main {
  padding: 100px 80px;
}

.copy-container h1 {
  font-size: 100px;
}

/*メールアドレスを打ち込むところ*/
input {
  border: 3px solid #000;
}

.contents {
  height: 500px;

  margin-top: 100px;
}

.section-title {
  border-bottom: 2px solid #dee7ec;

  font-size: 28px;

  padding-bottom: 15px;

  margin-bottom: 50px;
}

.contents-item {
  float: left;

  margin-right: 40px;
}

.contents-item p {
  font-size: 24px;
  margin-top: 30px;
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
