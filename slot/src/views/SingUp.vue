<template>
  <div>
    <body>
        <div class="header">
        <div class="header-logo">チームA</div>
        </div>

        <div class="main">
        <div class="copy-container">
            <h1>新規登録</h1>
        </div>
        
        <div class="contents">
            <h3 class="section-title">登録情報をご記入ください</h3>
            <div class="contents-item">
                <p>お名前</p>
                <input v-model="name"/>
                <p>メールアドレス</p>
                <input v-model="mail"/>
                <p>パスワード</p>
                <input v-model="password"/>
                <button @click="singup()">登録</button>
            </div>
        </div>
        
        <div class="contact-form">
        </div>
        </div>
        
        <div class="footer">
        <div class="footer-logo">hack-teamA-2021-summer</div>
        </div>
    </body>
    </div>    
</template>

<script>
export default {
  data: () => {
    return {
      name:"",
      mail: "",
      password: "",
    };
  },
  methods: {
    singup() {
      this.axios
        .post("http://localhost:8080/user", {
          name: this.name,
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
    background-color: #26d0c9;    color: #fff;
    height: 90px;
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
    float: left;
    font-size: 36px;
    padding: 20px 40px;
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
input{
    border: 4px solid #000;
    width: 400px;
    height: 30px;
}
.bottom{
    width: 40px;
    height: 40px;
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

.contents-item p{
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
</style>